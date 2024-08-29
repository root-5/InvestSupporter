package model

import (
	"time"
)

// 上場銘柄一覧
type StocksInfo struct {
	Code               string `json:"Code"`               // 銘柄コード
	CompanyName        string `json:"CompanyName"`        // 会社名
	CompanyNameEnglish string `json:"CompanyNameEnglish"` // 会社名（英語）
	Sector17Code       int    `json:"Sector17Code"`       // 17業種コード
	Sector33Code       int    `json:"Sector33Code"`       // 33業種コード
	ScaleCategory      string `json:"ScaleCategory"`      // 規模コード
	MarketCode         int    `json:"MarketCode"`         // 市場区分コード
}

// 財務情報
type FinancialInfo struct {
	Code                                                                         string    `json:"Code"`                                                                         // 銘柄コード
	DisclosedDate                                                                time.Time `json:"DisclosedDate"`                                                                // 開示日
	DisclosedTime                                                                time.Time `json:"DisclosedTime"`                                                                // 開示時刻
	NetSales                                                                     int       `json:"NetSales"`                                                                     // 売上高
	OperatingProfit                                                              int       `json:"OperatingProfit"`                                                              // 営業利益
	OrdinaryProfit                                                               int       `json:"OrdinaryProfit"`                                                               // 経常利益
	Profit                                                                       int       `json:"Profit"`                                                                       // 当期純利益
	EarningsPerShare                                                             float64   `json:"EarningsPerShare"`                                                             // 一株あたり当期純利益
	TotalAssets                                                                  int       `json:"TotalAssets"`                                                                  // 総資産
	Equity                                                                       int       `json:"Equity"`                                                                       // 純資産
	EquityToAssetRatio                                                           float64   `json:"EquityToAssetRatio"`                                                           // 自己資本比率
	BookValuePerShare                                                            float64   `json:"BookValuePerShare"`                                                            // 一株あたり純資産
	CashFlowsFromOperatingActivities                                             int       `json:"CashFlowsFromOperatingActivities"`                                             // 営業活動によるキャッシュ・フロー
	CashFlowsFromInvestingActivities                                             int       `json:"CashFlowsFromInvestingActivities"`                                             // 投資活動によるキャッシュ・フロー
	CashFlowsFromFinancingActivities                                             int       `json:"CashFlowsFromFinancingActivities"`                                             // 財務活動によるキャッシュ・フロー
	CashAndEquivalents                                                           int       `json:"CashAndEquivalents"`                                                           // 現金及び現金同等物期末残高
	ResultDividendPerShareAnnual                                                 float64   `json:"ResultDividendPerShareAnnual"`                                                 // 一株あたり配当実績合計
	ResultPayoutRatioAnnual                                                      float64   `json:"ResultPayoutRatioAnnual"`                                                      // 配当性向
	ForecastDividendPerShareAnnual                                               float64   `json:"ForecastDividendPerShareAnnual"`                                               // 一株あたり配当予想合計
	ForecastPayoutRatioAnnual                                                    float64   `json:"ForecastPayoutRatioAnnual"`                                                    // 予想配当性向
	NextYearForecastDividendPerShareAnnual                                       float64   `json:"NextYearForecastDividendPerShareAnnual"`                                       // 一株あたり配当予想翌事業年度合計
	NextYearForecastPayoutRatioAnnual                                            float64   `json:"NextYearForecastPayoutRatioAnnual"`                                            // 翌事業年度予想配当性向
	ForecastNetSales                                                             int       `json:"ForecastNetSales"`                                                             // 売上高予想_期末
	ForecastOperatingProfit                                                      int       `json:"ForecastOperatingProfit"`                                                      // 営業利益予想_期末
	ForecastOrdinaryProfit                                                       int       `json:"ForecastOrdinaryProfit"`                                                       // 経常利益予想_期末
	ForecastProfit                                                               int       `json:"ForecastProfit"`                                                               // 当期純利益予想_期末
	ForecastEarningsPerShare                                                     float64   `json:"ForecastEarningsPerShare"`                                                     // 一株あたり当期純利益予想_期末
	NextYearForecastNetSales                                                     int       `json:"NextYearForecastNetSales"`                                                     // 売上高予想_翌事業年度期末
	NextYearForecastOperatingProfit                                              int       `json:"NextYearForecastOperatingProfit"`                                              // 営業利益予想_翌事業年度期末
	NextYearForecastOrdinaryProfit                                               int       `json:"NextYearForecastOrdinaryProfit"`                                               // 経常利益予想_翌事業年度期末
	NextYearForecastProfit                                                       int       `json:"NextYearForecastProfit"`                                                       // 当期純利益予想_翌事業年度期末
	NextYearForecastEarningsPerShare                                             float64   `json:"NextYearForecastEarningsPerShare"`                                             // 一株あたり当期純利益予想_翌事業年度期末
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock int       `json:"NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock"` // 期末発行済株式数
}

// 上場銘柄テーブルと財務情報テーブルを結合したデータ
type FinancialInfoForApi struct {
	StocksInfo    StocksInfo    `json:"StocksInfo"`    // 上場銘柄一覧
	FinancialInfo FinancialInfo `json:"FinancialInfo"` // 財務情報
}
