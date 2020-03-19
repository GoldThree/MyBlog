package helper

import "github.com/satori/go.uuid"

// NewV4 returns a random uuid4 string.
func NewV4() string {

	return uuid.Must(uuid.NewV4(), nil).String()

}
