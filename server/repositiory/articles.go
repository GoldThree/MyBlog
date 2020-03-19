package repositiory

import (
	"github.com/globalsign/mgo/bson"

	"MyBlog/mongo"
	"MyBlog/server/models"
)

func Find(offset, limit int) ([]models.Article, error) {

	C := mongo.MDC("blogs", "articles")
	as := make([]models.Article, 0)
	err := C.Find(bson.M(map[string]interface{}{})).Skip(offset).Limit(limit).All(&as)
	if err != nil {
		return nil, err
	}

	return as, nil
}

func FindByUuid(uuid string, offset, limit int) ([]models.Article, error) {

	C := mongo.MDC("blogs", "articles")
	as := make([]models.Article, 0)
	err := C.Find(bson.M(map[string]interface{}{"uuid": uuid})).Skip(offset).Limit(limit).All(&as)
	if err != nil {
		return nil, err
	}

	return as, nil
}

func Post(data map[string]interface{}) error {

	C := mongo.MDC("blogs", "articles")
	return C.Insert(bson.M(data))
}

func GetArticle(uuid string) (models.Article, error) {

	C := mongo.MDC("blogs", "articles")
	query := map[string]interface{}{"uuid": uuid}

	result := models.Article{}
	err := C.Find(bson.M(query)).One(&result)
	if err != nil {
		return models.Article{}, err
	}
	return result, nil
}

func DeleteArticle(uuid string) error {

	C := mongo.MDC("blogs", "articles")
	query := map[string]interface{}{"uuid": uuid}

	return C.Remove(bson.M(query))
}
