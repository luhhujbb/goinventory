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
    Value []string `json:"value"`
    Not bool `json:"not,omitempty"`
}

/* main inventory declaration */
var inventory map[string]map[string]interface{}
var stores []ivtype.Store
const defaultKey = "resources.yml"

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

func GetResource(id string) *map[string]interface{}{
    resource := inventory[id]
    return &resource
}

func GetInventory() *map[string]map[string]interface{}{
    return &inventory
}

func GetFilteredInventory(tagFilter TagFilter) *map[string]map[string]interface{}{
    return &inventory
}

func FastSync(){
    loadInventory()
}

func ConfigureInventory(config interface{}){
    tconf := config.(map[string]interface{})
    tsto := utils.InterfaceToIDDictViper(tconf["store"])
    stores = store.ConfToStore(&tsto, defaultKey)
    log.Print(stores)
    loadInventory()
}
