package utils

import (
    "strconv"
)

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

func InterfaceToIDDictYaml(input interface{}) map[string]map[string]string {
    output := make(map[string]map[string]string)
    for key,val := range input.(map[string]interface{}){
            tval := make(map[string]string)
            for k,v := range val.(map[interface{}]interface{}){
                switch v.(type) {
                case string: tval[k.(string)] = v.(string)
                case int: tval[k.(string)] = strconv.Itoa(v.(int))
                }
            }
            output[key] = tval
    }
    return output
}

func ContainsIntoAI(input []interface{}, el string) bool {
    for _,val := range input {
        if el == val.(string) {
            return true
        }
    }
    return false
}

func Contains(input []string, el string) bool {
    for _,val := range input {
        if el == val {
            return true
        }
    }
    return false
}
