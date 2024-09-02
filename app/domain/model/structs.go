package model

import (
	"database/sql"
)

// 上場銘柄一覧
type StocksInfo struct {
	Code               string        `json:"Code"`               // 銘柄コード
	CompanyName        string        `json:"CompanyName"`        // 会社名
	CompanyNameEnglish string        `json:"CompanyNameEnglish"` // 会社名（英語）
	Sector17Code       sql.NullInt64 `json:"Sector17Code"`       // 17業種コード
	Sector33Code       sql.NullInt64 `json:"Sector33Code"`       // 33業種コード
	ScaleCategory      string        `json:"ScaleCategory"`      // 規模コード
	MarketCode         sql.NullInt64 `json:"MarketCode"`         // 市場区分コード
}

// 財務情報
type FinancialInfo struct {
	Code                                                                         string          `json:"Code"`                                                                         // 銘柄コード
	DisclosedDate                                                                sql.NullTime    `json:"DisclosedDate"`                                                                // 開示日
	DisclosedTime                                                                sql.NullTime    `json:"DisclosedTime"`                                                                // 開示時刻
	NetSales                                                                     sql.NullInt64   `json:"NetSales"`                                                                     // 売上高
	OperatingProfit                                                              sql.NullInt64   `json:"OperatingProfit"`                                                              // 営業利益
	OrdinaryProfit                                                               sql.NullInt64   `json:"OrdinaryProfit"`                                                               // 経常利益
	Profit                                                                       sql.NullInt64   `json:"Profit"`                                                                       // 当期純利益
	EarningsPerShare                                                             sql.NullFloat64 `json:"EarningsPerShare"`                                                             // 一株あたり当期純利益
	TotalAssets                                                                  sql.NullInt64   `json:"TotalAssets"`                                                                  // 総資産
	Equity                                                                       sql.NullInt64   `json:"Equity"`                                                                       // 純資産
	EquityToAssetRatio                                                           sql.NullFloat64 `json:"EquityToAssetRatio"`                                                           // 自己資本比率
	BookValuePerShare                                                            sql.NullFloat64 `json:"BookValuePerShare"`                                                            // 一株あたり純資産
	CashFlowsFromOperatingActivities                                             sql.NullInt64   `json:"CashFlowsFromOperatingActivities"`                                             // 営業活動によるキャッシュ・フロー
	CashFlowsFromInvestingActivities                                             sql.NullInt64   `json:"CashFlowsFromInvestingActivities"`                                             // 投資活動によるキャッシュ・フロー
	CashFlowsFromFinancingActivities                                             sql.NullInt64   `json:"CashFlowsFromFinancingActivities"`                                             // 財務活動によるキャッシュ・フロー
	CashAndEquivalents                                                           sql.NullInt64   `json:"CashAndEquivalents"`                                                           // 現金及び現金同等物期末残高
	ResultDividendPerShareAnnual                                                 sql.NullFloat64 `json:"ResultDividendPerShareAnnual"`                                                 // 一株あたり配当実績合計
	ResultPayoutRatioAnnual                                                      sql.NullFloat64 `json:"ResultPayoutRatioAnnual"`                                                      // 配当性向
	ForecastDividendPerShareAnnual                                               sql.NullFloat64 `json:"ForecastDividendPerShareAnnual"`                                               // 一株あたり配当予想合計
	ForecastPayoutRatioAnnual                                                    sql.NullFloat64 `json:"ForecastPayoutRatioAnnual"`                                                    // 予想配当性向
	NextYearForecastDividendPerShareAnnual                                       sql.NullFloat64 `json:"NextYearForecastDividendPerShareAnnual"`                                       // 一株あたり配当予想翌事業年度合計
	NextYearForecastPayoutRatioAnnual                                            sql.NullFloat64 `json:"NextYearForecastPayoutRatioAnnual"`                                            // 翌事業年度予想配当性向
	ForecastNetSales                                                             sql.NullInt64   `json:"ForecastNetSales"`                                                             // 売上高予想_期末
	ForecastOperatingProfit                                                      sql.NullInt64   `json:"ForecastOperatingProfit"`                                                      // 営業利益予想_期末
	ForecastOrdinaryProfit                                                       sql.NullInt64   `json:"ForecastOrdinaryProfit"`                                                       // 経常利益予想_期末
	ForecastProfit                                                               sql.NullInt64   `json:"ForecastProfit"`                                                               // 当期純利益予想_期末
	ForecastEarningsPerShare                                                     sql.NullFloat64 `json:"ForecastEarningsPerShare"`                                                     // 一株あたり当期純利益予想_期末
	NextYearForecastNetSales                                                     sql.NullInt64   `json:"NextYearForecastNetSales"`                                                     // 売上高予想_翌事業年度期末
	NextYearForecastOperatingProfit                                              sql.NullInt64   `json:"NextYearForecastOperatingProfit"`                                              // 営業利益予想_翌事業年度期末
	NextYearForecastOrdinaryProfit                                               sql.NullInt64   `json:"NextYearForecastOrdinaryProfit"`                                               // 経常利益予想_翌事業年度期末
	NextYearForecastProfit                                                       sql.NullInt64   `json:"NextYearForecastProfit"`                                                       // 当期純利益予想_翌事業年度期末
	NextYearForecastEarningsPerShare                                             sql.NullFloat64 `json:"NextYearForecastEarningsPerShare"`                                             // 一株あたり当期純利益予想_翌事業年度期末
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock sql.NullInt64   `json:"NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock"` // 期末発行済株式数
}

