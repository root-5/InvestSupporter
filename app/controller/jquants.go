// JQuants APIを利用するための関数をまとめたパッケージ
package jquants

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

// ====================================================================================
// 基本モジュール
// ====================================================================================
// HTTPクライアント
var httpClient = &http.Client{}

/* GETリクエストを行い、レスポンスボディを取得する関数
	- 入力) url - リクエスト先URL文字列
	- 入力) queryParam - クエリパラメータ構造体
	- 入力) headers - ヘッダー構造体
	- 入力) resBody - レスポンスボディ構造体のポインタ
	- 出力) err - エラー
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
		return fmt.Errorf("http.NewRequest Error: %v", err)
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
		return fmt.Errorf("httpClient.Do Error: %v", err)
	}
	defer resp.Body.Close()

	// ステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode Error: %v", resp.StatusCode)
	}

	// レスポンスボディを読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll Error: %v", err)
	}

	// レスポンスボディを構造体に変換
	if err := json.Unmarshal(body, resBody); err != nil {
		return fmt.Errorf("json.Unmarshal Error: %v", err)
	}

	return nil
}

/* POSTリクエストを行い、レスポンスボディを取得する関数
	- 入力) url - リクエスト先URL文字列
	- 入力) queryParam - クエリパラメータ構造体
	- 入力) reqBody - リクエストボディ構造体
	- 入力) resBody - レスポンスボディ構造体のポインタ
	- 出力) err - エラー
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
			return fmt.Errorf("json.Marshal Error: %v", err)
		}
	}

	// POSTリクエスト作成
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return fmt.Errorf("http.NewRequest Error: %v", err)
	}

	// ヘッダー設定
	req.Header.Set("Content-Type", "application/json")

	// リクエスト送信
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("httpClient.Do Error: %v", err)
	}
	defer resp.Body.Close()

	// ステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode Error: %v", resp.StatusCode)
	}

	// レスポンスボディを読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll Error: %v", err)
	}

	// レスポンスボディを構造体に変換
	if err := json.Unmarshal(body, resBody); err != nil {
		return fmt.Errorf("json.Unmarshal Error: %v", err)
	}

	return nil
}

// ====================================================================================
// レスポンス構造体
// ====================================================================================
// 上場銘柄一覧
type stockInfo struct {
	// Date              string `json:"Date"`
	Code              string `json:"Code"`
	CompanyName       string `json:"CompanyName"`
	CompanyNameEnglish string `json:"CompanyNameEnglish"`
	Sector17Code      string `json:"Sector17Code"`
	// Sector17CodeName  string `json:"Sector17CodeName"`
	Sector33Code      string `json:"Sector33Code"`
	// Sector33CodeName  string `json:"Sector33CodeName"`
	ScaleCategory     string `json:"ScaleCategory"`
	MarketCode        string `json:"MarketCode"`
	// MarketCodeName    string `json:"MarketCodeName"`
	MarginCode        string `json:"MarginCode"`
	// MarginCodeName    string `json:"MarginCodeName"`
}


// ====================================================================================
// API関数
// ====================================================================================
/* JQuants に登録したメールアドレスとパスワードを入力して、リフレッシュトークン（期限: 1週間）を取得する関数
	- 入力) email - JQuant に登録したメールアドレス
	- 入力) pass - JQuant に登録したパスワード
	- 出力) refreshToken - リフレッシュトークン
	- 出力) err - エラー
*/
func getRefreshToken() (refreshToken string, err error) {
	// return "eyJjdHkiOiJKV1QiLCJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiUlNBLU9BRVAifQ.lEvoH4XOE-grkUGtMLDwvuhiDtkIHEG4k2COyW9UtTqcxjeRKRVNmJqQZ2jb0WqNXcmUcEwAV6T6katBLZseg8Z2IkE0_-J2_BW1NRRYl_PbuXSH42SbuQgzTAtduLnMSliK17kGBSyvG_e0cd6GoivG0NRpmry9GUHgdYqvq_i0VYr2noWai1JOB3Is0_O2cJHPrSV3CdjH22X7u07qZWAQJekj1KioUAAQJb0Igu0I0-HiY2fO9cM-5tkobIJCdNy9zY-1iFKD3MtaXoCXqyXn6wpa-ao3ErfKBIwqV_b2OmuoiVbpT5Ab3fawcmq4bLcOMihOZJZ-oMVClmEriA.UrkYH_mZjHiBRE37.4uhGdOjUFDAahYCoD2XLf5JNuNZz7sLCMVMYOIGC4WJn5Eu88Pw7QP7tboeeqoKXv4FUoQ5Tve1mdvD3OwWfe8hlb0kXYYjqxaVSFlVPF2w0JJzHToiWRNBoRmzaNJe1ImJ3p8dVFrN7d7w5kyDxfKFTdnxvuYLZJMqmYdMUMn8M4ALKU-MTP9BxdA42qiCiomxAVLbX2zhM7VNC1Y_sXWnnSG9Lw9pUkWEoLV0ccJLyraRIa5oUSQ9WIuf-kuVb3Bu6KfwPFnl4lcEyaDoXObXd9b2xzTxmNm-CjAwHxFX5S577xaXR2oDEoQjpZjFFvjIM1ISBLFCVvx7zT6aB1Olpt-OnqaR2kumOy5jb7126R79TgdIsdyMbpNIDHxBtTXFhsFaI2bp5i9Uko3hC_TZdofyxmreYChHYPcCO1PO-j_857JnH5W_qExm3_HqQl5Io9ZmUaNW5UDfbzeCN2hp3-FAQ3fU41-kp1Vt1manP0-JAUagtDUtKc76j2S94JbZiSZiWEVKVI_cUZlfeePvok4R5g4qJslXSpql5vA27Hz_LoD8LOf1J_CvnylLDWStrOgth0QBz3mWJeH2G9n1Qww-SOd8bsKPtBXtqGroO5LfE9nTNPr1QFGDryNEjR653d1Sq-76N4oaRIZ9pv-JlclY0_0hNH6NkVgpZcWDEt4BL8wkEaZNJJSBuCHC6tX6W87_Ece6rkq9XYT-HQajNEF_eSePlAhw6RmjIGpUDdlf4UtOQNCnrqNfokca21zS6uM-2SGq0I-nQFtO-rkmjVEgXE8oXszNHG3P63yvPbof1YGK7Qz0cI27FSNUnIR42NpX0ft6U-KZoO96uJmf49zhS_dRUgv_8y7jaa-9zmwhCg-mC9ZLnn6FFWS4ONNKGCCdTNDhy1VEpDI7okX37Q56j61pebe6rhuLFDHeEUpEOBOwH8JhmwP7niTa6xbFm7HjooTnRlhY0SltMoulFAydgzCQNCIwKtsOMGqqqu7EJTMIxmeXQb8z4AGO-VD9vM9o3OUzN-ohnO2VTa3n_pAJOQf6x4kiDPfyyvTCUhRY3GL4ML-nwyntNVrC1wlHB1MiQoQjdrKCHA6Ppy2B07wYA4e9jqaPNZYXiiBh9q6XB1jojc9OPuHZJvKV64TQPIP9RRDt8GMYDv0jC2Wz7jKTbp_boim9DQTIgRaFsa3miSjYL6ikQOMW_Viq9Nm0WI32t2mTOQQ2gREhdWNLuks6YaXoaNI17_NYjp6Bl4NpaDD_-GOdeJiXtKgF64ZCAJJadGolkN5aQPsfT4oMsRLUwCX4S6f9b-dR0jH4ppTqz95XxT94uU2hN8E48GhelP4YIZlRqrw.WyO9VqhsLA4Og65fcm2VhA", nil
	// return "dummy_refreshtoken", nil
	// 環境変数からメールアドレスとパスワードを取得
	email := os.Getenv("JQUANTS_EMAIL")
	pass := os.Getenv("JQUANTS_PASS")

	// リクエスト先URL
	url := "https://api.jquants.com/v1/token/auth_user"

	// クエリパラメータ
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

	// レスポンスボディ
	type resBodyType struct {
		RefreshToken string `json:"refreshToken"`
	}
	var resBody resBodyType

	// POSTリクエスト
	err = post(url, queryParams, reqBody, &resBody)
	if err != nil {
		return "", fmt.Errorf("post Error: %v", err)
	}

	// トークンを取得
	refreshToken = resBody.RefreshToken

	return refreshToken, nil
}

