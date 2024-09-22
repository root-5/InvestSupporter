package model

import (
	"database/sql"
)

// 17業種情報
type Sector17Info struct {
	Sector17Code int    `json:"Sector17Code"` // 17業種コード
	Sector17Name string `json:"Sector17Name"` // 17業種名
}

// 33業種情報
type Sector33Info struct {
	Sector33Code int    `json:"Sector33Code"` // 33業種コード
	Sector33Name string `json:"Sector33Name"` // 33業種名
}

// 市場区分情報
type MarketInfo struct {
	MarketCode int    `json:"MarketCode"` // 市場区分コード
	MarketName string `json:"MarketName"` // 市場区分名
}

// 上場銘柄一覧
type StocksInfo struct {
	Code               string         `json:"Code"`               // 銘柄コード
	CompanyName        sql.NullString `json:"CompanyName"`        // 会社名
	CompanyNameEnglish sql.NullString `json:"CompanyNameEnglish"` // 会社名（英語）
	Sector17Code       sql.NullInt64  `json:"Sector17Code"`       // 17業種コード
	Sector33Code       sql.NullInt64  `json:"Sector33Code"`       // 33業種コード
	ScaleCategory      sql.NullString `json:"ScaleCategory"`      // 規模コード
	MarketCode         sql.NullInt64  `json:"MarketCode"`         // 市場区分コード
}

// 財務情報
type StatementInfo struct {
	DisclosureNumber                                                             int64           `json:"開示番号"`                 // 開示番号
	Code                                                                         string          `json:"銘柄コード"`                // 銘柄コード
	DisclosedDate                                                                sql.NullTime    `json:"開示日"`                  // 開示日
	TypeOfDocument                                                               string          `json:"開示書類種別"`               // 開示書類種別
	NetSales                                                                     sql.NullInt64   `json:"売上高"`                  // 売上高
	OperatingProfit                                                              sql.NullInt64   `json:"営業利益"`                 // 営業利益
	OrdinaryProfit                                                               sql.NullInt64   `json:"経常利益"`                 // 経常利益
	Profit                                                                       sql.NullInt64   `json:"当期純利益"`                // 当期純利益
	EarningsPerShare                                                             sql.NullFloat64 `json:"一株あたり当期純利益"`           // 一株あたり当期純利益
	TotalAssets                                                                  sql.NullInt64   `json:"総資産"`                  // 総資産
	Equity                                                                       sql.NullInt64   `json:"純資産"`                  // 純資産
	EquityToAssetRatio                                                           sql.NullFloat64 `json:"自己資本比率"`               // 自己資本比率
	BookValuePerShare                                                            sql.NullFloat64 `json:"一株あたり純資産"`             // 一株あたり純資産
	CashFlowsFromOperatingActivities                                             sql.NullInt64   `json:"営業活動によるキャッシュ・フロー"`     // 営業活動によるキャッシュ・フロー
	CashFlowsFromInvestingActivities                                             sql.NullInt64   `json:"投資活動によるキャッシュ・フロー"`     // 投資活動によるキャッシュ・フロー
	CashFlowsFromFinancingActivities                                             sql.NullInt64   `json:"財務活動によるキャッシュ・フロー"`     // 財務活動によるキャッシュ・フロー
	CashAndEquivalents                                                           sql.NullInt64   `json:"現金及び現金同等物期末残高"`        // 現金及び現金同等物期末残高
	ResultDividendPerShareAnnual                                                 sql.NullFloat64 `json:"一株あたり配当実績合計"`          // 一株あたり配当実績合計
	ResultPayoutRatioAnnual                                                      sql.NullFloat64 `json:"配当性向"`                 // 配当性向
	ForecastDividendPerShareAnnual                                               sql.NullFloat64 `json:"一株あたり配当予想合計"`          // 一株あたり配当予想合計
	ForecastPayoutRatioAnnual                                                    sql.NullFloat64 `json:"予想配当性向"`               // 予想配当性向
	NextYearForecastDividendPerShareAnnual                                       sql.NullFloat64 `json:"一株あたり配当予想翌事業年度合計"`     // 一株あたり配当予想翌事業年度合計
	NextYearForecastPayoutRatioAnnual                                            sql.NullFloat64 `json:"翌事業年度予想配当性向"`          // 翌事業年度予想配当性向
	ForecastNetSales                                                             sql.NullInt64   `json:"売上高予想_期末"`             // 売上高予想_期末
	ForecastOperatingProfit                                                      sql.NullInt64   `json:"営業利益予想_期末"`            // 営業利益予想_期末
	ForecastOrdinaryProfit                                                       sql.NullInt64   `json:"経常利益予想_期末"`            // 経常利益予想_期末
	ForecastProfit                                                               sql.NullInt64   `json:"当期純利益予想_期末"`           // 当期純利益予想_期末
	ForecastEarningsPerShare                                                     sql.NullFloat64 `json:"一株あたり当期純利益予想_期末"`      // 一株あたり当期純利益予想_期末
	NextYearForecastNetSales                                                     sql.NullInt64   `json:"売上高予想_翌事業年度期末"`        // 売上高予想_翌事業年度期末
	NextYearForecastOperatingProfit                                              sql.NullInt64   `json:"営業利益予想_翌事業年度期末"`       // 営業利益予想_翌事業年度期末
	NextYearForecastOrdinaryProfit                                               sql.NullInt64   `json:"経常利益予想_翌事業年度期末"`       // 経常利益予想_翌事業年度期末
	NextYearForecastProfit                                                       sql.NullInt64   `json:"当期純利益予想_翌事業年度期末"`      // 当期純利益予想_翌事業年度期末
	NextYearForecastEarningsPerShare                                             sql.NullFloat64 `json:"一株あたり当期純利益予想_翌事業年度期末"` // 一株あたり当期純利益予想_翌事業年度期末
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock sql.NullInt64   `json:"期末発行済株式数"`             // 期末発行済株式数
}

