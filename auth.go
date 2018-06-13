package sip

import (
	"bytes"
	"crypto/md5"
	"fmt"

	"github.com/baubles/go-sip/header"
)

func AuthSrvChallenge(req *Request, realm string) (res *Response) {
	res = CreateResponseFromRequest(req, 401)
	wwwauth := header.NewWWWAuthenticate()
	wwwauth.Credential = "Digest"
	wwwauth.SetRealm(realm)
	wwwauth.SetNonce(randString(16))
	// wwwauth.SetAlgorithm("md5")
	// wwwauth.SetOpaque(randString(16))
	res.SetWWWAuthenticate(wwwauth)
	return res
}

func AuthSrvVerify(req *Request, cred *Cred) (res *Response, ok bool) {
	auth := req.Authorization()
	buf := new(bytes.Buffer)
	username, _ := auth.Username()
	realm, _ := auth.Realm()
	uri, _ := auth.URI()
	nonce, _ := auth.Nonce()
	response, _ := auth.Response()

	buf.WriteString(username)
	buf.WriteString(":")
	buf.WriteString(realm)
	buf.WriteString(":")
	buf.WriteString(cred.Data)
	ha1 := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))

	// fmt.Println(string(buf.Bytes()), ha1)

	buf = new(bytes.Buffer)
	buf.WriteString(req.Method)
	buf.WriteString(":")
	buf.WriteString(uri)
	ha2 := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))

	pass := fmt.Sprintf("%x", md5.Sum([]byte(ha1+":"+nonce+":"+ha2)))

	if pass != response {
		return CreateResponseFromRequest(req, 500), false
	} else {
		return CreateResponseFromRequest(req, 200), true
	}
}

func GenerateAuthorizationResponse(method, username, realm, password, nonce, uri string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(username)
	buf.WriteString(":")
	buf.WriteString(realm)
	buf.WriteString(":")
	buf.WriteString(password)
	ha1 := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))

	buf = new(bytes.Buffer)
	buf.WriteString(method)
	buf.WriteString(":")
	buf.WriteString(uri)
	ha2 := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))

	res := fmt.Sprintf("%x", md5.Sum([]byte(ha1+":"+nonce+":"+ha2)))

	return res
}
