package main

import (
	"fmt"
	"math"
	"net/http"
)

func msg() string {
	return "<b>Code.education Rocks!</b>"
}

func loop(writer http.ResponseWriter, reqest *http.Request) {
	x := 0.0001
	for i := 0; i < 1000000000; i++ {
		x += math.Sqrt(x)
	}
	fmt.Println(msg())
}

func main() {
	http.HandleFunc("/", loop)
	http.ListenAndServe(":8000", nil)
}
