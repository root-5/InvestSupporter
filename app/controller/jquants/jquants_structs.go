package jquants

// ====================================================================================
// 各APIのレスポンスボディの構造体を定義
// ====================================================================================

// 上場銘柄一覧
type jquantsStockInfo struct {
	// Date               string `json:"Date"`               // 日付
	Code               string `json:"Code"`     // 銘柄コード
	CompanyName        string `json:"CoName"`   // 会社名
	CompanyNameEnglish string `json:"CoNameEn"` // 会社名（英語）
	Sector17Code       string `json:"S17"`      // 17業種コード
	// Sector17CodeName   string `json:"S17Nm"`              // 17業種コード名
	Sector33Code string `json:"S33"` // 33業種コード
	// Sector33CodeName   string `json:"S33Nm"`              // 33業種コード名
	ScaleCategory string `json:"ScaleCat"` // 規模コード
	MarketCode    string `json:"Mkt"`      // 市場区分コード
	// MarketCodeName     string `json:"MktNm"`              // 市場区分名
	// MarginCode         string `json:"Mrgn"`               // 貸借信用区分
	// MarginCodeName     string `json:"MrgnNm"`             // 貸借信用区分名
}

// 財務情報
type jquantsStatementInfo struct {
	DisclosedDate string `json:"DiscDate"` // 開示日
	// DisclosedTime                                                                string `json:"DiscTime"`       // 開示時刻
	Code             string `json:"Code"`    // 銘柄コード（5桁）
	DisclosureNumber string `json:"DiscNo"`  // 開示番号
	TypeOfDocument   string `json:"DocType"` // 開示書類種別
	// TypeOfCurrentPeriod                                                          string `json:"CurPerType"`     // 当会計期間の種類
	// CurrentPeriodStartDate                                                       string `json:"CurPerSt"`       // 当会計期間開始日
	// CurrentPeriodEndDate                                                         string `json:"CurPerEn"`       // 当会計期間終了日
	// CurrentFiscalYearStartDate                                                   string `json:"CurFYSt"`        // 当事業年度開始日
	// CurrentFiscalYearEndDate                                                     string `json:"CurFYEn"`        // 当事業年度終了日
	// NextFiscalYearStartDate                                                      string `json:"NxtFYSt"`        // 翌事業年度開始日
	// NextFiscalYearEndDate                                                        string `json:"NxtFYEn"`        // 翌事業年度終了日
	NetSales         string `json:"Sales"` // 売上高
	OperatingProfit  string `json:"OP"`    // 営業利益
	OrdinaryProfit   string `json:"OdP"`   // 経常利益
	Profit           string `json:"NP"`    // 当期純利益
	EarningsPerShare string `json:"EPS"`   // 一株あたり当期純利益
	// DilutedEarningsPerShare                                                      string `json:"DEPS"`           // 潜在株式調整後一株あたり当期純利益
	TotalAssets                      string `json:"TA"`     // 総資産
	Equity                           string `json:"Eq"`     // 純資産
	EquityToAssetRatio               string `json:"EqAR"`   // 自己資本比率
	BookValuePerShare                string `json:"BPS"`    // 一株あたり純資産
	CashFlowsFromOperatingActivities string `json:"CFO"`    // 営業活動によるキャッシュ・フロー
	CashFlowsFromInvestingActivities string `json:"CFI"`    // 投資活動によるキャッシュ・フロー
	CashFlowsFromFinancingActivities string `json:"CFF"`    // 財務活動によるキャッシュ・フロー
	CashAndEquivalents               string `json:"CashEq"` // 現金及び現金同等物期末残高
	// ResultDividendPerShare1stQuarter                                             string `json:"Div1Q"`          // 一株あたり配当実績_第1四半期末
	// ResultDividendPerShare2ndQuarter                                             string `json:"Div2Q"`          // 一株あたり配当実績_第2四半期末
	// ResultDividendPerShare3rdQuarter                                             string `json:"Div3Q"`          // 一株あたり配当実績_第3四半期末
	// ResultDividendPerShareFiscalYearEnd                                          string `json:"DivFY"`          // 一株あたり配当実績_期末
	ResultDividendPerShareAnnual string `json:"DivAnn"` // 一株あたり配当実績_合計
	// DistributionsPerUnitREIT                                                     string `json:"DivUnit"`        // 1口当たり分配金
	ResultTotalDividendPaidAnnual string `json:"DivTotalAnn"`    // 配当金総額
	ResultPayoutRatioAnnual       string `json:"PayoutRatioAnn"` // 配当性向
	// ForecastDividendPerShare1stQuarter                                           string `json:"FDiv1Q"`         // 一株あたり配当予想_第1四半期末
	// ForecastDividendPerShare2ndQuarter                                           string `json:"FDiv2Q"`         // 一株あたり配当予想_第2四半期末
	// ForecastDividendPerShare3rdQuarter                                           string `json:"FDiv3Q"`         // 一株あたり配当予想_第3四半期末
	// ForecastDividendPerShareFiscalYearEnd                                        string `json:"FDivFY"`         // 一株あたり配当予想_期末
	ForecastDividendPerShareAnnual string `json:"FDivAnn"` // 一株あたり配当予想_合計
	// ForecastDistributionsPerUnitREIT                                             string `json:"FDivUnit"`       // 1口当たり予想分配金
	ForecastTotalDividendPaidAnnual string `json:"FDivTotalAnn"`    // 予想配当金総額
	ForecastPayoutRatioAnnual       string `json:"FPayoutRatioAnn"` // 予想配当性向
	// NextYearForecastDividendPerShare1stQuarter                                   string `json:"NxFDiv1Q"`       // 一株あたり配当予想_翌事業年度第1四半期末
	// NextYearForecastDividendPerShare2ndQuarter                                   string `json:"NxFDiv2Q"`       // 一株あたり配当予想_翌事業年度第2四半期末
	// NextYearForecastDividendPerShare3rdQuarter                                   string `json:"NxFDiv3Q"`       // 一株あたり配当予想_翌事業年度第3四半期末
	// NextYearForecastDividendPerShareFiscalYearEnd                                string `json:"NxFDivFY"`       // 一株あたり配当予想_翌事業年度期末
	NextYearForecastDividendPerShareAnnual string `json:"NxFDivAnn"` // 一株あたり配当予想_翌事業年度合計
	// NextYearForecastDistributionsPerUnitREIT                                     string `json:"NxFDivUnit"`     // 1口当たり翌事業年度予想分配金
	NextYearForecastPayoutRatioAnnual string `json:"NxFPayoutRatioAnn"` // 翌事業年度予想配当性向
	// ForecastNetSales2ndQuarter                                                   string `json:"FSales2Q"`       // 売上高_予想_第2四半期末
	// ForecastOperatingProfit2ndQuarter                                            string `json:"FOP2Q"`          // 営業利益_予想_第2四半期末
	// ForecastOrdinaryProfit2ndQuarter                                             string `json:"FOdP2Q"`         // 経常利益_予想_第2四半期末
	// ForecastProfit2ndQuarter                                                     string `json:"FNP2Q"`          // 当期純利益_予想_第2四半期末
	// ForecastEarningsPerShare2ndQuarter                                           string `json:"FEPS2Q"`         // 一株あたり当期純利益_予想_第2四半期末
	// NextYearForecastNetSales2ndQuarter                                           string `json:"NxFSales2Q"`     // 売上高_予想_翌事業年度第2四半期末
	// NextYearForecastOperatingProfit2ndQuarter                                    string `json:"NxFOP2Q"`        // 営業利益_予想_翌事業年度第2四半期末
	// NextYearForecastOrdinaryProfit2ndQuarter                                     string `json:"NxFOdP2Q"`       // 経常利益_予想_翌事業年度第2四半期末
	// NextYearForecastProfit2ndQuarter                                             string `json:"NxFNp2Q"`        // 当期純利益_予想_翌事業年度第2四半期末
	// NextYearForecastEarningsPerShare2ndQuarter                                   string `json:"NxFEPS2Q"`       // 一株あたり当期純利益_予想_翌事業年度第2四半期末
	ForecastNetSales                                                             string `json:"FSales"`   // 売上高_予想_期末
	ForecastOperatingProfit                                                      string `json:"FOP"`      // 営業利益_予想_期末
	ForecastOrdinaryProfit                                                       string `json:"FOdP"`     // 経常利益_予想_期末
	ForecastProfit                                                               string `json:"FNP"`      // 当期純利益_予想_期末
	ForecastEarningsPerShare                                                     string `json:"FEPS"`     // 一株あたり当期純利益_予想_期末
	NextYearForecastNetSales                                                     string `json:"NxFSales"` // 売上高_予想_翌事業年度期末
	NextYearForecastOperatingProfit                                              string `json:"NxFOP"`    // 営業利益_予想_翌事業年度期末
	NextYearForecastOrdinaryProfit                                               string `json:"NxFOdP"`   // 経常利益_予想_翌事業年度期末
	NextYearForecastProfit                                                       string `json:"NxFNp"`    // 当期純利益_予想_翌事業年度期末
	NextYearForecastEarningsPerShare                                             string `json:"NxFEPS"`   // 一株あたり当期純利益_予想_翌事業年度期末
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock string `json:"ShOutFY"`  // 期末発行済株式数
}

