// JQuants API を利用するための関数をまとめたパッケージ
package jquants

import (
	log "app/controller/log"
	model "app/domain/model"
	"os"
	"reflect"
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
- arg) IdToken			ID トークン
- arg) err				エラー
- return) refreshToken	getRefreshToken 関数で取得したトークン
*/
func getIdToken(refreshToken string) (err error) {
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

	// fmt.Println(">> ID Token: " + IdToken)

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
			CompanyName:        stock.CompanyName,
			CompanyNameEnglish: stock.CompanyNameEnglish,
			Sector17Code:       convertStringToInt64(stock.Sector17Code),
			Sector33Code:       convertStringToInt64(stock.Sector33Code),
			ScaleCategory:      stock.ScaleCategory,
			MarketCode:         convertStringToInt64(stock.MarketCode),
		})
	}

	return stocksList, nil
}

/*
企業の財務情報を取得する関数
- arg) codeOrDate		銘柄コードまたは日付（YYYY-MM-DD）
- return) financials	企業の財務情報
- return) err			エラー
*/
func GetFinancialInfo(codeOrDate string) (financials []model.FinancialInfo, err error) {
	// リクエスト先URL
	url := "https://api.jquants.com/v1/fins/statements"

	// クエリパラメータ定義
	type queryParamsType struct {
		Code string
		Date string
	}
	var queryParams = queryParamsType{}

	// もしcodeOrDateがコードの場合は融合処理を行いデータをまとめる
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
		Statements []jquantsFinancialInfo `json:"statements"`
	}
	var resBody resBodyStruct

	// GETリクエスト
	err = get(url, queryParams, headers, &resBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, state := range resBody.Statements {
		// state の中身が空の場合は Code のみ入った構造体を返却
		if state.Code == "" {
			financials = append(financials, model.FinancialInfo{
				Code: codeOrDate,
			})
		} else {
			// 型変換（jquantsFinancialInfo 型の配列から model.FinancialInfo 型の配列に変換）
			financials = append(financials, model.FinancialInfo{
				Code:                                   state.Code,
				DisclosedDate:                          convertStringToTime(state.DisclosedDate),
				DisclosedTime:                          convertStringToTime(state.DisclosedTime),
				NetSales:                               convertStringToInt64(state.NetSales),
				OperatingProfit:                        convertStringToInt64(state.OperatingProfit),
				OrdinaryProfit:                         convertStringToInt64(state.OrdinaryProfit),
				Profit:                                 convertStringToInt64(state.Profit),
				EarningsPerShare:                       convertStringToFloat64(state.EarningsPerShare),
				TotalAssets:                            convertStringToInt64(state.TotalAssets),
				Equity:                                 convertStringToInt64(state.Equity),
				EquityToAssetRatio:                     convertStringToFloat64(state.EquityToAssetRatio),
				BookValuePerShare:                      convertStringToFloat64(state.BookValuePerShare),
				CashFlowsFromOperatingActivities:       convertStringToInt64(state.CashFlowsFromOperatingActivities),
				CashFlowsFromInvestingActivities:       convertStringToInt64(state.CashFlowsFromInvestingActivities),
				CashFlowsFromFinancingActivities:       convertStringToInt64(state.CashFlowsFromFinancingActivities),
				CashAndEquivalents:                     convertStringToInt64(state.CashAndEquivalents),
				ResultDividendPerShareAnnual:           convertStringToFloat64(state.ResultDividendPerShareAnnual),
				ResultPayoutRatioAnnual:                convertStringToFloat64(state.ResultPayoutRatioAnnual),
				ForecastDividendPerShareAnnual:         convertStringToFloat64(state.ForecastDividendPerShareAnnual),
				ForecastPayoutRatioAnnual:              convertStringToFloat64(state.ForecastPayoutRatioAnnual),
				NextYearForecastDividendPerShareAnnual: convertStringToFloat64(state.NextYearForecastDividendPerShareAnnual),
				NextYearForecastPayoutRatioAnnual:      convertStringToFloat64(state.NextYearForecastPayoutRatioAnnual),
				ForecastNetSales:                       convertStringToInt64(state.ForecastNetSales),
				ForecastOperatingProfit:                convertStringToInt64(state.ForecastOperatingProfit),
				ForecastOrdinaryProfit:                 convertStringToInt64(state.ForecastOrdinaryProfit),
				ForecastProfit:                         convertStringToInt64(state.ForecastProfit),
				ForecastEarningsPerShare:               convertStringToFloat64(state.ForecastEarningsPerShare),
				NextYearForecastNetSales:               convertStringToInt64(state.NextYearForecastNetSales),
				NextYearForecastOperatingProfit:        convertStringToInt64(state.NextYearForecastOperatingProfit),
				NextYearForecastOrdinaryProfit:         convertStringToInt64(state.NextYearForecastOrdinaryProfit),
				NextYearForecastProfit:                 convertStringToInt64(state.NextYearForecastProfit),
				NextYearForecastEarningsPerShare:       convertStringToFloat64(state.NextYearForecastEarningsPerShare),
				NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock: convertStringToInt64(state.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock),
			})
		}
	}

	// financials の中身が空の場合は Code のみ入った構造体を返却
	if len(financials) == 0 {
		financial := model.FinancialInfo{
			Code: codeOrDate,
		}
		// financials をスライスに変換して返す
		financials = []model.FinancialInfo{financial}
	}

	// もしcodeOrDateがコードの場合は融合処理を行いデータをまとめる
	if len(codeOrDate) == 4 || len(codeOrDate) == 5 {
		// 統合後の財務情報
		var financialsMerged model.FinancialInfo

		// APIから返却される内容は古いものから順になっているので、配列の最初の要素から順に処理する
		for _, financial := range financials {
			// 初回は統合後の財務情報にそのまま代入
			if financialsMerged.Code == "" {
				financialsMerged = financial
			} else {
				// 2回目以降は統合処理を行う、ただし新しいデータがない（nil）の場合はスキップ
				m := reflect.ValueOf(financial)
				merged := reflect.ValueOf(&financialsMerged).Elem()

				// フィールドごとに統合処理を行う
				for i := 0; i < m.NumField(); i++ {
					switch m.Field(i).Type().String() {
					case "sql.NullInt64":
						// フィールドが sql.NullInt64 型の場合
						if m.Field(i).FieldByName("Valid").Bool() {
							merged.Field(i).FieldByName("Int64").Set(m.Field(i).FieldByName("Int64"))
							merged.Field(i).FieldByName("Valid").Set(m.Field(i).FieldByName("Valid"))
						}
					case "sql.NullFloat64":
						// フィールドが sql.NullFloat64 型の場合
						if m.Field(i).FieldByName("Valid").Bool() {
							merged.Field(i).FieldByName("Float64").Set(m.Field(i).FieldByName("Float64"))
							merged.Field(i).FieldByName("Valid").Set(m.Field(i).FieldByName("Valid"))
						}
					case "sql.NullString":
						// フィールドが sql.NullString 型の場合
						if m.Field(i).FieldByName("Valid").Bool() {
							merged.Field(i).FieldByName("String").Set(m.Field(i).FieldByName("String"))
							merged.Field(i).FieldByName("Valid").Set(m.Field(i).FieldByName("Valid"))
						}
					case "sql.NullTime":
						// フィールドが sql.NullTime 型の場合
						if m.Field(i).FieldByName("Valid").Bool() {
							merged.Field(i).FieldByName("Time").Set(m.Field(i).FieldByName("Time"))
							merged.Field(i).FieldByName("Valid").Set(m.Field(i).FieldByName("Valid"))
						}
					default:
						if m.Field(i).Interface() != nil {
							merged.Field(i).Set(m.Field(i))
						}
					}
				}
			}
		}

		// 統合前の財務情報を初期化しなおして、統合後の財務情報を返却する
		financials = make([]model.FinancialInfo, 1)
		financials[0] = financialsMerged
	}

	return financials, nil
}
