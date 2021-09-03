package aws

import (
    "github.com/luhhujbb/goinventory/ivtype"
)

const ElasticacheDefaultKey = "elasticache"

var ecastores []ivtype.Store

func ConfigureECaStores(stores *[]ivtype.Store){
    ecastores = *stores
}
