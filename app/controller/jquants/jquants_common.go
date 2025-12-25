package jquants

import (
	"app/controller/log"
	"bytes"
	"database/sql"
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

// HTTPクライアント
var httpClient = &http.Client{}

// ====================================================================================
// 共通関数
// ====================================================================================

/*
GETリクエストを行い、レスポンスボディを取得する関数
  - arg) url		リクエスト先URL文字列
  - arg) queryParam	クエリパラメータ構造体
  - arg) headers	ヘッダー構造体
  - arg) resBody	レスポンスボディ構造体のポインタ
  - return) err		エラー
*/
func get[T any](reqUrl string, queryParams any, headers any, resBody *T) (err error) {
	// クエリパラメータをreqURLに追加
	if queryParams != struct{}{} {
		queryParamVal := reflect.ValueOf(queryParams)
		queryParamType := reflect.TypeOf(queryParams)
		params := url.Values{}
		for i := 0; i < queryParamVal.NumField(); i++ {
			tag := queryParamType.Field(i).Tag.Get("json")
			key := ""
			omitempty := false
			if tag != "" {
				parts := strings.Split(tag, ",")
				key = parts[0]
				if len(parts) > 1 && parts[1] == "omitempty" {
					omitempty = true
				}
			} else {
				key = strings.ToLower(queryParamType.Field(i).Name)
			}

			val := fmt.Sprintf("%v", queryParamVal.Field(i).Interface())
			if omitempty && val == "" {
				continue
			}
			params.Add(key, val)
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
			headerName := headerType.Field(i).Tag.Get("json")
			if headerName == "" {
				headerName = strings.ToLower(headerType.Field(i).Name)
			}
			req.Header.Set(headerName, fmt.Sprintf("%v", headerVal.Field(i).Interface()))
		}
	}

	// リクエスト送信
	var resp *http.Response
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		resp, err = httpClient.Do(req)

		// ネットワークエラー時の再試行
		if err != nil {
			log.Error(err)
			fmt.Println("レスポンスがありません、30秒待機後に再リクエストします")
			time.Sleep(30 * time.Second)
			continue
		}

		// レートリミットに抵触した場合は5秒待機してリトライ
		if resp.StatusCode == 429 {
			resp.Body.Close()
			fmt.Printf("レートリミットに抵触しました (リトライ: %d/%d)\n", i+1, maxRetries)
			time.Sleep(5 * time.Second)
			continue
		}

		// 正常にレスポンスが返ってきた場合はループを抜ける
		break
	}
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("リトライ回数を超過しました")
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
  - arg) url		リクエスト先URL文字列
  - arg) queryParam	クエリパラメータ構造体
  - arg) reqBody	リクエストボディ構造体
  - arg) resBody	レスポンスボディ構造体のポインタ
  - return) err		エラー
*/
func post[T any](reqUrl string, queryParams any, reqBody any, resBody *T) (err error) {
	// クエリパラメータをreqURLに追加
	if queryParams != struct{}{} {
		queryParamVal := reflect.ValueOf(queryParams)
		queryParamType := reflect.TypeOf(queryParams)
		params := url.Values{}
		for i := 0; i < queryParamVal.NumField(); i++ {
			tag := queryParamType.Field(i).Tag.Get("json")
			key := ""
			omitempty := false
			if tag != "" {
				parts := strings.Split(tag, ",")
				key = parts[0]
				if len(parts) > 1 && parts[1] == "omitempty" {
					omitempty = true
				}
			} else {
				key = strings.ToLower(queryParamType.Field(i).Name)
			}

			val := fmt.Sprintf("%v", queryParamVal.Field(i).Interface())
			if omitempty && val == "" {
				continue
			}
			params.Add(key, val)
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
		fmt.Println("30秒待機後に再リクエストします")

		// 30秒待機したのち再リクエスト
		time.Sleep(30 * time.Second)
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Error(err)
			return err
		}
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
string 型の数値を sql.NullString 型に変換する関数
  - arg) stringValue	変換する文字列
  - return) stringValue	変換後の文字列
*/
func convertStringToString(stringValue string) (stringOrNilValue sql.NullString) {
	if stringValue == "" {
		stringOrNilValue = sql.NullString{String: "", Valid: false}
	} else {
		stringOrNilValue = sql.NullString{String: stringValue, Valid: true}
	}
	return stringOrNilValue
}

/*
string 型の数値を sql.NullInt64 型に変換する関数
  - arg) stringValue	変換する文字列
  - return) intValue	変換後の整数
*/
func convertStringToInt64(stringValue string) (intValue sql.NullInt64) {
	if stringValue == "" {
		intValue = sql.NullInt64{Int64: 0, Valid: false}
	} else {
		// 文字列を整数に変換
		intOnlyValue, _ := strconv.ParseInt(stringValue, 10, 64)
		intValue = sql.NullInt64{Int64: intOnlyValue, Valid: true}
	}
	return intValue
}

/*
string 型の数値を sql.NullFloat64 型に変換する関数
  - arg) stringValue	変換する文字列
  - return) floatValue	変換後の浮動小数点数
*/
func convertStringToFloat64(stringValue string) (floatValue sql.NullFloat64) {
	if stringValue == "" {
		floatValue = sql.NullFloat64{Float64: 0, Valid: false}
	} else {
		// 文字列を浮動小数点数に変換
		floatOnlyValue, _ := strconv.ParseFloat(stringValue, 64)
		floatValue = sql.NullFloat64{Float64: floatOnlyValue, Valid: true}
	}
	return floatValue
}

/*
string 型の数値を sql.NullTime 型に変換する関数
  - arg) stringValue	変換する文字列
  - return) timeValue	変換後の時刻
*/
func convertStringToTime(stringValue string) (timeValue sql.NullTime) {
	if stringValue == "" {
		timeValue = sql.NullTime{Time: time.Time{}, Valid: false}
	} else {
		// 文字列を時刻に変換
		timeOnlyValue, _ := time.Parse("2006-01-02", stringValue)
		timeValue = sql.NullTime{Time: timeOnlyValue, Valid: true}
	}
	return timeValue
}

/*
any 型の数値を sql.NullFloat64 型に変換する関数
  - arg) anyValue		変換する値（"Null"という string か float64）
  - return) floatValue	変換後の浮動小数点数
*/
func convertAnyToFloat64(anyValue any) (floatValue sql.NullFloat64) {
	if anyValue == "Null" || anyValue == "" || anyValue == nil {
		floatValue = sql.NullFloat64{Float64: 0, Valid: false}
	} else {
		floatValue = sql.NullFloat64{Float64: anyValue.(float64), Valid: true}
	}
	return floatValue
}
