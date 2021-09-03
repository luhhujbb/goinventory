package aws

import (
    "github.com/luhhujbb/goinventory/ivtype"
)

const RDSDefaultKey = "rds"

var rdsstores []ivtype.Store

func ConfigureRDSStores(stores *[]ivtype.Store){
    rdsstores = *stores
}
