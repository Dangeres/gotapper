package main

type ResponseOrder struct {
	OrderId          string `json:"order_id"`
	ShopId           string `json:"shops_id"`
	SessionId        string `json:"sessions_id"`
	Id               int    `json:"id"`
	ExternalWaiterId string `json:"external_waiter_id"`
	FullPaySum       int    `json:"fullPaySum"`
	// UserId  any  `json:"users_id"`
}

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
	Amount      float64 `json:"amount"`
	Coefficient float64 `json:"coefficient"`
}

type ResponseUser struct {
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

type ResponseWaiter struct {
	User ResponseUser `json:"user"`
}

type ResponseResult struct {
	OrdersElementAll []ResponseOrderElementTapper `json:"ordersElementAll"`
	Prepayment       int                          `json:"prepayment"`
	Waiter           ResponseWaiter               `json:"waiter"`
	Orders           []ResponseOrder              `json:"orders"`
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
	Active  bool           `json:"active"`
	Result  ResponseResult `json:"result"`
	Success bool           `json:"success"`
}

type RequestBrowserId struct {
	BrowserId string `json:"browserId"`
	Plugin    string `json:"plugin"`
}

type RequestTapper struct {
	TableId   string           `json:"table_id"`
	Domen     string           `json:"domen"`
	Guest     int              `json:"guest"`
	Session   string           `json:"session"`
	BrowserId RequestBrowserId `json:"browser_id"`
}

type CustomPriceAmount struct {
	Price  int     `json:"price"`
	Amount float64 `json:"amount"`
}

type CustomPositions struct {
	PositionId string                       `json:"position_id"`
	Title      string                       `json:"title"`
	Orders     map[string]CustomPriceAmount `json:"orders"`
}

type Settings struct {
	SleepSeconds  int    `json:"sleep_seconds"`
	IsProcessData bool   `json:"is_process_data"`
	ShopToken     string `json:"shop_token"`
	CountTables   int    `json:"count_tables"`
}
