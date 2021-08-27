package store

import (
    s3 "github.com/luhhujbb/goinventory/store/nonatomic/s3store"
    "github.com/luhhujbb/goinventory/ivtype"
    "gopkg.in/yaml.v2"
    "errors"
)

func decode(format string,st string) (map[string]interface{},error){
    switch format {
    case "yaml": t := make(map[string]interface{})
                if err := yaml.Unmarshal([]byte(st), &t); err != nil {
                     return t, err
                 } else {
                     return t, nil
                 }
    }
    return make(map[string]interface{}), errors.New("unknown format")
}

func loadFromAtomic(store ivtype.Store)(map[string]interface{},error){
    return nil, nil
}

func loadFromNonAtomic(store ivtype.Store)(string,error){
    switch store.Type {
    case "s3": return s3.Load(store)
    }
    return "", nil
}

func LoadFromStore(store ivtype.Store) (interface{},error){
    if store.Atomic {
        return loadFromAtomic(store)
    } else {
        st, err := loadFromNonAtomic(store)
        if err != nil {
            return make(map[string]interface{}), err
        } else {
            return decode(store.Format,st)
        }
    }
}

func ConfToStore(storesspecs *map[string]map[string]string,defaultkey string) []ivtype.Store{
    resStore := make([]ivtype.Store,len(*storesspecs))
    idx := 0
    for storetype, storespec := range *storesspecs {
        var key string
        var format string
        if storespec["key"] == "" {
            key = defaultkey
        } else {
            key = storespec["key"]
        }
        if storespec["format"] == "" {
            format = "yaml"
        } else {
            format = storespec["format"]
        }
        switch storetype {
        case "s3":
            resStore[idx] = ivtype.Store{
                Key: key,
                Bucket: storespec["bucket"],
                Type: "s3",
                Atomic: false, // Indicate what is send
                Format: format,
            }
        case "file":
            resStore[idx] = ivtype.Store{
                Key: key,
                Bucket: storespec["bucket"],
                Type: "file",
                Atomic: false, // Indicate what is send
                Format: format,
            }
        }
        idx++
    }
    return resStore
}