// 株価四本値
type jquantsPriceInfo struct {
	Date string `json:"Date"` // 日付
	Code string `json:"Code"` // 銘柄コード
	// Open                      string  `json:"Open"`                       // 始値（調整前）
	// High                      string  `json:"High"`                       // 高値（調整前）
	// Low                       string  `json:"Low"`                        // 安値（調整前）
	// Close                     string  `json:"Close"`                      // 終値（調整前）
	// UpperLimit                string  `json:"UpperLimit"`                 // 日通ストップ高を記録したか、否かを表すフラグ
	// LowerLimit                string  `json:"LowerLimit"`                 // 日通ストップ安を記録したか、否かを表すフラグ
	// Volume                    string  `json:"Volume"`                     // 取引高（調整前）
	// TurnoverValue             string  `json:"TurnoverValue"`              // 取引代金
	AdjustmentFactor float64 `json:"AdjFactor"` // 調整係数
	AdjustmentOpen   any     `json:"AdjO"`      // 調整済み始値（売買がない場合に Null になるため any）
	AdjustmentHigh   any     `json:"AdjH"`      // 調整済み高値（売買がない場合に Null になるため any）
	AdjustmentLow    any     `json:"AdjL"`      // 調整済み安値（売買がない場合に Null になるため any）
	AdjustmentClose  any     `json:"AdjC"`      // 調整済み終値（売買がない場合に Null になるため any）
	AdjustmentVolume any     `json:"AdjVo"`     // 調整済み取引高（売買がない場合に Null になるため any）
	// MorningOpen               string  `json:"MorningOpen"`                // 前場始値
	// MorningHigh               string  `json:"MorningHigh"`                // 前場高値
	// MorningLow                string  `json:"MorningLow"`                 // 前場安値
	// MorningClose              string  `json:"MorningClose"`               // 前場終値
	// MorningUpperLimit         string  `json:"MorningUpperLimit"`          // 前場ストップ高を記録したか、否かを表すフラグ
	// MorningLowerLimit         string  `json:"MorningLowerLimit"`          // 前場ストップ安を記録したか、否かを表すフラグ
	// MorningVolume             string  `json:"MorningVolume"`              // 前場売買高
	// MorningTurnoverValue      string  `json:"MorningTurnoverValue"`       // 前場取引代金
	// MorningAdjustmentOpen     string  `json:"MorningAdjustmentOpen"`      // 調整済み前場始値
	// MorningAdjustmentHigh     string  `json:"MorningAdjustmentHigh"`      // 調整済み前場高値
	// MorningAdjustmentLow      string  `json:"MorningAdjustmentLow"`       // 調整済み前場安値
	// MorningAdjustmentClose    string  `json:"MorningAdjustmentClose"`     // 調整済み前場終値
	// MorningAdjustmentVolume   string  `json:"MorningAdjustmentVolume"`    // 調整済み前場売買高
	// AfternoonOpen             string  `json:"AfternoonOpen"`              // 後場始値
	// AfternoonHigh             string  `json:"AfternoonHigh"`              // 後場高値
	// AfternoonLow              string  `json:"AfternoonLow"`               // 後場安値
	// AfternoonClose            string  `json:"AfternoonClose"`             // 後場終値
	// AfternoonUpperLimit       string  `json:"AfternoonUpperLimit"`        // 後場ストップ高を記録したか、否かを表すフラグ
	// AfternoonLowerLimit       string  `json:"AfternoonLowerLimit"`        // 後場ストップ安を記録したか、否かを表すフラグ
	// AfternoonVolume           string  `json:"AfternoonVolume"`            // 後場売買高
	// AfternoonTurnoverValue    string  `json:"AfternoonTurnoverValue"`     // 後場取引代金
	// AfternoonAdjustmentOpen   string  `json:"AfternoonAdjustmentOpen"`    // 調整済み後場始値
	// AfternoonAdjustmentHigh   string  `json:"AfternoonAdjustmentHigh"`    // 調整済み後場高値
	// AfternoonAdjustmentLow    string  `json:"AfternoonAdjustmentLow"`     // 調整済み後場安値
	// AfternoonAdjustmentClose  string  `json:"AfternoonAdjustmentClose"`   // 調整済み後場終値
	// AfternoonAdjustmentVolume string  `json:"AfternoonAdjustmentVolume"`  // 調整済み後場売買高
}
