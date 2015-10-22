package utilities

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// Debugging functions for logging requests and responses

func DumpRequest(req *http.Request) {
	content, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println("Error dumping request: %v\n", err)
	} else {
		fmt.Println(string(content))
	}
}

func DumpResponse(req *http.Response) {
	content, err := httputil.DumpResponse(req, true)
	if err != nil {
		log.Println("Error dumping response: %v\n", err)
	} else {
		fmt.Println(string(content))
	}
}
