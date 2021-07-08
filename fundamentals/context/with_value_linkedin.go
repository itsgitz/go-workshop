// This source code is the example of WithTimeout usage insipired by post from linkedin
// https://www.linkedin.com/pulse/its-time-understand-golang-contexts-lucas-schenkel-schieferdecker/
package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(testHandler)
	http.Handle("/test", addContextMiddleware(handler))

	http.ListenAndServe(":8000", nil)
}

func addContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create child context with value using parent context from r.Context()
		ctx := context.WithValue(r.Context(), "Name", "Ryouko")

		// change the default request to request with context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Context:", r.Context().Value("Name"))
}
