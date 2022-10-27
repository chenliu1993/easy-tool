package types

type Accounts struct {
	Data []Account `json:"Data"`
}

type Account struct {
	ConsoleLogin int    `json:"ConsoleLogin"`
	CountryCode  string `json:"CountryCode"`
	CreateTime   string `json:"CreateTime"`
	Email        string `json:"Email"`
	Name         string `json:"Name"`
	NickName     string `json:"NickName"`
	PhoneNum     string `json:"PhoneNum"`
	Remark       string `json:"Remark"`
	Uid          int64  `json:"Uid"`
	Uin          int64  `json:"Uin"`
}