// 上場銘柄-財務情報（API用）
type FinancialInfoForApi struct {
	Code                                                                         string          `json:"Code"`                                                                         // 銘柄コード
	CompanyName                                                                  string          `json:"CompanyName"`                                                                  // 会社名
	CompanyNameEnglish                                                           string          `json:"CompanyNameEnglish"`                                                           // 会社名（英語）
	Sector17Code                                                                 sql.NullInt64   `json:"Sector17Code"`                                                                 // 17業種コード
	Sector33Code                                                                 sql.NullInt64   `json:"Sector33Code"`                                                                 // 33業種コード
	ScaleCategory                                                                string          `json:"ScaleCategory"`                                                                // 規模コード
	MarketCode                                                                   sql.NullInt64   `json:"MarketCode"`                                                                   // 市場区分コード
	DisclosedDate                                                                sql.NullTime    `json:"DisclosedDate"`                                                                // 開示日
	NetSales                                                                     sql.NullInt64   `json:"NetSales"`                                                                     // 売上高
	OperatingProfit                                                              sql.NullInt64   `json:"OperatingProfit"`                                                              // 営業利益
	OrdinaryProfit                                                               sql.NullInt64   `json:"OrdinaryProfit"`                                                               // 経常利益
	Profit                                                                       sql.NullInt64   `json:"Profit"`                                                                       // 当期純利益
	EarningsPerShare                                                             sql.NullFloat64 `json:"EarningsPerShare"`                                                             // 一株あたり当期純利益
	TotalAssets                                                                  sql.NullInt64   `json:"TotalAssets"`                                                                  // 総資産
	Equity                                                                       sql.NullInt64   `json:"Equity"`                                                                       // 純資産
	EquityToAssetRatio                                                           sql.NullFloat64 `json:"EquityToAssetRatio"`                                                           // 自己資本比率
	BookValuePerShare                                                            sql.NullFloat64 `json:"BookValuePerShare"`                                                            // 一株あたり純資産
	CashFlowsFromOperatingActivities                                             sql.NullInt64   `json:"CashFlowsFromOperatingActivities"`                                             // 営業活動によるキャッシュ・フロー
	CashFlowsFromInvestingActivities                                             sql.NullInt64   `json:"CashFlowsFromInvestingActivities"`                                             // 投資活動によるキャッシュ・フロー
	CashFlowsFromFinancingActivities                                             sql.NullInt64   `json:"CashFlowsFromFinancingActivities"`                                             // 財務活動によるキャッシュ・フロー
	CashAndEquivalents                                                           sql.NullInt64   `json:"CashAndEquivalents"`                                                           // 現金及び現金同等物期末残高
	ResultDividendPerShareAnnual                                                 sql.NullFloat64 `json:"ResultDividendPerShareAnnual"`                                                 // 一株あたり配当実績合計
	ResultPayoutRatioAnnual                                                      sql.NullFloat64 `json:"ResultPayoutRatioAnnual"`                                                      // 配当性向
	ForecastDividendPerShareAnnual                                               sql.NullFloat64 `json:"ForecastDividendPerShareAnnual"`                                               // 一株あたり配当予想合計
	ForecastPayoutRatioAnnual                                                    sql.NullFloat64 `json:"ForecastPayoutRatioAnnual"`                                                    // 予想配当性向
	NextYearForecastDividendPerShareAnnual                                       sql.NullFloat64 `json:"NextYearForecastDividendPerShareAnnual"`                                       // 一株あたり配当予想翌事業年度合計
	NextYearForecastPayoutRatioAnnual                                            sql.NullFloat64 `json:"NextYearForecastPayoutRatioAnnual"`                                            // 翌事業年度予想配当性向
	ForecastNetSales                                                             sql.NullInt64   `json:"ForecastNetSales"`                                                             // 売上高予想_期末
	ForecastOperatingProfit                                                      sql.NullInt64   `json:"ForecastOperatingProfit"`                                                      // 営業利益予想_期末
	ForecastOrdinaryProfit                                                       sql.NullInt64   `json:"ForecastOrdinaryProfit"`                                                       // 経常利益予想_期末
	ForecastProfit                                                               sql.NullInt64   `json:"ForecastProfit"`                                                               // 当期純利益予想_期末
	ForecastEarningsPerShare                                                     sql.NullFloat64 `json:"ForecastEarningsPerShare"`                                                     // 一株あたり当期純利益予想_期末
	NextYearForecastNetSales                                                     sql.NullInt64   `json:"NextYearForecastNetSales"`                                                     // 売上高予想_翌事業年度期末
	NextYearForecastOperatingProfit                                              sql.NullInt64   `json:"NextYearForecastOperatingProfit"`                                              // 営業利益予想_翌事業年度期末
	NextYearForecastOrdinaryProfit                                               sql.NullInt64   `json:"NextYearForecastOrdinaryProfit"`                                               // 経常利益予想_翌事業年度期末
	NextYearForecastProfit                                                       sql.NullInt64   `json:"NextYearForecastProfit"`                                                       // 当期純利益予想_翌事業年度期末
	NextYearForecastEarningsPerShare                                             sql.NullFloat64 `json:"NextYearForecastEarningsPerShare"`                                             // 一株あたり当期純利益予想_翌事業年度期末
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock sql.NullInt64   `json:"NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock"` // 期末発行済株式数
}
