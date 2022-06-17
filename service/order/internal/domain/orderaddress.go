package domain

type OrderAddress struct {
	ID              int64
	User            int64
	OrderSn         string
	RecipientName   string
	RecipientMobile string
	Province        string
	City            string
	Districts       string
	Address         string
	PostCode        string
}
