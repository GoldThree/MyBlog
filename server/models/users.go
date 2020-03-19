package models

type User struct {
	UUID         string `json:"uuid"`
	UserName     string `json:"username"`
	Phone        string `json:"phone"`
	Gender       string `json:"gender"`
	PasswordSalt string `json:"password_salt"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type UserMongo struct {
	UUID         string `bson:"uuid"`
	UserName     string `bson:"username"`
	Phone        string `bson:"phone"`
	Gender       string `bson:"gender"`
	PasswordSalt string `bson:"password_salt"`
	PasswordHash string `bson:"password_hash"`
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

func TransMongoUserToServerUser(user *UserMongo) *User {

	return &User{
		UUID:         user.UUID,
		UserName:     user.UserName,
		Phone:        user.Phone,
		Gender:       user.Gender,
		PasswordSalt: user.PasswordSalt,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

}
