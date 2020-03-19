package repositiory

import (
	"time"

	"github.com/globalsign/mgo/bson"

	"MyBlog/helper"
	"MyBlog/mongo"
	"MyBlog/server/models"
)

func SignUp(user map[string]interface{}) error {
	C := mongo.MDC("blogs", "users")

	return C.Insert(bson.M(user))
}

func GetUserByPhone(phone string) (*models.User, error) {
	C := mongo.MDC("blogs", "users")
	query := map[string]interface{}{"phone": phone}
	user := &models.UserMongo{}
	err := C.Find(bson.M(query)).One(user)
	if err != nil {
		return nil, err
	}

	return models.TransMongoUserToServerUser(user), nil
}

func GetUserByUuid(uuid string) (*models.User, error) {
	C := mongo.MDC("blogs", "users")
	query := map[string]interface{}{"uuid": uuid}
	user := &models.UserMongo{}
	err := C.Find(bson.M(query)).One(user)
	if err != nil {
		return nil, err
	}

	return models.TransMongoUserToServerUser(user), nil
}

func UpdatePassword(uuid, passwordHash string) error {

	C := mongo.MDC("blogs", "users")
	selector := map[string]interface{}{"uuid": uuid}
	update := map[string]interface{}{
		"password_hash": passwordHash,
		"update_at":     helper.TimeToTimestamp(time.Now()),
	}
	return C.Update(selector, update)
}
