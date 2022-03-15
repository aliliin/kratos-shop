package domain

type Address struct {
	ID        int64
	UserID    int64
	IsDefault int
	Mobile    string
	Name      string
	Province  string
	City      string
	Districts string
	Address   string
	PostCode  string
}
