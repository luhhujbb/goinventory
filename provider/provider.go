package provider

import (
    store "github.com/luhhujbb/goinventory/store"
)

type Provider struct {
    Stores []store.Store
    Region []strings
}
