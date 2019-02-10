package usercrudhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	logger "log"
	"net/http"

	dbrepo "github.com/gohttpexamples/restaurant/dao/dbrepository"
	domain "github.com/gohttpexamples/restaurant/dao/domain"
	"github.com/gohttpexamples/restaurant/dbrepo/userrepo"
	customerrors "github.com/gohttpexamples/restaurant/delivery/restapplication/packages/errors"
	"github.com/gohttpexamples/restaurant/delivery/restapplication/packages/httphandlers"
	mthdroutr "github.com/gohttpexamples/restaurant/delivery/restapplication/packages/mthdrouter"
	"github.com/gohttpexamples/restaurant/delivery/restapplication/packages/resputl"
	"github.com/gorilla/mux"
)

type UserCrudHandlerr struct {
	httphandlers.BaseHandler
	usersvc userrepo.Repository
}
type RestCrudHandler struct {
	Mongo1 *dbrepo.MongoRepository
	httphandlers.BaseHandler
}

func NewRestCrudHandler(mongor *dbrepo.MongoRepository) *RestCrudHandler {
	return &RestCrudHandler{Mongo1: mongor}
}

func (p *RestCrudHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

//Get http method to get data
func (p *RestCrudHandler) Get(r *http.Request) resputl.SrvcRes {
	//pathParam := mux.Vars(r)
	//usID := pathParam["typeOfFood"]
	foodType := r.URL.Query()["typeOfFood"]
	fmt.Println("get", foodType)
	if len(foodType) == 1 {
		result, _ := p.Mongo1.FindByTypeOfFood(foodType[0])
		//fmt.Print(err, result)
		responseObj := transformobjListToResponse(result)

		return resputl.Response200OK(responseObj)

	}
	nameOfFood := r.URL.Query()["name"]
	fmt.Println("get", nameOfFood)
	if len(nameOfFood) == 1 {
		result, _ := p.Mongo1.FindByName(nameOfFood[0])
		//fmt.Print(err, result)
		responseObj := transformobjListToResponse(result)

		return resputl.Response200OK(responseObj)

	}
	searchQry := r.URL.Query()["searchTerm"]
	fmt.Println("get", searchQry)
	if len(searchQry) == 1 {
		result, _ := p.Mongo1.Search(nameOfFood[0])
		//fmt.Print(err, result)
		responseObj := transformobjListToResponse(result)

		return resputl.Response200OK(responseObj)

	}
	pathParam := mux.Vars(r)
	usID := pathParam["id"]
	fmt.Println(usID)
	if usID == "" {

		//return resputl.Response200OK(generateSampleResponseObj())
		resp, err := p.Mongo1.GetAll()

		if err != nil {
			return resputl.ReponseCustomError(err)
		}
		responseObj := transformobjListToResponse(resp)
		return resputl.Response200OK(responseObj)
	} else {
		resp, err := p.Mongo1.Get(usID)

		if err != nil {
			return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
		}

		return resputl.Response200OK(resp)
	}

	return resputl.Response200OK("im working")
}

//Post method creates new temporary schedule
func (p *RestCrudHandler) Post(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateUserCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var requestdata *domain.Restaurant
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	id, err := p.Mongo1.Store(requestdata)
	if err != nil {
		logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK(&RestaurantRespDTO{DBID: id})
}

//Put method modifies temporary schedule contents
func (p *RestCrudHandler) Put(r *http.Request) resputl.SrvcRes {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateUserCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}
	logger.Printf("Received PUT request to UPDATE schedule %s ", string(body))
	var requestdata *domain.Restaurant
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	//f := userrepo.Factory{}
	//userObj := f.NewUser(requestdata.FirstName, requestdata.LastName, requestdata.Age)
	err = p.Mongo1.Update(requestdata)
	fmt.Println(err)
	if err != nil {
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error Finding Data"), "")
	}

	return resputl.Response200OK("Updated Succefully")
}

//Delete method removes temporary schedule from db
func (p *RestCrudHandler) Delete(r *http.Request) resputl.SrvcRes {

	pathParam := mux.Vars(r)
	usID := pathParam["id"]
	if usID == "" {

		return resputl.Response200OK("give User ID")
	} else {
		err := p.Mongo1.Delete(usID)

		if err != nil {
			return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
		}

		return resputl.Response200OK("DELETE Succefully")
	}
}
