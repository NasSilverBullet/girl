package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

type Girl struct {
	URL         string
	Method      string
	Body        io.Reader
	ShowReqBody bool
	ShowResBody bool
}

func run() error {
	g := Girl{
		URL:         "https://google.com",
		Method:      http.MethodGet,
		ShowReqBody: true,
		ShowResBody: false,
	}

	req, err := http.NewRequest(g.Method, g.URL, g.Body)
	if err != nil {
		return err
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
