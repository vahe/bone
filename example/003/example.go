package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
)

var (
	router = bone.New()
)

func Generate() *bone.Mux {
	muxx := bone.New()

	muxx.GetFunc("*/test/:val", TestHandler)

	return muxx
}

func main() {
	router.Handle("/index/*", Generate())

	http.ListenAndServe(":8080", router)
}

func TestHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RequestURI)
	fmt.Println(bone.GetValue(req, "val"))
	rw.Write([]byte(bone.GetValue(req, "val")))
}
