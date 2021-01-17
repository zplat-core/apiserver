package service

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestActivateLink(t *testing.T) {
	email := "saltbo@foxmail.com"
	token := "aasdfdssssdffqwwessdfzxvd"
	key := encodeToKey(email, token)
	fmt.Println(key)
	email1, token2 := decodeFromKey(key)
	assert.Equal(t, email, email1)
	assert.Equal(t, token, token2)
}
