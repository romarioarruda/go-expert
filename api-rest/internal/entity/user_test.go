package entity

import (
	"go-expert/api-rest/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(test *testing.T) {
	user, err := entity.NewUser("Jhon Dow", "j@j.com", "123")

	assert.Nil(test, err)
}
