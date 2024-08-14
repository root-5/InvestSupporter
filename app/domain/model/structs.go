package model

// 上場銘柄一覧
type StockInfo struct {
	Code              string `json:"Code"`
	CompanyName       string `json:"CompanyName"`
	CompanyNameEnglish string `json:"CompanyNameEnglish"`
	Sector17Code      int `json:"Sector17Code"`
	Sector33Code      int `json:"Sector33Code"`
	ScaleCategory     string `json:"ScaleCategory"`
	MarketCode        int `json:"MarketCode"`
}
