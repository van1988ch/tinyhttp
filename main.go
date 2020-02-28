package main

import (
	"fmt"
	"net/http"
	"tinyhttp/tinyhttp"
)


func main()  {
	engine := tinyhttp.New()
	engine.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k,v := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	})

	engine.Run(":9999")
}
