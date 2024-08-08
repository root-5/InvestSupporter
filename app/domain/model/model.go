package model

// 上場銘柄一覧
type StockInfo struct {
	// Date              string `json:"Date"`
	Code              string `json:"Code"`
	CompanyName       string `json:"CompanyName"`
	CompanyNameEnglish string `json:"CompanyNameEnglish"`
	Sector17Code      string `json:"Sector17Code"`
	// Sector17CodeName  string `json:"Sector17CodeName"`
	Sector33Code      string `json:"Sector33Code"`
	// Sector33CodeName  string `json:"Sector33CodeName"`
	ScaleCategory     string `json:"ScaleCategory"`
	MarketCode        string `json:"MarketCode"`
	// MarketCodeName    string `json:"MarketCodeName"`
	// MarginCode        string `json:"MarginCode"`  // Light プランでは取得できない
	// MarginCodeName    string `json:"MarginCodeName"`
}