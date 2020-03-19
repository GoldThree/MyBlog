package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"MyBlog/helper"
	"MyBlog/server/repositiory"
)

func SignUp(username, gender, phone, password string) error {

	passwordSalt := helper.GenerateString(64)
	passwordHash := hashPassword(passwordSalt, password)

	timestamp := helper.TimeToTimestamp(time.Now())

	user := map[string]interface{}{
		"uuid":          helper.NewV4(),
		"username":      username,
		"gender":        gender,
		"phone":         phone,
		"password_salt": passwordSalt,
		"password_hash": passwordHash,
		"created_at":    timestamp,
		"updated_at":    timestamp,
	}

	return repositiory.SignUp(user)
}

func Login(phone, password string) error {

	user, err := repositiory.GetUserByPhone(phone)
	if err != nil {
		return err
	}

	passwordSalt := user.PasswordSalt
	passwordHash := hashPassword(passwordSalt, password)
	if user.PasswordHash != passwordHash {
		return errors.New("password_is_wrong")
	}

	return nil
}

func UpdatePassword(uuid, oldPassword, newPassword string) error {

	user, err := repositiory.GetUserByUuid(uuid)
	if err != nil {
		return err
	}

	passwordSalt := user.PasswordSalt
	passwordHash := hashPassword(passwordSalt, oldPassword)
	if user.PasswordHash != passwordHash {
		return errors.New("password_is_wrong")
	}

	newPasswordHash := hashPassword(user.PasswordSalt, newPassword)

	return repositiory.UpdatePassword(uuid, newPasswordHash)
}

// hashPassword 给定密码和混淆码，生成一个加密密码，这里使用hmac256。
func hashPassword(salt, password string) string {

	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	mdStr := hex.EncodeToString(h.Sum(nil))

	return strings.ToLower(mdStr)
}
