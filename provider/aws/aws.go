package aws

import (
    "github.com/luhhujbb/goinventory/utils"
    "github.com/luhhujbb/goinventory/store"
)

const AWSDefaultKey = "aws"

func defaultKey(sdkey string) string {
    return AWSDefaultKey + "/" + sdkey
}

func loadInventory(){

}

func ConfigureAWS(config interface{}){
    tconf := config.(map[string]interface{})
    tsto := utils.InterfaceToIDDictViper(tconf["store"])
    //initiate store
    ec2Stores := store.ConfToStore(&tsto, defaultKey(EC2DefaultKey))
    rdsStores := store.ConfToStore(&tsto, defaultKey(RDSDefaultKey))
    elasticacheStores := store.ConfToStore(&tsto, defaultKey(ElasticacheDefaultKey))
    ConfigureEC2Stores(&ec2Stores)
    ConfigureRDSStores(&rdsStores)
    ConfigureECaStores(&elasticacheStores)
    loadInventory()
}
