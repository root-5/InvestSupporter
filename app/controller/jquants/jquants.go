// JQuants API を利用するための関数をまとめたパッケージ
package jquants

import (
	log "app/controller/log"
	model "app/domain/model"
	"fmt"
	"os"
	"time"
)

// ====================================================================================
// API関数
// ====================================================================================

/*
JQuants に登録したメールアドレスとパスワードを入力して、リフレッシュトークン（期限: 1週間）を取得する関数
- arg) refreshToken	リフレッシュトークン
- arg) err			エラー
*/
func getRefreshToken() (err error) {
	// 環境変数からメールアドレスとパスワードを取得
	email := os.Getenv("JQUANTS_EMAIL")
	pass := os.Getenv("JQUANTS_PASS")

	// 環境変数からリフレッシュトークンと前回取得時刻を取得
	refreshToken = os.Getenv("JQUANTS_REFRESH_TOKEN")
	refreshTokenTime, _ := time.Parse(time.RFC3339, os.Getenv("JQUANTS_REFRESH_TOKEN_TIME"))

	// リフレッシュトークンが存在し、取得時刻から1週間以内の場合はリフレッシュトークンを返す
	if refreshToken != "" && time.Since(refreshTokenTime) < 7*24*time.Hour {
		return nil
	}

	// リクエスト先URL
	url := "https://api.jquants.com/v1/token/auth_user"

	// クエリパラメータ定義
	type queryParamsType struct{}
	queryParams := queryParamsType{}

	// リクエストボディ
	type reqBodyType struct {
		Mailaddress string `json:"mailaddress"`
		Password    string `json:"password"`
	}
	reqBody := reqBodyType{
		Mailaddress: email,
		Password:    pass,
	}

	// レスポンスボディ定義
	type resBodyType struct {
		RefreshToken string `json:"refreshToken"`
	}
	var resBody resBodyType

	// POSTリクエスト
	err = post(url, queryParams, reqBody, &resBody)
	if err != nil {
		log.Error(err)
		return err
	}

	// リフレッシュトークンを取得
	refreshToken = resBody.RefreshToken

	// リフレッシュトークンと現在時刻を環境変数に保存
	os.Setenv("JQUANTS_REFRESH_TOKEN", refreshToken)
	os.Setenv("JQUANTS_REFRESH_TOKEN_TIME", time.Now().Format(time.RFC3339))

	return nil
}

/*
リフレッシュトークンを渡して、ID トークン（期限: 24時間）を取得する関数
- arg) idToken			ID トークン
- arg) err				エラー
- return) refreshToken	getRefreshToken 関数で取得したトークン
*/
func getIdToken(refreshToken string) (err error) {
	// 環境変数から ID トークンと前回取得時刻を取得
	idToken = os.Getenv("JQUANTS_ID_TOKEN")
	idTokenTime, _ := time.Parse(time.RFC3339, os.Getenv("JQUANTS_ID_TOKEN_TIME"))

	// ID トークンが存在し、取得時刻から24時間以内の場合は ID トークンを返す
	if idToken != "" && time.Since(idTokenTime) < 24*time.Hour {
		return nil
	}

	// リクエスト先URL
	url := "https://api.jquants.com/v1/token/auth_refresh"

	// クエリパラメータ定義
	type queryParamsType struct {
		RefreshToken string
	}
	queryParam := queryParamsType{
		RefreshToken: refreshToken,
	}

	// リクエストボディ
	type reqBodyType struct{}
	reqBody := reqBodyType{}

	// レスポンスボディ定義
	type resBodyStruct struct {
		IdToken string `json:"idToken"`
	}
	var resBody resBodyStruct

	// POSTリクエスト
	err = post(url, queryParam, reqBody, &resBody)
	if err != nil {
		log.Error(err)
		return err
	}

	// IDトークンを取得
	idToken = resBody.IdToken

	// IDトークンと現在時刻を環境変数に保存
	os.Setenv("JQUANTS_ID_TOKEN", idToken)
	os.Setenv("JQUANTS_ID_TOKEN_TIME", time.Now().Format(time.RFC3339))

	// テスト用のグローバル変数にも保存
	IdTokenForTest = idToken

	return nil
}

/*
リフレッシュトークンを取得した上でIDトークンを取得する関数
- arg) err	エラー
*/
func setIdToken() (err error) {
	// リフレッシュトークンを取得
	err = getRefreshToken()
	if err != nil {
		log.Error(err)
		return err
	}

	// ID トークンを取得
	err = getIdToken(refreshToken)
	if err != nil {
		log.Error(err)
		return err
	}

	// fmt.Println(">> ID Token: " + idToken)

	return nil
}

