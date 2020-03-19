package server

import (
	"MyBlog/helper"
	"MyBlog/server/models"
	"MyBlog/server/repositiory"
	"time"
)

func Find(offset, limit int) ([]models.Article, error) {

	articles, err := repositiory.Find(offset, limit)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func FindByUuid(uuid string, offset, limit int) ([]models.Article, error) {

	articles, err := repositiory.FindByUuid(uuid, offset, limit)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func Post(title, content, authorUuid string) error {

	timestamp := helper.TimeToTimestamp(time.Now())

	data := map[string]interface{}{
		"uuid":        helper.NewV4(),
		"author_uuid": authorUuid,
		"title":       title,
		"Content":     content,
		"created_at":  timestamp,
		"updated_at":  timestamp,
	}

	err := repositiory.Post(data)
	if err != nil {
		return err
	}
	return nil
}

func GetArticle(uuid string) (models.Article, error) {

	article, err := repositiory.GetArticle(uuid)
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func DeleteArticle(uuid string) error {

	return repositiory.DeleteArticle(uuid)
}
