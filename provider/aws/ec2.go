package aws

import (
    "github.com/luhhujbb/goinventory/ivtype"
)

const EC2DefaultKey = "ec2"

var ec2stores []ivtype.Store

func ConfigureEC2Stores(stores *[]ivtype.Store){
    ec2stores = *stores
}
