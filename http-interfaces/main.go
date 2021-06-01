package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

//read html from any page!
func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Writer Interface, Reader Iterface
	lw := logWriter{}

	io.Copy(lw, resp.Body)

	//make takes a type and size
	// 	bs := make([]byte, 99999)
	// 	resp.Body.Read(bs)
	// 	fmt.Println(string(bs))
	// }
}

//custom Writer!
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes -> ", len(bs))
	return len(bs), nil
}
