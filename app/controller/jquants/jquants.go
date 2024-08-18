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
// 初期化関数
// ====================================================================================
func Init() {
	schedulerStart()
}

// ====================================================================================
// API関数
// ====================================================================================

/* JQuants に登録したメールアドレスとパスワードを入力して、リフレッシュトークン（期限: 1週間）を取得する関数
> refreshToken	リフレッシュトークン
> err			エラー
*/
func getRefreshToken() (refreshToken string, err error) {
	// 環境変数からメールアドレスとパスワードを取得
	email := os.Getenv("JQUANTS_EMAIL")
	pass := os.Getenv("JQUANTS_PASS")

	// 環境変数からリフレッシュトークンと前回取得時刻を取得
	refreshToken = os.Getenv("JQUANTS_REFRESH_TOKEN")
	refreshTokenTime, _ := time.Parse(time.RFC3339, os.Getenv("JQUANTS_REFRESH_TOKEN_TIME"))

	// リフレッシュトークンが存在し、取得時刻から1週間以内の場合はリフレッシュトークンを返す
	if refreshToken != "" && time.Since(refreshTokenTime) < 7*24*time.Hour {
		return refreshToken, nil
	}

	// リクエスト先URL
	url := "https://api.jquants.com/v1/token/auth_user"

	// クエリパラメータ定義
	type queryParamsType struct {}
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
		return "", err
	}

	// リフレッシュトークンを取得
	refreshToken = resBody.RefreshToken

	// リフレッシュトークンと現在時刻を環境変数に保存
	os.Setenv("JQUANTS_REFRESH_TOKEN", refreshToken)
	os.Setenv("JQUANTS_REFRESH_TOKEN_TIME", time.Now().Format(time.RFC3339))

	return refreshToken, nil
}

/* リフレッシュトークンを渡して、ID トークン（期限: 24時間）を取得する関数
	- refreshToken	getRefreshToken 関数で取得したトークン
	> idToken		ID トークン
	> err			エラー
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
	queryParam := queryParamsType {
		RefreshToken: refreshToken,
	}

	// リクエストボディ
	type reqBodyType struct {}
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

	return nil
}

/* リフレッシュトークンを取得した上でIDトークンを取得する関数
	> err	エラー
*/
func setIdToken() (err error) {
	// リフレッシュトークンを取得
	refreshToken, err := getRefreshToken()
	if err != nil {
		return err
	}

	// ID トークンを取得
	err = getIdToken(refreshToken)
	if err != nil {
		return err
	}

	fmt.Println(">> ID Token: " + idToken)

	return nil
}

/* 上場銘柄一覧を取得する関数
	> stocksList	上場銘柄情報の配列
*/
func GetStocksInfo() (stocksList []model.StocksInfo, err error) {
	// リクエスト先URL
	url := "https://api.jquants.com/v1/listed/info"

	// クエリパラメータ定義
	type queryParamsType struct {}
	queryParams := queryParamsType{}

	// ヘッダー定義
	type headersType struct {
		Authorization string `json:"Authorization"`
	}
	headers := headersType {
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
			Code:              stock.Code,
			CompanyName:       stock.CompanyName,
			CompanyNameEnglish: stock.CompanyNameEnglish,
			Sector17Code:      convertStringToInt(stock.Sector17Code),
			Sector33Code:      convertStringToInt(stock.Sector33Code),
			ScaleCategory:     stock.ScaleCategory,
			MarketCode:        convertStringToInt(stock.MarketCode),
		})
	}

	return stocksList, nil
}