// 上場銘柄-財務情報（API用）
type BasicInfoForApi struct {
	Code                                                                         string          `json:"銘柄コード"`                // 銘柄コード
	CompanyName                                                                  sql.NullString  `json:"会社名"`                  // 会社名
	CompanyNameEnglish                                                           sql.NullString  `json:"会社名（英語）"`              // 会社名（英語）
	Sector17Name                                                                 sql.NullString  `json:"17業種コード"`              // 17業種コード
	Sector33Name                                                                 sql.NullString  `json:"33業種コード"`              // 33業種コード
	ScaleCategory                                                                sql.NullString  `json:"規模コード"`                // 規模コード
	MarketName                                                                   sql.NullString  `json:"市場区分コード"`              // 市場区分コード
	DisclosedDate                                                                sql.NullTime    `json:"開示日"`                  // 開示日
	TypeOfDocument                                                               sql.NullString  `json:"開示書類種別"`               // 開示書類種別
	NetSales                                                                     sql.NullInt64   `json:"売上高"`                  // 売上高
	OperatingProfit                                                              sql.NullInt64   `json:"営業利益"`                 // 営業利益
	OrdinaryProfit                                                               sql.NullInt64   `json:"経常利益"`                 // 経常利益
	Profit                                                                       sql.NullInt64   `json:"当期純利益"`                // 当期純利益
	EarningsPerShare                                                             sql.NullFloat64 `json:"一株あたり当期純利益"`           // 一株あたり当期純利益
	TotalAssets                                                                  sql.NullInt64   `json:"総資産"`                  // 総資産
	Equity                                                                       sql.NullInt64   `json:"純資産"`                  // 純資産
	EquityToAssetRatio                                                           sql.NullFloat64 `json:"自己資本比率"`               // 自己資本比率
	BookValuePerShare                                                            sql.NullFloat64 `json:"一株あたり純資産"`             // 一株あたり純資産
	CashFlowsFromOperatingActivities                                             sql.NullInt64   `json:"営業活動によるキャッシュ・フロー"`     // 営業活動によるキャッシュ・フロー
	CashFlowsFromInvestingActivities                                             sql.NullInt64   `json:"投資活動によるキャッシュ・フロー"`     // 投資活動によるキャッシュ・フロー
	CashFlowsFromFinancingActivities                                             sql.NullInt64   `json:"財務活動によるキャッシュ・フロー"`     // 財務活動によるキャッシュ・フロー
	CashAndEquivalents                                                           sql.NullInt64   `json:"現金及び現金同等物期末残高"`        // 現金及び現金同等物期末残高
	ResultDividendPerShareAnnual                                                 sql.NullFloat64 `json:"一株あたり配当実績合計"`          // 一株あたり配当実績合計
	ResultPayoutRatioAnnual                                                      sql.NullFloat64 `json:"配当性向"`                 // 配当性向
	ForecastDividendPerShareAnnual                                               sql.NullFloat64 `json:"一株あたり配当予想合計"`          // 一株あたり配当予想合計
	NextYearForecastDividendPerShareAnnual                                       sql.NullFloat64 `json:"一株あたり配当予想翌事業年度合計"`     // 一株あたり配当予想翌事業年度合計
	NextYearForecastPayoutRatioAnnual                                            sql.NullFloat64 `json:"翌事業年度予想配当性向"`          // 翌事業年度予想配当性向
	ForecastNetSales                                                             sql.NullInt64   `json:"売上高予想_期末"`             // 売上高予想_期末
	ForecastOperatingProfit                                                      sql.NullInt64   `json:"営業利益予想_期末"`            // 営業利益予想_期末
	ForecastOrdinaryProfit                                                       sql.NullInt64   `json:"経常利益予想_期末"`            // 経常利益予想_期末
	ForecastProfit                                                               sql.NullInt64   `json:"当期純利益予想_期末"`           // 当期純利益予想_期末
	ForecastEarningsPerShare                                                     sql.NullFloat64 `json:"一株あたり当期純利益予想_期末"`      // 一株あたり当期純利益予想_期末
	NextYearForecastNetSales                                                     sql.NullInt64   `json:"売上高予想_翌事業年度期末"`        // 売上高予想_翌事業年度期末
	NextYearForecastOperatingProfit                                              sql.NullInt64   `json:"営業利益予想_翌事業年度期末"`       // 営業利益予想_翌事業年度期末
	NextYearForecastOrdinaryProfit                                               sql.NullInt64   `json:"経常利益予想_翌事業年度期末"`       // 経常利益予想_翌事業年度期末
	NextYearForecastProfit                                                       sql.NullInt64   `json:"当期純利益予想_翌事業年度期末"`      // 当期純利益予想_翌事業年度期末
	NextYearForecastEarningsPerShare                                             sql.NullFloat64 `json:"一株あたり当期純利益予想_翌事業年度期末"` // 一株あたり当期純利益予想_翌事業年度期末
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock sql.NullInt64   `json:"期末発行済株式数"`             // 期末発行済株式数
}

// 株価四本値情報
type PriceInfo struct {
	Date             string          `json:"日付"`     // 日付
	Code             string          `json:"銘柄コード"`  // 銘柄コード
	AdjustmentOpen   sql.NullFloat64 `json:"調整後始値"`  // 調整後始値
	AdjustmentHigh   sql.NullFloat64 `json:"調整後高値"`  // 調整後高値
	AdjustmentLow    sql.NullFloat64 `json:"調整後安値"`  // 調整後安値
	AdjustmentClose  sql.NullFloat64 `json:"調整後終値"`  // 調整後終値
	AdjustmentVolume sql.NullFloat64 `json:"調整後出来高"` // 調整後出来高
}
