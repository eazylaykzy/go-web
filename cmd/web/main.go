package main

import (
	"fmt"
	handlers2 "github.com/eazylaykzy/go-web/pkg/handlers"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers2.Home)
	http.HandleFunc("/about", handlers2.About)

	fmt.Println(fmt.Sprintf("Server started on port:%s", port))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return
	}
}
