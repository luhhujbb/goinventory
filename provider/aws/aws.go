package aws

import (
    "context"
    "log"
    "github.com/luhhujbb/goinventory/ivtype"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
    "bytes"
)

const AWSDefaultKey = "aws"

func defaultKey(sdkey string) string {
    return AWSDefaultKey + "/" + sdkey
}

func ConfigureAWS(config interface{}){
    tconf := config.(map[string]interface{})
    tsto := utils.InterfaceToIDDictViper(tconf["store"])
    //initiate store
    ec2Stores = store.ConfToStore(&tsto, defaultKey(EC2DefaultKey))
    rdsStores = store.ConfToStore(&tsto, defaultKey(RDSDefaultKey))
    elasticacheStores = store.ConfToStore(&tsto, defaultKey(ElasticacheDefaultKey))
    log.Print(stores)
    loadInventory()
}
