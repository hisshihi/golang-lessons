package payments

type PaymentInfo struct {
	ID          int
	Description string
	CountUSD    int
	IsCancelled bool
}
