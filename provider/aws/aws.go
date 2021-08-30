package aws




func ConfigureAWS(config interface{}){
    tconf := config.(map[string]interface{})
    tsto := utils.InterfaceToIDDictViper(tconf["store"])
    ec2stores = store.ConfToStore(&tsto, EC2DefaultKey)
    log.Print(stores)
    loadInventory()
}
