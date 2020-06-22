package handler

import (
	"../model"
	"encoding/base64"
)

func MakeCredential(oauth model.OAuth) string {
	credential := base64.URLEncoding.EncodeToString([]byte(oauth.ConsumerKey + ":" + oauth.ConsumerSecret))
	return credential
}
