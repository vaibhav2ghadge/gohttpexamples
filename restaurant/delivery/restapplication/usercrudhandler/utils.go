package usercrudhandler

import (
	d "github.com/gohttpexamples/restaurant/dao/domain"
)

func transformobjListToResponse(resp []*d.Restaurant) RestaurantGetListRespDTO {
	responseObj := RestaurantGetListRespDTO{}
	for _, obj := range resp {
		userObj := RestaurantRespDTO{
			DBID:         obj.DBID,
			Name:         obj.Name,
			Address:      obj.Address,
			AddressLine2: obj.AddressLine2,
			URL:          obj.URL,
			Outcode:      obj.Outcode,
			Postcode:     obj.Postcode,
			Rating:       obj.Rating,
			Type_of_food: obj.Type_of_food,
		}
		responseObj.Rest = append(responseObj.Rest, userObj)
	}
	responseObj.Count = len(responseObj.Rest)

	return responseObj
}
