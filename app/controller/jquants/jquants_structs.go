package jquants

// ====================================================================================
// 各APIのレスポンスボディの構造体を定義
// ====================================================================================

// 上場銘柄一覧
type jquantsStockInfo struct {
	// Date               string `json:"Date"`               // 日付
	Code               string `json:"Code"`               // 銘柄コード
	CompanyName        string `json:"CompanyName"`        // 会社名
	CompanyNameEnglish string `json:"CompanyNameEnglish"` // 会社名（英語）
	Sector17Code       string `json:"Sector17Code"`       // 17業種コード
	// Sector17CodeName   string `json:"Sector17CodeName"`   // 17業種コード名
	Sector33Code       string `json:"Sector33Code"`       // 33業種コード
	// Sector33CodeName   string `json:"Sector33CodeName"`   // 33業種コード名
	ScaleCategory      string `json:"ScaleCategory"`      // 規模コード
	MarketCode         string `json:"MarketCode"`         // 市場区分コード
	// MarketCodeName     string `json:"MarketCodeName"`     // 市場区分名
	// MarginCode         string `json:"MarginCode"`         // 貸借信用区分
	// MarginCodeName     string `json:"MarginCodeName"`     // 貸借信用区分名
}

// 財務情報
type jquantsStatementInfo struct {
	DisclosedDate                                                                string `json:"DisclosedDate"`                                                                // 開示日
	// DisclosedTime                                                                string `json:"DisclosedTime"`                                                                // 開示時刻
	Code                                                                         string `json:"LocalCode"`                                                                    // 銘柄コード（5桁）
	DisclosureNumber                                                             string `json:"DisclosureNumber"`                                                             // 開示番号
	TypeOfDocument                                                               string `json:"TypeOfDocument"`                                                               // 開示書類種別
	// TypeOfCurrentPeriod                                                          string `json:"TypeOfCurrentPeriod"`                                                          // 当会計期間の種類
	// CurrentPeriodStartDate                                                       string `json:"CurrentPeriodStartDate"`                                                       // 当会計期間開始日
	// CurrentPeriodEndDate                                                         string `json:"CurrentPeriodEndDate"`                                                         // 当会計期間終了日
	// CurrentFiscalYearStartDate                                                   string `json:"CurrentFiscalYearStartDate"`                                                   // 当事業年度開始日
	// CurrentFiscalYearEndDate                                                     string `json:"CurrentFiscalYearEndDate"`                                                     // 当事業年度終了日
	// NextFiscalYearStartDate                                                      string `json:"NextFiscalYearStartDate"`                                                      // 翌事業年度開始日
	// NextFiscalYearEndDate                                                        string `json:"NextFiscalYearEndDate"`                                                        // 翌事業年度終了日
	NetSales                                                                     string `json:"NetSales"`                                                                     // 売上高
	OperatingProfit                                                              string `json:"OperatingProfit"`                                                              // 営業利益
	OrdinaryProfit                                                               string `json:"OrdinaryProfit"`                                                               // 経常利益
	Profit                                                                       string `json:"Profit"`                                                                       // 当期純利益
	EarningsPerShare                                                             string `json:"EarningsPerShare"`                                                             // 一株あたり当期純利益
	// DilutedEarningsPerShare                                                      string `json:"DilutedEarningsPerShare"`                                                      // 潜在株式調整後一株あたり当期純利益
	TotalAssets                                                                  string `json:"TotalAssets"`                                                                  // 総資産
	Equity                                                                       string `json:"Equity"`                                                                       // 純資産
	EquityToAssetRatio                                                           string `json:"EquityToAssetRatio"`                                                           // 自己資本比率
	BookValuePerShare                                                            string `json:"BookValuePerShare"`                                                            // 一株あたり純資産
	CashFlowsFromOperatingActivities                                             string `json:"CashFlowsFromOperatingActivities"`                                             // 営業活動によるキャッシュ・フロー
	CashFlowsFromInvestingActivities                                             string `json:"CashFlowsFromInvestingActivities"`                                             // 投資活動によるキャッシュ・フロー
	CashFlowsFromFinancingActivities                                             string `json:"CashFlowsFromFinancingActivities"`                                             // 財務活動によるキャッシュ・フロー
	CashAndEquivalents                                                           string `json:"CashAndEquivalents"`                                                           // 現金及び現金同等物期末残高
	// ResultDividendPerShare1stQuarter                                             string `json:"ResultDividendPerShare1stQuarter"`                                             // 一株あたり配当実績_第1四半期末
	// ResultDividendPerShare2ndQuarter                                             string `json:"ResultDividendPerShare2ndQuarter"`                                             // 一株あたり配当実績_第2四半期末
	// ResultDividendPerShare3rdQuarter                                             string `json:"ResultDividendPerShare3rdQuarter"`                                             // 一株あたり配当実績_第3四半期末
	// ResultDividendPerShareFiscalYearEnd                                          string `json:"ResultDividendPerShareFiscalYearEnd"`                                          // 一株あたり配当実績_期末
	ResultDividendPerShareAnnual                                                 string `json:"ResultDividendPerShareAnnual"`                                                 // 一株あたり配当実績_合計
	// DistributionsPerUnitREIT                                                     string `json:"DistributionsPerUnitREIT"`                                                     // 1口当たり分配金
	ResultTotalDividendPaidAnnual                                                string `json:"ResultTotalDividendPaidAnnual"`                                                // 配当金総額
	ResultPayoutRatioAnnual                                                      string `json:"ResultPayoutRatioAnnual"`                                                      // 配当性向
	// ForecastDividendPerShare1stQuarter                                           string `json:"ForecastDividendPerShare1stQuarter"`                                           // 一株あたり配当予想_第1四半期末
	// ForecastDividendPerShare2ndQuarter                                           string `json:"ForecastDividendPerShare2ndQuarter"`                                           // 一株あたり配当予想_第2四半期末
	// ForecastDividendPerShare3rdQuarter                                           string `json:"ForecastDividendPerShare3rdQuarter"`                                           // 一株あたり配当予想_第3四半期末
	// ForecastDividendPerShareFiscalYearEnd                                        string `json:"ForecastDividendPerShareFiscalYearEnd"`                                        // 一株あたり配当予想_期末
	ForecastDividendPerShareAnnual                                               string `json:"ForecastDividendPerShareAnnual"`                                               // 一株あたり配当予想_合計
	// ForecastDistributionsPerUnitREIT                                             string `json:"ForecastDistributionsPerUnitREIT"`                                             // 1口当たり予想分配金
	ForecastTotalDividendPaidAnnual                                              string `json:"ForecastTotalDividendPaidAnnual"`                                              // 予想配当金総額
	ForecastPayoutRatioAnnual                                                    string `json:"ForecastPayoutRatioAnnual"`                                                    // 予想配当性向
	// NextYearForecastDividendPerShare1stQuarter                                   string `json:"NextYearForecastDividendPerShare1stQuarter"`                                   // 一株あたり配当予想_翌事業年度第1四半期末
	// NextYearForecastDividendPerShare2ndQuarter                                   string `json:"NextYearForecastDividendPerShare2ndQuarter"`                                   // 一株あたり配当予想_翌事業年度第2四半期末
	// NextYearForecastDividendPerShare3rdQuarter                                   string `json:"NextYearForecastDividendPerShare3rdQuarter"`                                   // 一株あたり配当予想_翌事業年度第3四半期末
	// NextYearForecastDividendPerShareFiscalYearEnd                                string `json:"NextYearForecastDividendPerShareFiscalYearEnd"`                                // 一株あたり配当予想_翌事業年度期末
	NextYearForecastDividendPerShareAnnual                                       string `json:"NextYearForecastDividendPerShareAnnual"`                                       // 一株あたり配当予想_翌事業年度合計
	// NextYearForecastDistributionsPerUnitREIT                                     string `json:"NextYearForecastDistributionsPerUnitREIT"`                                     // 1口当たり翌事業年度予想分配金
	NextYearForecastPayoutRatioAnnual                                            string `json:"NextYearForecastPayoutRatioAnnual"`                                            // 翌事業年度予想配当性向
	// ForecastNetSales2ndQuarter                                                   string `json:"ForecastNetSales2ndQuarter"`                                                   // 売上高_予想_第2四半期末
	// ForecastOperatingProfit2ndQuarter                                            string `json:"ForecastOperatingProfit2ndQuarter"`                                            // 営業利益_予想_第2四半期末
	// ForecastOrdinaryProfit2ndQuarter                                             string `json:"ForecastOrdinaryProfit2ndQuarter"`                                             // 経常利益_予想_第2四半期末
	// ForecastProfit2ndQuarter                                                     string `json:"ForecastProfit2ndQuarter"`                                                     // 当期純利益_予想_第2四半期末
	// ForecastEarningsPerShare2ndQuarter                                           string `json:"ForecastEarningsPerShare2ndQuarter"`                                           // 一株あたり当期純利益_予想_第2四半期末
	// NextYearForecastNetSales2ndQuarter                                           string `json:"NextYearForecastNetSales2ndQuarter"`                                           // 売上高_予想_翌事業年度第2四半期末
	// NextYearForecastOperatingProfit2ndQuarter                                    string `json:"NextYearForecastOperatingProfit2ndQuarter"`                                    // 営業利益_予想_翌事業年度第2四半期末
	// NextYearForecastOrdinaryProfit2ndQuarter                                     string `json:"NextYearForecastOrdinaryProfit2ndQuarter"`                                     // 経常利益_予想_翌事業年度第2四半期末
	// NextYearForecastProfit2ndQuarter                                             string `json:"NextYearForecastProfit2ndQuarter"`                                             // 当期純利益_予想_翌事業年度第2四半期末
	// NextYearForecastEarningsPerShare2ndQuarter                                   string `json:"NextYearForecastEarningsPerShare2ndQuarter"`                                   // 一株あたり当期純利益_予想_翌事業年度第2四半期末
	ForecastNetSales                                                             string `json:"ForecastNetSales"`                                                             // 売上高_予想_期末
	ForecastOperatingProfit                                                      string `json:"ForecastOperatingProfit"`                                                      // 営業利益_予想_期末
	ForecastOrdinaryProfit                                                       string `json:"ForecastOrdinaryProfit"`                                                       // 経常利益_予想_期末
	ForecastProfit                                                               string `json:"ForecastProfit"`                                                               // 当期純利益_予想_期末
	ForecastEarningsPerShare                                                     string `json:"ForecastEarningsPerShare"`                                                     // 一株あたり当期純利益_予想_期末
	NextYearForecastNetSales                                                     string `json:"NextYearForecastNetSales"`                                                     // 売上高_予想_翌事業年度期末
	NextYearForecastOperatingProfit                                              string `json:"NextYearForecastOperatingProfit"`                                              // 営業利益_予想_翌事業年度期末
	NextYearForecastOrdinaryProfit                                               string `json:"NextYearForecastOrdinaryProfit"`                                               // 経常利益_予想_翌事業年度期末
	NextYearForecastProfit                                                       string `json:"NextYearForecastProfit"`                                                       // 当期純利益_予想_翌事業年度期末
	NextYearForecastEarningsPerShare                                             string `json:"NextYearForecastEarningsPerShare"`                                             // 一株あたり当期純利益_予想_翌事業年度期末
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock string `json:"NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock"` // 期末発行済株式数
}

