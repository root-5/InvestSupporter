// JQuants APIを利用するための関数をまとめたパッケージ
package jquants

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
OSTリクエストを行い、レスポンスボディを取得する関数

	- 入力) url - リクエスト先URL
	- 入力) reqBody - リクエストボディ
	- 出力) body - レスポンスボディ
	- 出力) err - エラー
*/
// post: POSTリクエストを行い、レスポンスボディを取得
func post(url string, reqBody interface{}) (body []byte, err error) {
	// リクエストボディをJSONに変換
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal Error: %v", err)
	}

	// POSTリクエスト
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return nil, fmt.Errorf("http.Post Error: %v", err)
	}
	defer resp.Body.Close()

	// ステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode Error: %v", resp.StatusCode)
	}

	// レスポンスボディを読み込み
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll Error: %v", err)
	}

	return body, nil
}

/*
JQuants に登録したメールアドレスとパスワードを入力して、リフレッシュトークンを取得する関数

	- 入力) email - JQuant に登録したメールアドレス
	- 入力) pass - JQuant に登録したパスワード
	- 出力) refreshToken - トークン
	- 出力) err - エラー
*/
func GetRefreshToken(email string, pass string) (refreshToken string, err error) {
	fmt.Println(">> GetRefreshToken 開始")

	// return "dummy_token", nil

	// リクエストボディ/レスポンスボディの構造体
	type reqBodyStruct struct {
		Mailaddress string `json:"mailaddress"`
		Password    string `json:"password"`
	}
	type resBodyStruct struct {
		RefreshToken string `json:"refreshToken"`
	}

	// リクエスト先URL
	url := "https://api.jquants.com/v1/token/auth_user"

	// リクエストボディ
	reqBody := reqBodyStruct{
		Mailaddress: email,
		Password:    pass,
	}

	// POSTリクエスト
	body, err := post(url, reqBody)
	if err != nil {
		return "", fmt.Errorf("post Error: %v", err)
	}

	// レスポンスボディを構造体に変換
	var resBody resBodyStruct
	if err := json.Unmarshal(body, &resBody); err != nil {
		return "", fmt.Errorf("json.Unmarshal Error: %v", err)
	}

	// トークンを取得
	refreshToken = resBody.RefreshToken

	return refreshToken, nil
}