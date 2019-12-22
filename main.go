package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
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

func NewGirl() *Girl {
	return &Girl{
		URL:         "https://example.com/",
		Method:      http.MethodGet,
		Headers:     map[string]string{"Content-Type": "application/json", "Authorization": "Bearer access-token", "tenant-id": "testid", "application-key": "testkey"},
		Body:        bytes.NewReader([]byte(`hoge`)),
		ShowReqBody: true,
		ShowResBody: false,
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	g := NewGirl()

	req, err := http.NewRequest(g.Method, g.URL, g.Body)
	if err != nil {
		return err
	}

	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	reqDump, err := httputil.DumpRequestOut(req, g.ShowReqBody)
	if err != nil {
		return err
	}

	fmt.Println(string(reqDump))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resDump, err := httputil.DumpResponse(res, g.ShowResBody)
	if err != nil {
		return err
	}

	fmt.Println(string(resDump))
	return nil
}
