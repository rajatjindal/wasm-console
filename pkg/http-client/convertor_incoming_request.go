package httpclient

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rajatjindal/wasm-console/internal/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

type IncomingRequest = types.IncomingRequest

// convert the IncomingRequest to http.Request
func NewHttpRequest(ir IncomingRequest) (req *http.Request, err error) {
	// convert the http method to string
	method, err := methodToString(ir.Method())
	if err != nil {
		return nil, err
	}

	// convert the path with query to a url
	var url string
	if pathWithQuery := ir.PathWithQuery(); pathWithQuery.None() {
		url = ""
	} else {
		url = *pathWithQuery.Some()
	}

	// convert the body to a reader
	var body io.Reader
	if consumeResult := ir.Consume(); consumeResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming request %s", *consumeResult.Err())
	} else if streamResult := consumeResult.OK().Stream(); streamResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming requests's stream %s", streamResult.Err())
	} else {
		body = NewReader(*streamResult.OK())
	}

	// create a new request
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// update additional fields
	toHttpHeader(ir.Headers(), &req.Header)

	return req, nil
}

func methodToString(m types.Method) (string, error) {
	if m.Connect() {
		return "CONNECT", nil
	}
	if m.Delete() {
		return "DELETE", nil
	}
	if m.Get() {
		return "GET", nil
	}
	if m.Head() {
		return "HEAD", nil
	}
	if m.Options() {
		return "OPTIONS", nil
	}
	if m.Patch() {
		return "PATCH", nil
	}
	if m.Post() {
		return "POST", nil
	}
	if m.Put() {
		return "PUT", nil
	}
	if m.Trace() {
		return "TRACE", nil
	}
	if other := m.Other(); other != nil {
		return *other, fmt.Errorf("unknown http method 'other'")
	}
	return "", fmt.Errorf("failed to convert http method")
}

func toHttpHeader(src types.Fields, dest *http.Header) {
	for _, f := range src.Entries().Slice() {
		key := string(f.F0)
		value := string(cm.List[uint8](f.F1).Slice())
		dest.Add(key, value)
	}
}
