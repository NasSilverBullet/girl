package main

import (
	"fmt"
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

func run() error {
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		return err
	}

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return err
	}

	fmt.Println(string(reqDump))

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resDump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return err
	}

	fmt.Println(string(resDump))
	return nil
}
