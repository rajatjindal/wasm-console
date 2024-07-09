package httpclient

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rajatjindal/wasm-console/internal/wasi/cli/stdout"
	"github.com/rajatjindal/wasm-console/internal/wasi/http/types"
	"github.com/rajatjindal/wasm-console/internal/wasi/io/streams"
	"github.com/ydnar/wasm-tools-go/cm"
)

var _ http.ResponseWriter = &responseOutparamWriter{}

type responseOutparamWriter struct {
	// wasi response outparam is set at the end of http_trigger_handle
	outparam types.ResponseOutparam
	// wasi response
	response types.OutgoingResponse
	// wasi http headers
	wasiHeaders types.Fields
	// go httpHeaders are reconciled on call to WriteHeader, Flush or at the end of http_trigger_handle
	httpHeaders http.Header
	// wasi response body is set on first write because it can only be called once
	body *types.OutgoingBody
	// wasi response stream is set on first write because it can only be called once
	stream *streams.OutputStream

	statuscode int
}

func (row *responseOutparamWriter) Header() http.Header {
	return row.httpHeaders
}

func (row *responseOutparamWriter) Write(buf []byte) (int, error) {
	// acquire the response body's resource handle on first call to write
	if row.body == nil {
		bodyResult := row.response.Body()
		if bodyResult.IsErr() {
			return 0, fmt.Errorf("failed to acquire resource handle to response body: %s", bodyResult.Err())
		}
		row.body = bodyResult.OK()

		writeResult := row.body.Write()
		if writeResult.IsErr() {
			return 0, fmt.Errorf("failed to acquire resource handle for response body's stream: %s", writeResult.Err())
		}
		row.stream = writeResult.OK()
	}

	// //TODO: determine if we need to do these to fulfill the ResponseWriter contract
	// // call WriteHeader(http.StatusOK) if it hasn't been called yet
	// // call DetectContentType if headers doesn't contain content-type yet
	// // if total data is under "a few" KB and there are no flush calls, Content-Length is added automatically

	contents := cm.ToList(buf)
	writeResult := row.stream.Write(contents)
	if writeResult.IsErr() {
		if writeResult.Err().Closed() {
			return 0, fmt.Errorf("failed to write to response body's stream: closed")
		}

		//TODO: possible nil error here
		return 0, fmt.Errorf("failed to write to response body's stream: %s", writeResult.Err().LastOperationFailed().ToDebugString())
	}

	result := cm.OK[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](row.response)
	types.ResponseOutparamSet(row.outparam, result)

	return int(contents.Len()), nil
}

func (row *responseOutparamWriter) WriteHeader(statusCode int) {
	row.statuscode = statusCode
	row.reconcile()
}

// reconcile headers from go to wasi
func (row *responseOutparamWriter) reconcileHeaders() error {
	for key, vals := range row.httpHeaders {
		// convert each value distincly
		fieldVals := []types.FieldValue{}
		for _, val := range vals {
			fieldVals = append(fieldVals, types.FieldValue(cm.ToList([]uint8(val))))
		}

		if result := row.wasiHeaders.Set(types.FieldKey(key), cm.ToList(fieldVals)); result.IsErr() {
			switch *result.Err() {
			case types.HeaderErrorInvalidSyntax:
				return fmt.Errorf("failed to set header %s to [%s]: invalid syntax", key, strings.Join(vals, ","))
			case types.HeaderErrorForbidden:
				return fmt.Errorf("failed to set forbidden header key %s", key)
			case types.HeaderErrorImmutable:
				return fmt.Errorf("failed to set header on immutable header fields")
			default:
				return fmt.Errorf("not sure what happened here?")
			}
		}
	}

	//TODO: handle deleted headers

	return nil
}

// convert the ResponseOutparam to http.ResponseWriter
func NewHttpResponseWriter(out types.ResponseOutparam) *responseOutparamWriter {
	row := &responseOutparamWriter{
		outparam:    out,
		httpHeaders: http.Header{},
		wasiHeaders: types.NewFields(),
	}

	return row
}

func (row *responseOutparamWriter) reconcile() {
	err := row.reconcileHeaders()
	if err != nil {
		//TODO
	}

	//setting headers after this cause panic
	//TODO: debug
	row.response = types.NewOutgoingResponse(row.wasiHeaders)

	//set status code
	row.response.SetStatusCode(types.StatusCode(row.statuscode))
}

func println(msg string) {
	stdout.GetStdout().Write(cm.ToList([]byte(msg)))
	stdout.GetStdout().Write(cm.ToList([]byte("\n")))
}
