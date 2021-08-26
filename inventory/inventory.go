package goinventory

import (
    "github.com/luhhujbb/goinventory/ivtype"
    "github.com/luhhujbb/goinventory/store"
)

/* main inventory declaration */
var inventory map[string]map[string]string

func loadInventory (stores []ivtype.Store){
    for _ , st := range stores {
        ivFromStore, err := store.LoadFromStore(st)
        if err != nil {
            inventory = ivFromStore.(map[string]map[string]string)
            break
        }
    }
}

func GetResource(id string) map[string]string{
    return inventory[id]
}
