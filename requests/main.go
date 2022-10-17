package requests

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetToken(w http.ResponseWriter, req *http.Request) {
	resp := GetTokenResponse{1234, getMD5Hash("123")}
	js, _ := json.Marshal(resp)
	fmt.Fprint(w, js)
}

func GetRates(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hi!")
}
func GetBalance(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hi!")
}
func SetTransaction(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hi!")
}
