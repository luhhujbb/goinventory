package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/luhhujbb/goinventory/inventory"
    "encoding/json"
    "log"
)

//Entities Getter

func makeEntitityGetter(entityType string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var entity map[string]string
        switch entityType {
            case "resource": entity = *inventory.GetResource(ps.ByName("id"))
            case "group": entity = *inventory.GetGroup(ps.ByName("id"))
            case "alias": entity = *inventory.GetAlias(ps.ByName("id"))
        }
        response := ApiResponse{State: "success", Data: entity}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func makeEntitiesGetter(entityType string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var entities map[string]map[string]string
        switch entityType {
            case "resource": entities = *inventory.GetResources()
            case "group": entities = *inventory.GetGroups()
            case "alias": entities = *inventory.GetAliases()
        }
        response := ApiResponse{State: "success", Data: entities}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func makeFilteredEntitiesGetter(entityType string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var tagFilter inventory.TagFilter
        var response ApiResponse
        err := json.NewDecoder(r.Body).Decode(&tagFilter)
        if err != nil {
            log.Print(err)
            getInventory := makeEntitiesGetter(entityType)
            getInventory(w,r,ps)
        } else {
            var entities map[string]map[string]string
            switch entityType {
                case "resource": entities = *inventory.GetResources()
                case "group": entities = *inventory.GetGroups()
                case "alias": entities = *inventory.GetAliases()
            }
            response = ApiResponse{State: "success", Data: entities}
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
        }
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
    router.GET("/inventory/resource/:id", makeEntitityGetter("resource"))
    router.GET("/inventory/resource", makeEntitiesGetter("resource"))
    router.POST("/inventory/resource", makeFilteredEntitiesGetter("resource"))
}