// 株価四本値
type jquantsPriceInfo struct {
	Date                      string `json: "Date"`                      // 日付
	Code                      string `json: "Code"`                      // 銘柄コード
	// Open                      string `json: "Open"`                      // 始値（調整前）
	// High                      string `json: "High"`                      // 高値（調整前）
	// Low                       string `json: "Low"`                       // 安値（調整前）
	// Close                     string `json: "Close"`                     // 終値（調整前）
	// UpperLimit                string `json: "UpperLimit"`                // 日通ストップ高を記録したか、否かを表すフラグ
	// LowerLimit                string `json: "LowerLimit"`                // 日通ストップ安を記録したか、否かを表すフラグ
	// Volume                    string `json: "Volume"`                    // 取引高（調整前）
	// TurnoverValue             string `json: "TurnoverValue"`             // 取引代金
	AdjustmentFactor          string `json: "AdjustmentFactor"`          // 調整係数
	AdjustmentOpen            any    `json: "AdjustmentOpen"`            // 調整済み始値（売買がない場合に Null になるため any）
	AdjustmentHigh            any    `json: "AdjustmentHigh"`            // 調整済み高値（売買がない場合に Null になるため any）
	AdjustmentLow             any    `json: "AdjustmentLow"`             // 調整済み安値（売買がない場合に Null になるため any）
	AdjustmentClose           any    `json: "AdjustmentClose"`           // 調整済み終値（売買がない場合に Null になるため any）
	AdjustmentVolume          any    `json: "AdjustmentVolume"`          // 調整済み取引高（売買がない場合に Null になるため any）
	// MorningOpen               string `json: "MorningOpen"`               // 前場始値
	// MorningHigh               string `json: "MorningHigh"`               // 前場高値
	// MorningLow                string `json: "MorningLow"`                // 前場安値
	// MorningClose              string `json: "MorningClose"`              // 前場終値
	// MorningUpperLimit         string `json: "MorningUpperLimit"`         // 前場ストップ高を記録したか、否かを表すフラグ
	// MorningLowerLimit         string `json: "MorningLowerLimit"`         // 前場ストップ安を記録したか、否かを表すフラグ
	// MorningVolume             string `json: "MorningVolume"`             // 前場売買高
	// MorningTurnoverValue      string `json: "MorningTurnoverValue"`      // 前場取引代金
	// MorningAdjustmentOpen     string `json: "MorningAdjustmentOpen"`     // 調整済み前場始値
	// MorningAdjustmentHigh     string `json: "MorningAdjustmentHigh"`     // 調整済み前場高値
	// MorningAdjustmentLow      string `json: "MorningAdjustmentLow"`      // 調整済み前場安値
	// MorningAdjustmentClose    string `json: "MorningAdjustmentClose"`    // 調整済み前場終値
	// MorningAdjustmentVolume   string `json: "MorningAdjustmentVolume"`   // 調整済み前場売買高
	// AfternoonOpen             string `json: "AfternoonOpen"`             // 後場始値
	// AfternoonHigh             string `json: "AfternoonHigh"`             // 後場高値
	// AfternoonLow              string `json: "AfternoonLow"`              // 後場安値
	// AfternoonClose            string `json: "AfternoonClose"`            // 後場終値
	// AfternoonUpperLimit       string `json: "AfternoonUpperLimit"`       // 後場ストップ高を記録したか、否かを表すフラグ
	// AfternoonLowerLimit       string `json: "AfternoonLowerLimit"`       // 後場ストップ安を記録したか、否かを表すフラグ
	// AfternoonVolume           string `json: "AfternoonVolume"`           // 後場売買高
	// AfternoonTurnoverValue    string `json: "AfternoonTurnoverValue"`    // 後場取引代金
	// AfternoonAdjustmentOpen   string `json: "AfternoonAdjustmentOpen"`   // 調整済み後場始値
	// AfternoonAdjustmentHigh   string `json: "AfternoonAdjustmentHigh"`   // 調整済み後場高値
	// AfternoonAdjustmentLow    string `json: "AfternoonAdjustmentLow"`    // 調整済み後場安値
	// AfternoonAdjustmentClose  string `json: "AfternoonAdjustmentClose"`  // 調整済み後場終値
	// AfternoonAdjustmentVolume string `json: "AfternoonAdjustmentVolume"` // 調整済み後場売買高
}
