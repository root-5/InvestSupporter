// JQuants API を利用するための関数をまとめたパッケージ
package jquants

import (
	"app/controller/log"
	"app/domain/model"
	"os"
	"strconv"
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
func fetchRefreshToken() (err error) {
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
- arg) IdToken			ID トークン
- arg) err				エラー
- return) refreshToken	getRefreshToken 関数で取得したトークン
*/
func fetchIdToken(refreshToken string) (err error) {
	// 環境変数から ID トークンと前回取得時刻を取得
	IdToken = os.Getenv("JQUANTS_ID_TOKEN")
	idTokenTime, _ := time.Parse(time.RFC3339, os.Getenv("JQUANTS_ID_TOKEN_TIME"))

	// ID トークンが存在し、取得時刻から24時間以内の場合は ID トークンを返す
	if IdToken != "" && time.Since(idTokenTime) < 24*time.Hour {
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
	IdToken = resBody.IdToken

	// IDトークンと現在時刻を環境変数に保存
	os.Setenv("JQUANTS_ID_TOKEN", IdToken)
	os.Setenv("JQUANTS_ID_TOKEN_TIME", time.Now().Format(time.RFC3339))

	return nil
}

/*
リフレッシュトークンを取得した上でIDトークンを取得する関数
- arg) err	エラー
*/
func setIdToken() (err error) {
	// リフレッシュトークンを取得
	err = fetchRefreshToken()
	if err != nil {
		log.Error(err)
		return err
	}

	// ID トークンを取得
	err = fetchIdToken(refreshToken)
	if err != nil {
		log.Error(err)
		return err
	}

	// fmt.Println(">> ID Token: " + IdToken)

	return nil
}

/*
上場銘柄一覧を取得する関数
- arg) stocksList	上場銘柄情報の配列
*/
func FetchStocksInfo() (stocksList []model.StocksInfo, err error) {
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
		Authorization: IdToken,
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
			CompanyName:        convertStringToString(stock.CompanyName),
			CompanyNameEnglish: convertStringToString(stock.CompanyNameEnglish),
			Sector17Code:       convertStringToInt64(stock.Sector17Code),
			Sector33Code:       convertStringToInt64(stock.Sector33Code),
			ScaleCategory:      convertStringToString(stock.ScaleCategory),
			MarketCode:         convertStringToInt64(stock.MarketCode),
		})
	}

	return stocksList, nil
}

