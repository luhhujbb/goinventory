package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
)

type ApiResponse struct {
    State string `json:"state"`
    Data interface{} `json:"data"`
    Message string `json:"msg,omitempty"`
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    response := ApiResponse{State: "success", Data: "", Message:"Welcome To Inventory!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func Server() {
    router := httprouter.New()
    ConfigureInventoryRoute(router)
    router.GET("/", index)
    http.ListenAndServe(":8080", router)
}
