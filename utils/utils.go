package utils

func InterfaceToIDDictViper(input interface{}) map[string]map[string]string {
    output := make(map[string]map[string]string)
    for key,val := range input.(map[string]interface{}){
            tval := make(map[string]string)
            for k,v := range val.(map[string]interface{}){
                tval[k] = v.(string)
            }
            output[key] = tval
    }
    return output
}

func InterfaceToIDDictYaml(input interface{}) map[string]map[string]interface{} {
    output := make(map[string]map[string]interface{})
    for key,val := range input.(map[string]interface{}){
            tval := make(map[string]interface{})
            for k,v := range val.(map[interface{}]interface{}){
                    tval[k.(string)] = v
            }
            output[key] = tval
    }
    return output
}
