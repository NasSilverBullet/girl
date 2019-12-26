package girl

import (
	"io"
	"net/http"
	"net/http/httputil"
)

// Girl ...
type Girl struct {
	URL         string
	Method      string
	Headers     map[string]string
	Body        io.Reader
	ShowReqBody bool
	ShowResBody bool
}

func New(url string, method string, headers map[string]string, body io.Reader, showReqBody bool, showResBody bool) *Girl {
	return &Girl{
		URL:         url,
		Method:      method,
		Headers:     headers,
		Body:        body,
		ShowReqBody: showReqBody,
		ShowResBody: showResBody,
	}
}

func (g *Girl) Request() ([]byte, []byte, error) {
	req, err := http.NewRequest(g.Method, g.URL, g.Body)
	if err != nil {
		return nil, nil, err
	}

	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	reqDump, err := httputil.DumpRequestOut(req, g.ShowReqBody)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	resDump, err := httputil.DumpResponse(res, g.ShowResBody)
	if err != nil {
		return nil, nil, err
	}

	return reqDump, resDump, nil
}
