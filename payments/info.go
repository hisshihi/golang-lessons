package payments

type PaymentInfo struct {
	Description string `json:"description"`
	CountUSD    int    `json:"count_usd"`
	IsCancelled bool   `json:"is_cancelled"`
}
