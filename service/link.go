package service

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func Link2SignIn(redirect string) string {
	return fmt.Sprintf("/moreu/signin?redirect=%s", redirect)
}

func Link2ServerError(err error) string {
	return fmt.Sprintf("/moreu/error?msg=%s", err.Error())
}

func Link2Forbidden() string {
	return fmt.Sprintf("/moreu/forbidden")
}

func ActivateLink(origin, email, token string) string {
	return fmt.Sprintf("%s/moreu/signin/%s", origin, encodeToKey(email, token))
}

func PasswordRestLink(origin, email, token string) string {
	return fmt.Sprintf("%s/moreu/password-reset/%s", origin, encodeToKey(email, token))
}

var base64Encode = base64.URLEncoding.EncodeToString
var base64Decode = base64.URLEncoding.DecodeString

const moreuSplitKey = "|moreu|"

func encodeToKey(email, token string) string {
	return base64Encode([]byte(email + moreuSplitKey + token))
}

func decodeFromKey(key string) (email, token string) {
	bb, _ := base64Decode(key)
	sss := strings.Split(string(bb), moreuSplitKey)
	email = sss[0]
	token = sss[1]
	return
}
