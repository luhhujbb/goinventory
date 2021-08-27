package api

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/luhhujbb/goinventory/inventory"
    "github.com/luhhujbb/goinventory/ivtype"
    "encoding/json"
)



func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func getInventoryResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    resource := *inventory.GetResource(ps.ByName("id"))
    response := ivtype.ApiResponse{State: "success", Data: resource}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}


func Server() {
    router := httprouter.New()
    router.GET("/inventory/resource/:id", getInventoryResource)
    router.GET("/", index)
    router.GET("/hello/:name", hello)

    http.ListenAndServe(":8080", router)
}
