package usercrudhandler

import (
	"fmt"
	"net/http"

	dbrepo "github.com/gohttpexamples/restaurant/dao/dbrepository"
	"github.com/gohttpexamples/restaurant/dbrepo/userrepo"
	"github.com/gohttpexamples/restaurant/delivery/restapplication/packages/httphandlers"
	mthdroutr "github.com/gohttpexamples/restaurant/delivery/restapplication/packages/mthdrouter"
	"github.com/gohttpexamples/restaurant/delivery/restapplication/packages/resputl"
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
	fmt.Println("get")
	return resputl.Response200OK("im working")
	/*
		pathParam := mux.Vars(r)
		usID := pathParam["id"]
		if usID == "" {

			//return resputl.Response200OK(generateSampleResponseObj())
			resp, err := p.usersvc.GetAll()

			if err != nil {
				return resputl.ReponseCustomError(err)
			}

			responseObj := transformobjListToResponse(resp)

			return resputl.Response200OK(responseObj)
		} else {
			obj, err := p.usersvc.GetByID(usID)

			if err != nil {
				return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
			}

			userObj := UserGetRespDTO{
				ID:        obj.ID,
				FirstName: obj.Firstname,
				LastName:  obj.Lastname,
				CreatedOn: obj.CreatedOn,
				Age:       obj.Age,
			}

			return resputl.Response200OK(userObj)

		}
	*/
}

/*
//Post method creates new temporary schedule
func (p *UserCrudHandler) Post(r *http.Request) resputl.SrvcRes {
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
	var requestdata *UserCreateReqDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	f := userrepo.Factory{}
	userObj := f.NewUser(requestdata.FirstName, requestdata.LastName, requestdata.Age)
	id, err := p.usersvc.Create(userObj)
	if err != nil {
		logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK(&UserCreateRespDTO{ID: id})
}

//Put method modifies temporary schedule contents
func (p *UserCrudHandler) Put(r *http.Request) resputl.SrvcRes {

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
	var requestdata *domain.User
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	//f := userrepo.Factory{}
	//userObj := f.NewUser(requestdata.FirstName, requestdata.LastName, requestdata.Age)
	err = p.usersvc.Update(requestdata)
	if err != nil {
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error Finding Data"), "")
	}

	return resputl.Response200OK("Updated Succefully")
}

//Delete method removes temporary schedule from db
func (p *UserCrudHandler) Delete(r *http.Request) resputl.SrvcRes {

	pathParam := mux.Vars(r)
	usID := pathParam["id"]
	if usID == "" {

		return resputl.Response200OK("give User ID")
	} else {
		err := p.usersvc.Delete(usID)

		if err != nil {
			return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
		}

		return resputl.Response200OK("DELETE Succefully")
	}
}
*/
