package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/luhhujbb/goinventory/inventory"
    "encoding/json"
)



func getInventoryResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    resource := *inventory.GetResource(ps.ByName("id"))
    response := ApiResponse{State: "success", Data: resource}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func getInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    iv := *inventory.GetInventory()
    response := ApiResponse{State: "success", Data: iv}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func getFilteredInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var tagFilter inventory.TagFilter
    var response ApiResponse
    err := json.NewDecoder(r.Body).Decode(&tagFilter)
    if err != nil {
        response = ApiResponse{State: "error", Data: nil, Message: "no tags defined for filtering"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
        http.Error(w, err.Error(), http.StatusBadRequest)
    } else {
        iv := *inventory.GetFilteredInventory(tagFilter)
        response = ApiResponse{State: "success", Data: iv}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }

}

func fastSync(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    inventory.FastSync()
    response := ApiResponse{State: "success", Data: "", Message: "Fsync submitted"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}


func ConfigureInventoryRoute(router *httprouter.Router){
    router.GET("/inventory/fsync", fastSync)
    router.GET("/inventory/resource/:id", getInventoryResource)
    router.GET("/inventory/resource", getInventory)
    router.POST("/inventory/resource", getFilteredInventory)
}
