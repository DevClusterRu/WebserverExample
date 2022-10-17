package requests

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendResponse(v interface{}, w http.ResponseWriter) {
	js, _ := json.Marshal(v)
	fmt.Fprint(w, string(js))
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))

	return hex.EncodeToString(hash[:])
}

func GetToken(w http.ResponseWriter, req *http.Request) {
	SendResponse(GetTokenResponse{1234, getMD5Hash("123")}, w)
}

func GetRates(w http.ResponseWriter, req *http.Request) {
	SendResponse(GetRatesResponse{1.234}, w)
}
func GetBalance(w http.ResponseWriter, req *http.Request) {
	SendResponse(GetBalanceResponse{123}, w)
}
func SetTransaction(w http.ResponseWriter, req *http.Request) {
	SendResponse(SetTransactionResponse{0, ""}, w)
}
