package goinventory

import (
    store "github.com/luhhujbb/goinventory/store"
)

/* main inventory declaration */
var inventory map[string]map[string]string

func loadInventory (stores []store.Store){

}

func GetResource(id string) map[string]string{
    return inventory[id]
}
