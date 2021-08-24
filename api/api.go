package goinventory

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Server() {
    router := httprouter.New()
    router.GET("/", index)
    router.GET("/hello/:name", hello)

    http.ListenAndServe(":8080", router)
}
