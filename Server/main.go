package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	fmt.Println("get transaction")
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("test")
		fmt.Fprint(writer, "Hi my client")
	case <-ctx.Done():
		fmt.Println("Time out ==> ", ctx.Err())
		http.Error(writer, ctx.Err().Error(), http.StatusRequestTimeout)
	}

}
