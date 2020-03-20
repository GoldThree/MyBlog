package server

import (
	"MyBlog/helper"
	"MyBlog/server/models"
	"MyBlog/server/repositiory"
	"errors"
	"time"
)

func Find(offset, limit int, uuid, token string) ([]models.Article, error) {

	isTrueToken, err := checkToken(uuid, token)
	if err != nil {
		return nil, err
	}

	if !isTrueToken {
		return nil, errors.New("error_token")
	}

	articles, err := repositiory.Find(offset, limit)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func FindByUuid(uuid, token string, offset, limit int) ([]models.Article, error) {

	isTrueToken, err := checkToken(uuid, token)
	if err != nil {
		return nil, err
	}

	if !isTrueToken {
		return nil, errors.New("error_token")
	}

	articles, err := repositiory.FindByUuid(uuid, offset, limit)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func Post(title, content, authorUuid, token string) error {

	isTrueToken, err := checkToken(authorUuid, token)
	if err != nil {
		return err
	}

	if !isTrueToken {
		return errors.New("error_token")
	}

	timestamp := helper.TimeToTimestamp(time.Now())

	data := map[string]interface{}{
		"uuid":        helper.NewV4(),
		"author_uuid": authorUuid,
		"title":       title,
		"Content":     content,
		"created_at":  timestamp,
		"updated_at":  timestamp,
	}

	err = repositiory.Post(data)
	if err != nil {
		return err
	}
	return nil
}

func GetArticle(uuid, token string) (models.Article, error) {

	isTrueToken, err := checkToken(uuid, token)
	if err != nil {
		return models.Article{}, err
	}

	if !isTrueToken {
		return models.Article{}, errors.New("error_token")
	}

	article, err := repositiory.GetArticle(uuid)
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func DeleteArticle(uuid, token string) error {

	isTrueToken, err := checkToken(uuid, token)
	if err != nil {
		return err
	}

	if !isTrueToken {
		return errors.New("error_token")
	}

	return repositiory.DeleteArticle(uuid)
}
