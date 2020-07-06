package handler

import (
	"encoding/base64"
	"repEATer/model"
)

func MakeCredential(oauth model.OAuth) string {
	credential := base64.URLEncoding.EncodeToString([]byte(oauth.ConsumerKey + ":" + oauth.ConsumerSecret))
	return credential
}