/*
企業の財務情報を取得する関数
- arg) codeOrDate		銘柄コードまたは日付（YYYY-MM-DD）
- return) statements	企業の財務情報
- return) err			エラー
*/
func FetchStatementsInfo(codeOrDate string) (statements []model.StatementInfo, err error) {
	// リクエスト先URL
	url := "https://api.jquants.com/v1/fins/statements"

	// クエリパラメータ定義
	type queryParamsType struct {
		Code string
		Date string
	}
	var queryParams = queryParamsType{}

	// codeOrDate が銘柄コードか日付かでクエリパラメータを変更
	if len(codeOrDate) == 4 || len(codeOrDate) == 5 {
		queryParams = queryParamsType{
			Code: codeOrDate,
		}
	} else {
		queryParams = queryParamsType{
			Date: codeOrDate,
		}
	}

	// ヘッダー定義
	type headersType struct {
		Authorization string `json:"Authorization"`
	}
	headers := headersType{
		Authorization: IdToken,
	}

	// レスポンスボディ定義
	type resBodyStruct struct {
		ResStatements []jquantsStatementInfo `json:"statements"`
	}
	var resBody resBodyStruct

	// GETリクエスト
	err = get(url, queryParams, headers, &resBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resStatement := range resBody.ResStatements {
		// resStatement の中身が空の場合は空の構造体を返却
		if resStatement.Code == "" {
			statements = append(statements, model.StatementInfo{})
		} else {
			// 開示番号を int64 に変換
			disclosureNumber, err := strconv.ParseInt(resStatement.DisclosureNumber, 10, 64)
			if err != nil {
				log.Error(err)
				return nil, err
			}

			// 型変換（jquantsStatementInfo 型の配列から model.StatementInfo 型の配列に変換）
			statements = append(statements, model.StatementInfo{
				DisclosureNumber:                       disclosureNumber,
				Code:                                   resStatement.Code,
				DisclosedDate:                          convertStringToTime(resStatement.DisclosedDate),
				TypeOfDocument:                         resStatement.TypeOfDocument,
				NetSales:                               convertStringToInt64(resStatement.NetSales),
				OperatingProfit:                        convertStringToInt64(resStatement.OperatingProfit),
				OrdinaryProfit:                         convertStringToInt64(resStatement.OrdinaryProfit),
				Profit:                                 convertStringToInt64(resStatement.Profit),
				EarningsPerShare:                       convertStringToFloat64(resStatement.EarningsPerShare),
				TotalAssets:                            convertStringToInt64(resStatement.TotalAssets),
				Equity:                                 convertStringToInt64(resStatement.Equity),
				EquityToAssetRatio:                     convertStringToFloat64(resStatement.EquityToAssetRatio),
				BookValuePerShare:                      convertStringToFloat64(resStatement.BookValuePerShare),
				CashFlowsFromOperatingActivities:       convertStringToInt64(resStatement.CashFlowsFromOperatingActivities),
				CashFlowsFromInvestingActivities:       convertStringToInt64(resStatement.CashFlowsFromInvestingActivities),
				CashFlowsFromFinancingActivities:       convertStringToInt64(resStatement.CashFlowsFromFinancingActivities),
				CashAndEquivalents:                     convertStringToInt64(resStatement.CashAndEquivalents),
				ResultDividendPerShareAnnual:           convertStringToFloat64(resStatement.ResultDividendPerShareAnnual),
				ResultPayoutRatioAnnual:                convertStringToFloat64(resStatement.ResultPayoutRatioAnnual),
				ForecastDividendPerShareAnnual:         convertStringToFloat64(resStatement.ForecastDividendPerShareAnnual),
				ForecastPayoutRatioAnnual:              convertStringToFloat64(resStatement.ForecastPayoutRatioAnnual),
				NextYearForecastDividendPerShareAnnual: convertStringToFloat64(resStatement.NextYearForecastDividendPerShareAnnual),
				NextYearForecastPayoutRatioAnnual:      convertStringToFloat64(resStatement.NextYearForecastPayoutRatioAnnual),
				ForecastNetSales:                       convertStringToInt64(resStatement.ForecastNetSales),
				ForecastOperatingProfit:                convertStringToInt64(resStatement.ForecastOperatingProfit),
				ForecastOrdinaryProfit:                 convertStringToInt64(resStatement.ForecastOrdinaryProfit),
				ForecastProfit:                         convertStringToInt64(resStatement.ForecastProfit),
				ForecastEarningsPerShare:               convertStringToFloat64(resStatement.ForecastEarningsPerShare),
				NextYearForecastNetSales:               convertStringToInt64(resStatement.NextYearForecastNetSales),
				NextYearForecastOperatingProfit:        convertStringToInt64(resStatement.NextYearForecastOperatingProfit),
				NextYearForecastOrdinaryProfit:         convertStringToInt64(resStatement.NextYearForecastOrdinaryProfit),
				NextYearForecastProfit:                 convertStringToInt64(resStatement.NextYearForecastProfit),
				NextYearForecastEarningsPerShare:       convertStringToFloat64(resStatement.NextYearForecastEarningsPerShare),
				NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock: convertStringToInt64(resStatement.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock),
			})
		}
	}
	return statements, nil
}

/*
株価を取得する関数
- arg) codeOrDate			銘柄コードまたは日付（YYYY-MM-DD）
- return) prices			株価情報
- return) splitStockCodes	分割銘柄コード
- return) err				エラー
*/
func FetchPricesInfo(codeOrDate string) (prices []model.PriceInfo, splitStockCodes []string, err error) {
	// リクエスト先URL
	url := "https://api.jquants.com/v1/prices/daily_quotes"

	// クエリパラメータ定義
	type queryParamsType struct {
		Code           string
		Date           string
		Pagination_key string
	}
	var queryParams = queryParamsType{}

	// codeOrDate が銘柄コードか日付かでクエリパラメータを変更
	if len(codeOrDate) == 4 || len(codeOrDate) == 5 {
		queryParams = queryParamsType{
			Code: codeOrDate,
		}
	} else {
		queryParams = queryParamsType{
			Date: codeOrDate,
		}
	}

	// ヘッダー定義
	type headersType struct {
		Authorization string `json:"Authorization"`
	}
	headers := headersType{
		Authorization: IdToken,
	}

	// レスポンスボディ定義
	type resBodyStruct struct {
		Daily_quotes   []jquantsPriceInfo `json:"daily_quotes"`
		Pagination_key string             `json:"pagination_key"`
	}
	var resBody resBodyStruct

	// ページネーションキーが存在する限りループ
	needRoop := true
	lastPaginationKey := ""

	for needRoop {
		lastPaginationKey = resBody.Pagination_key
		queryParams.Pagination_key = resBody.Pagination_key

		// GETリクエスト
		err = get(url, queryParams, headers, &resBody)
		if err != nil {
			log.Error(err)
			return nil, nil, err
		}

		if resBody.Pagination_key == lastPaginationKey {
			needRoop = false
		}

		for _, price := range resBody.Daily_quotes {
			if price.AdjustmentFactor != 1 {
				splitStockCodes = append(splitStockCodes, price.Code)
			}

			// 型変換（jquantsPriceInfo 型の配列から model.StockPrice 型の配列に変換）
			prices = append(prices, model.PriceInfo{
				Date:             price.Date,
				Code:             price.Code,
				AdjustmentOpen:   convertAnyToFloat64(price.AdjustmentOpen),
				AdjustmentHigh:   convertAnyToFloat64(price.AdjustmentHigh),
				AdjustmentLow:    convertAnyToFloat64(price.AdjustmentLow),
				AdjustmentClose:  convertAnyToFloat64(price.AdjustmentClose),
				AdjustmentVolume: convertAnyToFloat64(price.AdjustmentVolume),
			})
		}
	}

	return prices, splitStockCodes, nil
}
