package requests

type GetTokenResponse struct {
	Id  int64  `json:"id"`
	Key string `json:"key"`
}

type GetRatesResponse struct {
	Rate float64 `json:"rate"`
}

type GetBalanceResponse struct {
	Balance int64 `json:"balance"`
}

type SetTransactionResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