/* リフレッシュトークンを渡して、ID トークン（期限: 24時間）を取得する関数
	- 入力) refreshToken - getRefreshToken 関数で取得したトークン
	- 出力) idToken - ID トークン
	- 出力) err - エラー
*/
func getIdToken(refreshToken string) (idToken string, err error) {
	// return "eyJraWQiOiJHQXNvU2xxUzMyUktLT2lVYm1xcjU3ekdYNE1TVFhsWFBrbDNJTmhWKzNzPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI3ZWVjNGQ3Ni03NzhmLTQ1NzAtOGY2Ni0zNjRkZjc0MWZkNmIiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmFwLW5vcnRoZWFzdC0xLmFtYXpvbmF3cy5jb21cL2FwLW5vcnRoZWFzdC0xX0FGNjByeXJ2NCIsImNvZ25pdG86dXNlcm5hbWUiOiI3ZWVjNGQ3Ni03NzhmLTQ1NzAtOGY2Ni0zNjRkZjc0MWZkNmIiLCJvcmlnaW5fanRpIjoiNWRlZTcxODgtYzhhNy00Njk3LWJiZGMtN2VmYmFkNTAwYTVmIiwiYXVkIjoiNXZyN2xiOGppdThvZmhvZmJmYWxmbm1waWkiLCJldmVudF9pZCI6Ijk4Njk1MWM2LTQxNGQtNDIzNy04Nzk4LWZmY2EwZjE0ODJiYiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzE4NTQwNTM4LCJleHAiOjE3MTg2MjY5NzMsImlhdCI6MTcxODU0MDU3MywianRpIjoiODI2ZTQ2ZGItYTRmZC00Y2M0LTkzNjUtMjdkZTNmNDhiMTA1IiwiZW1haWwiOiJ3YXRlci52aWxsYWdlLnJpZ2h0Lm5lYXJAZ21haWwuY29tIn0.Wv00zXyDqhGetvei6Uet38bxIM87RgyTZTUeh7chgTHuya1FgfWeF-QUdIO6CjiKz87TZdTPC7hhzD6VcJMzRMUv3J7cf6RYcBSiEd1siLJEhtd06qzdROwLOURJSclFKO9HQo0RhU_YmUUj-sWyFJVbtrvYesqQN2jR0dDcQYJs3ITiS3QDl9sQfZvKAXVPJOTU_rm-mCYBpdL1iJAxRNCO2dwqqMTx9rDyplf383ndV4WaUwB9dJefj-LRJdn_r052muQQ68PJfispvOgkqr9GGHtilgcZnXBSN2vn2L0EowqE7cRWHf4DPXM8PAm3WL2E8BiuvNKD_3XZNoFOEA", nil
	// return "dummy_idtoken", nil

	// リクエスト先URL
	url := "https://api.jquants.com/v1/token/auth_refresh"

	// クエリパラメータ
	type queryParamsType struct {
		RefreshToken string
	}
	queryParam := queryParamsType {
		RefreshToken: refreshToken,
	}

	// リクエストボディ
	type reqBodyType struct {}
	reqBody := reqBodyType{}

	// レスポンスボディ
	type resBodyStruct struct {
		IdToken string `json:"idToken"`
	}
	var resBody resBodyStruct

	// POSTリクエスト
	err = post(url, queryParam, reqBody, &resBody)
	if err != nil {
		return "", fmt.Errorf("post Error: %v", err)
	}

	// トークンを取得
	idToken = resBody.IdToken

	return idToken, nil
}

