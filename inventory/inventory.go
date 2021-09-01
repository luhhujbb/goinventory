package inventory

import (
    "github.com/luhhujbb/goinventory/ivtype"
    "github.com/luhhujbb/goinventory/store"
    "github.com/luhhujbb/goinventory/utils"
    "log"
)

type TagFilter struct {
    Tags []Tag `json:"tags"`
}

type Tag struct {
    Name string `json:"name"`
    Value interface{} `json:"value"`
    Not bool `json:"not,omitempty"`
}

/* main inventory declaration */
var resources map[string]map[string]string
var groups map[string]map[string]string
var aliases map[string]map[string]string
var stores = make(map[string][]ivtype.Store)
const defaultResourceKey = "resources"
const defaultGroupKey = "groups"
const defaultAliasKey = "alias"

func loadInventory (){
    for entityType , ivstores := range stores{
        for _ , st := range ivstores {
            ivFromStore, err := store.LoadFromStore(st)
            if err == nil {
                switch entityType {
                    case "resource": resources = utils.InterfaceToIDDictYaml(ivFromStore)
                    case "group": groups = utils.InterfaceToIDDictYaml(ivFromStore)
                    case "alias": aliases = utils.InterfaceToIDDictYaml(ivFromStore)
                }
                break
                } else {
                    log.Print(err)
                }
            }
        }
    }

func tagMatcher (entity map[string]string, tag Tag) bool {
    var match bool
    switch tag.Value.(type) {
        case string: match = (tag.Value.(string) == entity[tag.Name])
        case []interface{}: match = utils.ContainsIntoAI(tag.Value.([]interface{}),entity[tag.Name])
    }
    if tag.Not {
        return !match
    } else {
        return match
    }
}

func tagsMatcher (entity map[string]string, tags []Tag) bool {
    for _,tag := range tags {
        if !tagMatcher(entity,tag){
            return false
        }
    }
    return true
}

//Filtered Entities
func getFilteredEntities(entities map[string]map[string]string,tagFilter TagFilter) *map[string]map[string]string{
    filteredEntities := make(map[string]map[string]string)
    for k,v := range entities {
        if tagsMatcher(v,tagFilter.Tags) {
            filteredEntities[k] = v
        }
    }
    return &filteredEntities
}

//Resource(s) Getter
func GetResource(id string) *map[string]string{
    resource := resources[id]
    return &resource
}

func GetResources() *map[string]map[string]string{
    return &resources
}

func GetFilteredResources(tagFilter TagFilter) *map[string]map[string]string{
    return getFilteredEntities(resources, tagFilter)
}

//Group(s) Getter
func GetGroup(id string) *map[string]string{
    group := groups[id]
    return &group
}

func GetGroups() *map[string]map[string]string{
    return &groups
}

func GetFilteredGroups(tagFilter TagFilter) *map[string]map[string]string{
    return getFilteredEntities(groups, tagFilter)
}

//Group(s) Getter
func GetAlias(id string) *map[string]string{
    alias := aliases[id]
    return &alias
}

func GetAliases() *map[string]map[string]string{
    return &aliases
}

func GetFilteredAliases(tagFilter TagFilter) *map[string]map[string]string{
    return getFilteredEntities(aliases, tagFilter)
}

func FastSync(){
    go loadInventory()
}

func ConfigureInventory(config interface{}){
    tconf := config.(map[string]interface{})
    tsto := utils.InterfaceToIDDictViper(tconf["store"])
    stores["resource"] = store.ConfToStore(&tsto,defaultResourceKey)
    stores["group"] = store.ConfToStore(&tsto,defaultGroupKey)
    stores["alias"] = store.ConfToStore(&tsto,defaultAliasKey)
    log.Print(stores)
    loadInventory()
    log.Print("Inventory loaded")
}
