package jquants

import (
	log "app/controller/log"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// ====================================================================================
// 共通変数
// ====================================================================================

// IDトークン
var idToken string

// HTTPクライアント
var httpClient = &http.Client{}

// ====================================================================================
// 共通関数
// ====================================================================================

/* GETリクエストを行い、レスポンスボディを取得する関数
- url			リクエスト先URL文字列
- queryParam	クエリパラメータ構造体
- headers		ヘッダー構造体
- resBody		レスポンスボディ構造体のポインタ
> err			エラー
*/
func get[T any](reqUrl string, queryParams any, headers any, resBody *T) (err error) {
	// クエリパラメータをreqURLに追加
	if queryParams != struct{}{} {
		queryParamVal := reflect.ValueOf(queryParams)
		queryParamType := reflect.TypeOf(queryParams)
		params := url.Values{}
		for i := 0; i < queryParamVal.NumField(); i++ {
			params.Add(strings.ToLower(queryParamType.Field(i).Name), fmt.Sprintf("%v", queryParamVal.Field(i).Interface()))
		}
		reqUrl += "?" + params.Encode()
	}

	// GETリクエスト作成
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Error(err)
		return err
	}

	// ヘッダー設定
	req.Header.Set("Content-Type", "application/json")
	if headers != struct{}{} {
		headerVal := reflect.ValueOf(headers)
		headerType := reflect.TypeOf(headers)
		for i := 0; i < headerVal.NumField(); i++ {
			req.Header.Set(strings.ToLower(headerType.Field(i).Name), fmt.Sprintf("%v", headerVal.Field(i).Interface()))
		}
	}

	// リクエスト送信
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	// レスポンスボディを読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return err
	}

	// レスポンスボディを構造体に変換
	if err := json.Unmarshal(body, resBody); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/* POSTリクエストを行い、レスポンスボディを取得する関数
	- url			リクエスト先URL文字列
	- queryParam	クエリパラメータ構造体
	- reqBody		リクエストボディ構造体
	- resBody		レスポンスボディ構造体のポインタ
	> err			エラー
*/
func post[T any](reqUrl string, queryParams any, reqBody any, resBody *T) (err error) {
	// クエリパラメータをreqURLに追加
	if queryParams != struct{}{} {
		queryParamVal := reflect.ValueOf(queryParams)
		queryParamType := reflect.TypeOf(queryParams)
		params := url.Values{}
		for i := 0; i < queryParamVal.NumField(); i++ {
			params.Add(strings.ToLower(queryParamType.Field(i).Name), fmt.Sprintf("%v", queryParamVal.Field(i).Interface()))
		}
		reqUrl += "?" + params.Encode()
	}
	
	// リクエストボディをJSONに変換
	reqBodyJson := []byte{}
	if reqBody != struct{}{} {
		reqBodyJson, err = json.Marshal(reqBody)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// POSTリクエスト作成
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(reqBodyJson))
	if err != nil {
		log.Error(err)
		return err
	}

	// ヘッダー設定
	req.Header.Set("Content-Type", "application/json")

	// リクエスト送信
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	// レスポンスボディを読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return err
	}

	// レスポンスボディを構造体に変換
	if err := json.Unmarshal(body, resBody); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/* string 型の数値を int 型に変換する関数
	- value		変換する文字列
	> intValue	変換後の整数
*/
func convertStringToInt(value string) (intValue int) {
    if value == "" {
        return 0 // デフォルト値を0に設定
    }
    // 文字列を整数に変換
    intValue, err := strconv.Atoi(value)
    if err != nil {
        return 0 // 変換に失敗した場合もデフォルト値を0に設定
    }
    return intValue
}