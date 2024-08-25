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
	"time"
)

// ====================================================================================
// 共通変数
// ====================================================================================

// リフレッシュトークン
var refreshToken string

// IDトークン
var IdToken string

// HTTPクライアント
var httpClient = &http.Client{}

// ====================================================================================
// 共通関数
// ====================================================================================

/*
GETリクエストを行い、レスポンスボディを取得する関数
  - return) url			リクエスト先URL文字列
  - return) queryParam	クエリパラメータ構造体
  - return) headers		ヘッダー構造体
  - return) resBody		レスポンスボディ構造体のポインタ
  - arg) err			エラー
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

	// ステータスコードが200以外の場合はエラー
	if resp.StatusCode != 200 {
		log.Error(fmt.Errorf("status Code: %d", resp.StatusCode))
		return fmt.Errorf("status Code: %d", resp.StatusCode)
	}

	// レスポンスボディを構造体に変換
	if err := json.Unmarshal(body, resBody); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
POSTリクエストを行い、レスポンスボディを取得する関数
  - return) url			リクエスト先URL文字列
  - return) queryParam	クエリパラメータ構造体
  - return) reqBody		リクエストボディ構造体
  - return) resBody		レスポンスボディ構造体のポインタ
  - arg) err			エラー
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

	// ステータスコードが200以外の場合はエラー
	if resp.StatusCode != 200 {
		log.Error(fmt.Errorf("status Code: %d", resp.StatusCode))
		return fmt.Errorf("status Code: %d", resp.StatusCode)
	}

	// レスポンスボディを構造体に変換
	if err := json.Unmarshal(body, resBody); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
string 型の数値を int 型に変換する関数
  - return) stringValue		変換する文字列
  - arg) intValue			変換後の整数
*/
func convertStringToInt(stringValue string) (intValue int) {
	if stringValue == "" {
		return 0 // デフォルト値を0に設定
	}
	// 文字列を整数に変換
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0 // 変換に失敗した場合もデフォルト値を0に設定
	}
	return intValue
}

/*
string 型の数値を float64 型に変換する関数
  - return) stringValue		変換する文字列
  - arg) floatValue			変換後の浮動小数点数
*/
func convertStringToFloat64(stringValue string) (floatValue float64) {
	if stringValue == "" {
		return 0 // デフォルト値を0に設定
	}
	// 文字列を浮動小数点数に変換
	floatValue, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0 // 変換に失敗した場合もデフォルト値を0に設定
	}
	return floatValue
}

/*
string 型の数値を time.Time 型に変換する関数
  - return) stringValue		変換する文字列
  - arg) timeValue			変換後の時刻
*/
func convertStringToTime(stringValue string) (timeValue time.Time) {
	if stringValue == "" {
		return time.Time{} // デフォルト値を0に設定
	}
	// 文字列を時刻に変換
	timeValue, err := time.Parse("2006-01-02", stringValue)
	if err != nil {
		return time.Time{} // 変換に失敗した場合もデフォルト値を0に設定
	}
	return timeValue
}
