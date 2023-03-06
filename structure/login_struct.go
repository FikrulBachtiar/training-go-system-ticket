package structure

type RequestLogin struct {
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	Terminal_code string `json:"terminal_code" validate:"required"`
	Imei          string `json:"imei" validate:"required"`
	Ip            string `json:"ip" validate:"required"`
}

type DataLogin struct {
	Token        string        `json:"token"`
	Ob_data      OBData        `json:"ob_data"`
	Config_param ConfigParam   `json:"config_param"`
	List_station []ListStation `json:"list_station"`
	List_fare    []ListFare    `json:"list_fare"`
	User_info    UserInfo      `json:"user_info"`
}

type OBData struct {
	Code     int     `json:"code"`
	Message  string  `json:"message"`
	Ob_code  string  `json:"ob_code"`
	Ob_cash  float64 `json:"ob_cash"`
	Cash_in  float64 `json:"cash_in"`
	Cash_out float64 `json:"cash_out"`
	Cash_now float64 `json:"cash_now"`
}

type ConfigParam struct {
	Max_sales      string  `json:"max_sales"`
	Pinalty_amount float64 `json:"pinalty_amount"`
}

type ListStation struct {
	Station_code string `json:"station_code"`
	Station_name string `json:"station_name"`
}

type ListFare struct {
	Origin_code      string  `json:"origin_code"`
	Destination_code string  `json:"destination_code"`
	Fare             float64 `json:"fare"`
}

type UserInfo struct {
	Fullname        string `json:"fullname"`
	Username        string `json:"username"`
	User_group_code string `json:"user_group_code"`
	User_group_name string `json:"user_group_name"`
	Station_code    string `json:"station_code"`
}