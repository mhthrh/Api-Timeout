package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/", nil)

	if err != nil {
		fmt.Println(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)

}
