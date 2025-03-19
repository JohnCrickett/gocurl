package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func help() {
	println("gocurl: try gocurl <flags> [urls]")
}

func main() {
	method := flag.String("X", "GET", "HTTP method to use.")
	verbose := flag.Bool("v", false, "Enable verbose mode")
	data := flag.String("d", "this would be the data", "Input data")
	headers := flag.String("H", "additional headers", "Additional headers to send")

	flag.Parse()
	urls := flag.Args()

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DisableKeepAlives = true
	c := &http.Client{Transport: t}

	for _, uv := range urls {
		m := http.MethodGet

		switch {
		case *method == "POST":
			m = http.MethodPost
		case *method == "DELETE":
			m = http.MethodDelete
		case *method == "PUT":
			m = http.MethodPut
		}
		u, err := url.Parse(uv)
		if err != nil {
			panic(err)
		}

		var body io.Reader = nil
		if data != nil {
			body = strings.NewReader(*data)
		}
		req, err := http.NewRequest(m, uv, body)

		if *headers != "" {
			s := strings.Fields(*headers)
			if len(s) != 2 {
				fmt.Fprintf(os.Stderr, "invalid additional headers: %s\n", *headers)
			} else {
				req.Header.Set(s[0][:len(s[0])-1], s[1])
			}
		}

		req.Header.Set("Accept", "*/*")
		if err != nil {
			fmt.Printf("Could not create request: %s\n", err)
			os.Exit(1)
		}
		if *verbose {
			fmt.Printf("> %s %s HTTP/%d.%d\n", *method, u.Path, req.ProtoMajor, req.ProtoMinor)
			fmt.Println("> Host:", req.Host)
			fmt.Println("> Accept:", req.Header["Accept"][0])
			fmt.Println(">")
		}
		res, err := c.Do(req)
		if err != nil {
			fmt.Printf("client: error unable to make http request: %s\n", err)
			os.Exit(1)
		}
		if *verbose {
			fmt.Printf("< %s %s\n", res.Request.Proto, res.Status)
			for hdr, v := range res.Header {
				fmt.Printf("< %s: %s\n", hdr, v[0])
			}
		}

		if res.StatusCode == http.StatusOK {
			io.Copy(os.Stdout, res.Body)
		} else {
			fmt.Println(res.Status)
		}
	}
}
