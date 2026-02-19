package spotify

import (
	"fmt"
	"net/http"
)

var AuthCodeChan = make(chan string)

func callbackHandler(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	code := q.Get("code")

	AuthCodeChan <- code

	fmt.Println("Login success")
}

func StartCallbackServer() {
	http.HandleFunc("/callback", callbackHandler)
	http.ListenAndServe(":8080", nil)
}
