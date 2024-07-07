package httpclient

import (
	"net/http"

	"github.com/rajatjindal/wasmshell/internal/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

// convert the IncomingRequest to http.Request
func NewOutgoingHttpRequest(req *http.Request) (types.OutgoingRequest, error) {
	headers := types.NewFields()
	toWasiHeader(req.Header, headers)

	or := types.NewOutgoingRequest(headers)
	or.SetAuthority(cm.Some(req.Host))
	or.SetMethod(toWasiMethod(req.Method))
	or.SetPathWithQuery(cm.Some(req.URL.Path + "?" + req.URL.Query().Encode()))

	switch req.URL.Scheme {
	case "http":
		or.SetScheme(cm.Some(types.SchemeHTTP()))
	case "https":
		or.SetScheme(cm.Some(types.SchemeHTTPS()))
	default:
		or.SetScheme(cm.Some(types.SchemeOther(req.URL.Scheme)))
	}

	return or, nil
}

func toWasiHeader(src http.Header, dest types.Fields) {
	for k, v := range src {
		key := types.FieldKey(k)
		fieldVals := []types.FieldValue{}

		for _, val := range v {
			fieldVals = append(fieldVals, types.FieldValue(cm.ToList([]uint8(val))))
		}

		//TODO(rjindal): check error
		_ = dest.Set(key, cm.ToList(fieldVals))
	}
}

func toWasiMethod(s string) types.Method {
	switch s {
	case http.MethodConnect:
		return types.MethodConnect()
	case http.MethodDelete:
		return types.MethodDelete()
	case http.MethodGet:
		return types.MethodGet()
	case http.MethodHead:
		return types.MethodHead()
	case http.MethodOptions:
		return types.MethodOptions()
	case http.MethodPatch:
		return types.MethodPatch()
	case http.MethodPost:
		return types.MethodPost()
	case http.MethodPut:
		return types.MethodPut()
	case http.MethodTrace:
		return types.MethodTrace()
	default:
		return types.MethodOther(s)
	}
}
