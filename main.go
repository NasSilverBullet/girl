package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/NasSilverBullet/girl/pkg/girl"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	g := girl.New(
		"https://example.com",
		http.MethodGet,
		map[string]string{"Content-Type": "application/json", "Authorization": "Bearer access-token", "tenant-id": "testid", "application-key": "testkey"},
		bytes.NewReader([]byte(`hoge`)),
		true,
		false,
	)

	reqDump, resDump, err := g.Request()
	if err != nil {
		return err
	}

	fmt.Printf("Request Dump:\n>>>>>%s<<<<<\n\n", string(reqDump))
	fmt.Printf("Response Dump:\n>>>>>%s<<<<<\n", string(resDump))
	return nil
}
