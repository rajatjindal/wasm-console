package commands

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	httpclient "github.com/rajatjindal/wasmshell/pkg/http-client"
)

func Curl(args []string) error {
	return run(strings.Join(args, " "))
}

func run(inp string) error {
	trimmedInp := strings.TrimSpace(inp)
	if trimmedInp == "" {
		return fmt.Errorf("no valid curl command found")
	}

	req, err := parseCurlCommand(trimmedInp)
	if err != nil {
		return err
	}

	resp, err := httpclient.Send(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := dumpResponse(resp, true)
	if err != nil {
		return err
	}

	fmt.Println(string(raw))

	return nil
}

func parseCurlCommand(curlCmd string) (*http.Request, error) {
	// Remove the 'curl ' prefix
	curlCmd = strings.TrimPrefix(curlCmd, "curl ")

	// Split command into parts by whitespace
	parts := strings.Fields(curlCmd)

	// Initialize request variables
	var method, urlStr, body string
	headers := make(http.Header)

	// Parse each part
	for i := 0; i < len(parts); i++ {
		switch parts[i] {
		case "-X", "--request":
			if i+1 < len(parts) {
				method = parts[i+1]
				i++
			}
		case "-H", "--header":
			if i+1 < len(parts) {
				header := strings.SplitN(parts[i+1], ":", 2)
				if len(header) == 2 {
					headers.Add(strings.TrimSpace(header[0]), strings.TrimSpace(header[1]))
				}
				i++
			}
		case "-d", "--data", "--data-raw", "--data-binary", "--data-urlencode":
			if i+1 < len(parts) {
				body = parts[i+1]
				i++
			}
		default:
			// Assume it's the URL if it doesn't start with a dash
			if !strings.HasPrefix(parts[i], "-") {
				urlStr = parts[i]
			}
		}
	}

	// Default method to GET if not specified
	if method == "" {
		method = http.MethodGet
	}

	// Parse the URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	bodyReader := strings.NewReader(body)
	req, err := http.NewRequest(method, parsedURL.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers to the request
	req.Header = headers

	return req, nil
}

// errNoBody is a sentinel error value used by failureToReadBody so we
// can detect that the lack of body was intentional.
var errNoBody = errors.New("sentinel error value")

// failureToReadBody is an io.ReadCloser that just returns errNoBody on
// Read. It's swapped in when we don't actually want to consume
// the body, but need a non-nil one, and want to distinguish the
// error from reading the dummy body.
type failureToReadBody struct{}

func (failureToReadBody) Read([]byte) (int, error) { return 0, errNoBody }
func (failureToReadBody) Close() error             { return nil }

var emptyBody = io.NopCloser(strings.NewReader(""))

func dumpResponse(resp *http.Response, body bool) ([]byte, error) {
	var b bytes.Buffer
	var err error
	save := resp.Body
	savecl := resp.ContentLength

	if !body {
		// For content length of zero. Make sure the body is an empty
		// reader, instead of returning error through failureToReadBody{}.
		if resp.ContentLength == 0 {
			resp.Body = emptyBody
		} else {
			resp.Body = failureToReadBody{}
		}
	} else if resp.Body == nil {
		resp.Body = emptyBody
	} else {
		save, resp.Body, err = drainBody(resp.Body)
		if err != nil {
			return nil, err
		}
	}
	err = resp.Write(&b)
	if err == errNoBody {
		err = nil
	}
	resp.Body = save
	resp.ContentLength = savecl
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return io.NopCloser(&buf), io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
