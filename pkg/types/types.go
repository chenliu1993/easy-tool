package types

type Accounts struct {
	Data []Account `json:"Data"`
}

type Account struct {
	ConsoleLogin int    `json:"ConsoleLogin"`
	CountryCode  int    `json:"CountryCode"`
	CreateTime   string `json:"CreateTime"`
	Email        string `json:"Email"`
	Name         string `json:"Name"`
	NickName     string `json:"NickName"`
	PhoneNum     string `json:"PhoneNum"`
	Remark       string `json:"Remark"`
	Uid          string `json:"Uid"`
	Uin          string `json:"Uin"`
}
