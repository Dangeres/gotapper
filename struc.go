package main

type ResponseOrderElementTapper struct {
	PositionId  string  `json:"position_id"`
	Sort        int     `json:"sort"`
	FinalPrice  int     `json:"finalPrice"`
	Price       int     `json:"price"`
	Name        string  `json:"name"`
	OrderId     string  `json:"order_id"`
	OrdersId    int     `json:"orders_id"`
	Id          int     `json:"id"`
	Status      string  `json:"status"`
	StatusPay   bool    `json:"status_pay"`
	Amount      int     `json:"amount"`
	Coefficient float64 `json:"coefficient"`
}

type UserTapper struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	LastName            string `json:"last_name"`
	SecondName          string `json:"second_name"`
	Phone               string `json:"phone"`
	Password            string `json:"password"`
	RolesId             int    `json:"roles_id"`
	Email               string `json:"email"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	LastShopId          int    `json:"last_shop_id"`
	UrlPasswordRecovery string `json:"url_password_recovery"`
	Messenger           string `json:"messenger"`
	DeletedAt           string `json:"deleted_at"`
}

type WaiterTapper struct {
	User UserTapper `json:"user"`
}

type ResponseResultTapper struct {
	OrdersElementAll []ResponseOrderElementTapper `json:"ordersElementAll"`
	Prepayment       int                          `json:"prepayment"`
	Waiter           WaiterTapper                 `json:"waiter"`
	// TableCode        string                       `json:"tableCode"`
	Guest            int    `json:"guest"`
	Discount         int    `json:"discount"`
	ResponseTemplate string `json:"responseTemplate"`
	Recipient        string `json:"recipient"`
	Name             string `json:"name"`
	TipsType         string `json:"tips_type"`
	LoyaltySystem    string `json:"loyalty_system"`
	FeeType          string `json:"fee_type"`
}

type ResponseTapper struct {
	Active  bool                 `json:"active"`
	Result  ResponseResultTapper `json:"result"`
	Success bool                 `json:"success"`
}

type BrowserId struct {
	BrowserId string `json:"browserId"`
	Plugin    string `json:"plugin"`
}

type RequestTapper struct {
	TableId   string    `json:"table_id"`
	Domen     string    `json:"domen"`
	Guest     int       `json:"guest"`
	Session   string    `json:"session"`
	BrowserId BrowserId `json:"browser_id"`
}
