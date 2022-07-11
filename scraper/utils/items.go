package utils

//type RealEstateItem struct {
//	Meta    interface{} `json:"meta"`
//	Contact interface{} `json:"contact"`
//}
type RealEstateItem struct {
	Meta         interface{} `json:"meta"`
	Contact      interface{} `json:"contact"`
	Address      string      `json:"address"`
	PropertyInfo interface{} `json:"propertyInfo"`
}