/* メールアドレスとパスワードを入力して、ID トークン（期限: 24時間）を取得する関数
	- 入力) email - JQuant に登録したメールアドレス
	- 入力) pass - JQuant に登録したパスワード
	- 出力) idToken - ID トークン
	- 出力) err - エラー
*/
func SetIdToken() (idToken string, err error) {
	// return "idToken_for_test ", nil

	// リフレッシュトークンを取得
	refreshToken, err := getRefreshToken()
	if err != nil {
		return "", fmt.Errorf("getRefreshToken Error: %v", err)
	}

	// ID トークンを取得
	idToken, err = getIdToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("getIdToken Error: %v", err)
	}

	return idToken, nil
}

/* 上場銘柄一覧を取得する関数
	- 入力) idToken - SetIdToken 関数で取得したトークン
	- 出力) stockList - 上場銘柄情報の配列
*/
func GetStockList(idToken string) (stockList []stockInfo, err error) {
	// リクエスト先URL
	url := "https://api.jquants.com/v1/listed/info"

	// クエリパラメータ
	type queryParamsType struct {}
	queryParams := queryParamsType{}

	// ヘッダー
	type headersType struct {
		Authorization string `json:"Authorization"`
	}
	headers := headersType {
		Authorization: idToken,
	}
	fmt.Println(headers)

	// レスポンスボディ
	type resBodyStruct struct {
		Info []stockInfo `json:"info"`
	}
	var resBody resBodyStruct

	// GETリクエスト
	err = get(url, queryParams, headers, &resBody)
	if err != nil {
		return nil, fmt.Errorf("post Error: %v", err)
	}

	// 上場銘柄一覧を取得
	stockList = resBody.Info

	return stockList, nil
}
