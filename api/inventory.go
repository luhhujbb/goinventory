package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/luhhujbb/goinventory/inventory"
    "encoding/json"
)



func GetInventoryResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    resource := *inventory.GetResource(ps.ByName("id"))
    response := ApiResponse{State: "success", Data: resource}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
