package usercrudhandler

type RestaurantRespDTO struct {
	DBID         string  `json:"id" bson:"_id"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"addressLine2" bson:"addressLine2"`
	URL          string  `json:"url" bson:"url"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	Type_of_food string  `json:"type_of_food" bson:"type_of_food"`
}

type RestaurantGetListRespDTO struct {
	Rest  []RestaurantRespDTO `json:"users"`
	Count int                 `json:"count"`
}
