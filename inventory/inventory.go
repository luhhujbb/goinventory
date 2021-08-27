package inventory

import (
    "github.com/luhhujbb/goinventory/ivtype"
    "github.com/luhhujbb/goinventory/store"
    "github.com/luhhujbb/goinventory/utils"
    "log"
)

/* main inventory declaration */
var inventory map[string]map[string]string
var stores []ivtype.Store
const defaultKey = "resources.yml"

func loadInventory (){
    for _ , st := range stores {
        ivFromStore, err := store.LoadFromStore(st)
        if err != nil {
            //need to add payload decoder whent format is != from ""
            inventory = utils.InterfaceToIDDict(ivFromStore)
            break
        }
    }
}

func GetResource(id string) *map[string]string{
    resource := inventory[id]
    return &resource
}

func GetInventory() *map[string]map[string]string{
    return &inventory
}

func ConfigureInventory(config interface{}){
    tconf := config.(map[string]interface{})
    tsto := utils.InterfaceToIDDict(tconf["store"])
    stores = store.ConfToStore(&tsto, defaultKey)
    log.Print(stores)
    loadInventory()
}
