package store

import (
    s3 "github.com/luhhujbb/goinventory/store/nonatomic/s3store"
    "github.com/luhhujbb/goinventory/ivtype"
)

func loadFromAtomic(store ivtype.Store)(interface{},error){
    return nil, nil
}

func loadFromNonAtomic(store ivtype.Store)(interface{},error){
    switch store.Type {
    case "s3": return s3.Load(store)
    }
    return nil, nil
}

func LoadFromStore(store ivtype.Store) (interface{},error){
    if store.Atomic {
        return loadFromAtomic(store)
    } else {
        return loadFromAtomic(store)
    }
}