/*
上場銘柄一覧を取得する関数
- arg) stocksList	上場銘柄情報の配列
*/
func GetStocksInfo() (stocksList []model.StocksInfo, err error) {
	// リクエスト先URL
	url := "https://api.jquants.com/v1/listed/info"

	// クエリパラメータ定義
	type queryParamsType struct{}
	queryParams := queryParamsType{}

	// ヘッダー定義
	type headersType struct {
		Authorization string `json:"Authorization"`
	}
	headers := headersType{
		Authorization: idToken,
	}

	// レスポンスボディ定義
	type resBodyStruct struct {
		Info []jquantsStockInfo `json:"info"`
	}
	var resBody resBodyStruct

	// GETリクエスト
	err = get(url, queryParams, headers, &resBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 型変換（jquantsStockInfo 型の配列から model.StockInfo 型の配列に変換）
	for _, stock := range resBody.Info {
		stocksList = append(stocksList, model.StocksInfo{
			Code:               stock.Code,
			CompanyName:        stock.CompanyName,
			CompanyNameEnglish: stock.CompanyNameEnglish,
			Sector17Code:       convertStringToInt(stock.Sector17Code),
			Sector33Code:       convertStringToInt(stock.Sector33Code),
			ScaleCategory:      stock.ScaleCategory,
			MarketCode:         convertStringToInt(stock.MarketCode),
		})
	}

	return stocksList, nil
}

/*
企業の財務情報を取得する関数
- arg) codeOrDate		銘柄コードまたは日付（YYYY-MM-DD）
- return) financialInfo	企業の財務情報
- return) err			エラー
*/
func GetFinancialInfo(codeOrDate string) (financialInfo []model.FinancialInfo, err error) {
	// リクエスト先URL
	url := "https://api.jquants.com/v1/financial/info"

	// クエリパラメータ定義
	type queryParamsType struct {
		Code string
	}
	queryParams := queryParamsType{
		Code: fmt.Sprint(codeOrDate),
	}

	// ヘッダー定義
	type headersType struct {
		Authorization string `json:"Authorization"`
	}
	headers := headersType{
		Authorization: idToken,
	}

	// レスポンスボディ定義
	type resBodyStruct struct {
		Statements []jquantsFinancialInfo `json:"statements"`
	}
	var resBody resBodyStruct

	// GETリクエスト
	err = get(url, queryParams, headers, &resBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 型変換（jquantsFinancialInfo 型の配列から model.FinancialInfo 型の配列に変換）
	for _, state := range resBody.Statements {
		financialInfo = append(financialInfo, model.FinancialInfo{
			Code:                                   state.Code,
			DisclosedDate:                          convertStringToTime(state.DisclosedDate),
			DisclosedTime:                          convertStringToTime(state.DisclosedTime),
			NetSales:                               convertStringToInt(state.NetSales),
			OperatingProfit:                        convertStringToInt(state.OperatingProfit),
			OrdinaryProfit:                         convertStringToInt(state.OrdinaryProfit),
			Profit:                                 convertStringToInt(state.Profit),
			EarningsPerShare:                       convertStringToFloat64(state.EarningsPerShare),
			TotalAssets:                            convertStringToInt(state.TotalAssets),
			Equity:                                 convertStringToInt(state.Equity),
			EquityToAssetRatio:                     convertStringToFloat64(state.EquityToAssetRatio),
			BookValuePerShare:                      convertStringToFloat64(state.BookValuePerShare),
			CashFlowsFromOperatingActivities:       convertStringToInt(state.CashFlowsFromOperatingActivities),
			CashFlowsFromInvestingActivities:       convertStringToInt(state.CashFlowsFromInvestingActivities),
			CashFlowsFromFinancingActivities:       convertStringToInt(state.CashFlowsFromFinancingActivities),
			CashAndEquivalents:                     convertStringToInt(state.CashAndEquivalents),
			ResultDividendPerShareAnnual:           convertStringToFloat64(state.ResultDividendPerShareAnnual),
			ResultPayoutRatioAnnual:                convertStringToFloat64(state.ResultPayoutRatioAnnual),
			ForecastDividendPerShareAnnual:         convertStringToFloat64(state.ForecastDividendPerShareAnnual),
			ForecastPayoutRatioAnnual:              convertStringToFloat64(state.ForecastPayoutRatioAnnual),
			NextYearForecastDividendPerShareAnnual: convertStringToFloat64(state.NextYearForecastDividendPerShareAnnual),
			NextYearForecastPayoutRatioAnnual:      convertStringToFloat64(state.NextYearForecastPayoutRatioAnnual),
			ForecastNetSales:                       convertStringToInt(state.ForecastNetSales),
			ForecastOperatingProfit:                convertStringToInt(state.ForecastOperatingProfit),
			ForecastOrdinaryProfit:                 convertStringToInt(state.ForecastOrdinaryProfit),
			ForecastProfit:                         convertStringToInt(state.ForecastProfit),
			ForecastEarningsPerShare:               convertStringToFloat64(state.ForecastEarningsPerShare),
			NextYearForecastNetSales:               convertStringToInt(state.NextYearForecastNetSales),
			NextYearForecastOperatingProfit:        convertStringToInt(state.NextYearForecastOperatingProfit),
			NextYearForecastOrdinaryProfit:         convertStringToInt(state.NextYearForecastOrdinaryProfit),
			NextYearForecastProfit:                 convertStringToInt(state.NextYearForecastProfit),
			NextYearForecastEarningsPerShare:       convertStringToFloat64(state.NextYearForecastEarningsPerShare),
			NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock: convertStringToInt(state.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock),
		})
	}

	return financialInfo, nil
}
