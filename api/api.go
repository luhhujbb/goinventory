package api

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

type ApiResponse struct {
    State string "json: state"
    Data interface{} "json: data"
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome To Inventory!\n")
}

func Server() {
    router := httprouter.New()
    router.GET("/inventory/resource/:id", GetInventoryResource)
    router.GET("/", index)
    http.ListenAndServe(":8080", router)
}
