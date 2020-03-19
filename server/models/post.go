package models

type Article struct {
	UUID       string `bson:"uuid"`
	AuthorUuid string `bson:"author_uuid"`
	Title      string `bson:"title"`
	Content    string `bson:"content"`
	CreatedAt  int64  `bson:"created_at"`
	UpdatedAt  int64  `bson:"updated_at"`
}

type ArticleInMongo struct {
	UUID       string `bson:"uuid"`
	AuthorUuid string `bson:"author_uuid"`
	Title      string `bson:"title"`
	Content    string `bson:"content"`
	CreatedAt  int64  `bson:"created_at"`
	UpdatedAt  int64  `bson:"updated_at"`
}
