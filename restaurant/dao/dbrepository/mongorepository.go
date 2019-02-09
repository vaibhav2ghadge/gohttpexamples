package dbrepository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	domain "github.com/gohttpexamples/restaurant/dao/domain"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "restaurant"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

//Find a Restaurant by Id
func (r *MongoRepository) Get(id domain.ID) (*domain.Restaurant, error) {
	result := domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//get all restaurant

func (r *MongoRepository) GetAll() ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

// find by name substring case insentive
func (r *MongoRepository) FindByName(name string) ([]*domain.Restaurant, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	result := []*domain.Restaurant{}
	//	name = "/"+name+"/"
	err := coll.Find(bson.M{"name": bson.RegEx{name, "i"}}).All(&result)
	//return result,err
	if err != nil {
		return result, err
	}
	return result, nil
}

//search in allfield
func (r *MongoRepository) Search(query string) ([]*domain.Restaurant, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	result := []*domain.Restaurant{}
	//	name = "/"+name+"/"
	err := coll.Find(bson.M{"$or": []bson.M{bson.M{"name": bson.RegEx{query, "i"}}, bson.M{"address": bson.RegEx{query, "i"}}, bson.M{"addressLine2": bson.RegEx{query, "i"}}, bson.M{"url": bson.RegEx{query, "i"}}, bson.M{"outcode": bson.RegEx{query, "i"}}, bson.M{"postcode": bson.RegEx{query, "i"}}, bson.M{"type_of_food": bson.RegEx{query, "i"}}}}).All(&result)
	//return result,err
	if err != nil {
		return result, err
	}
	return result, nil
}

//Store a Restaurantrecord
func (r *MongoRepository) Store(b *domain.Restaurant) (domain.ID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ID(0) == b.DBID {
		b.DBID = domain.NewID()
	}

	_, err := coll.UpsertId(b.DBID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.DBID, nil
}

//delete document from mongodb by id

func (r *MongoRepository) Delete(id domain.ID) error {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Remove(bson.M{"_id": id})
	return err

}
func (r *MongoRepository) CountByTypeOfFood(foodType string) (int, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	count, err := coll.Find(bson.M{"type_of_food": foodType}).Count()
	switch err {
	case nil:
		return count, nil
	case mgo.ErrNotFound:
		return 0, domain.ErrNotFound
	default:
		return 0, err
	}
}

//return all restaurant info that match to typeoffood

func (r *MongoRepository) FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error) {

	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"type_of_food": foodType}).All(&result)
	fmt.Println("result", result[0])
	//err := coll.Find({Postcode:Postcode}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//name of all matching restaurant by post code
func (r *MongoRepository) FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error) {

	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"postcode": postCode}).All(&result)
	fmt.Println("result", result[0])
	//err := coll.Find({Postcode:Postcode}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//print restaurant
func PrintRestaurant(r []*domain.Restaurant) {
	for _, obj := range r {
		fmt.Println(obj)
	}
}
func (r *MongoRepository) MarshalingFileData(filePath string) int {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var cnt int
	var data domain.Restaurant
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p := []byte(scanner.Text())
		json.Unmarshal(p, &data)

		data.DBID = domain.NewID()
		did, _ := r.Store(&data)
		if did == domain.ID(0) {
			fmt.Println("Error in Insert")
			break
		} else {
			cnt = cnt + 1
		}
	}
	return cnt
}
