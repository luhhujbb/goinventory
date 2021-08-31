package inventory

import (
    "github.com/luhhujbb/goinventory/ivtype"
    "github.com/luhhujbb/goinventory/store"
    "github.com/luhhujbb/goinventory/utils"
    "log"
)

type TagFilter struct {
    Tags []Tag `json:"tags"`
}

type Tag struct {
    Name string `json:"name"`
    Value interface{} `json:"value"`
    Not bool `json:"not,omitempty"`
}

/* main inventory declaration */
var inventory map[string]map[string]string
var stores []ivtype.Store
const defaultKey = "resources"

func loadInventory (){
    for _ , st := range stores {
        ivFromStore, err := store.LoadFromStore(st)
        if err == nil {
            inventory = utils.InterfaceToIDDictYaml(ivFromStore)
            break
        } else {
            log.Print(err)
        }
    }
}

func tagMatcher (entity map[string]string, tag Tag) bool {
    var match bool
    switch tag.Value.(type) {
        case string: match = (tag.Value.(string) == entity[tag.Name])
        case []interface{}: match = utils.ContainsIntoAI(tag.Value.([]interface{}),entity[tag.Name])
    }
    if tag.Not {
        return !match
    } else {
        return match
    }
}

func tagsMatcher (entity map[string]string, tags []Tag) bool {
    for _,tag := range tags {
        if !tagMatcher(entity,tag){
            return false
        }
    }
    return true
}

func GetResource(id string) *map[string]string{
    resource := inventory[id]
    return &resource
}

func GetInventory() *map[string]map[string]string{
    return &inventory
}

func GetFilteredInventory(tagFilter TagFilter) *map[string]map[string]string{
    filteredInventory := make(map[string]map[string]string)
    for k,v := range inventory {
        if tagsMatcher(v,tagFilter.Tags) {
            filteredInventory[k] = v
        }
    }
    return &filteredInventory
}

func FastSync(){
    go loadInventory()
}

func ConfigureInventory(config interface{}){
    tconf := config.(map[string]interface{})
    tsto := utils.InterfaceToIDDictViper(tconf["store"])
    stores = store.ConfToStore(&tsto, defaultKey)
    log.Print(stores)
    loadInventory()
    log.Print("Inventory loaded")
}
