// 主にスプレッドシートからの利用を想定したAPIを提供する
package api

import (
	log "app/controller/log"
	postgres "app/controller/postgres"
	"encoding/json"
	"fmt"
	"net/http"
)

// ====================================================================================
// API関数
// ====================================================================================

var port = "8080"

// APIサーバーを起動する関数
func StartServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
	// リクエストのメソッドによって処理を分岐
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	default:
		fmt.Fprintf(w, "Method not allowed")
	}
}

// GETリクエストを処理する関数
func getHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストパスを取得
	path := r.URL.Path

	// リクエストパスによって処理を分岐
	switch path {

	// テスト用
	case "/":
		fmt.Fprintf(w, "Hello, world")

	// 上場銘柄一覧を取得
	case "/financial":
		// postges から財務情報を取得
		data, err := postgres.GetFinancialInfoForApi()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// レスポンスを返す
		sendResponse(w, data)

	default:
		fmt.Fprintf(w, "Not found")
	}
}

// POSTリクエストを処理する関数
func postHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストパスを取得
	path := r.URL.Path

	// リクエストパスによって処理を分岐
	switch path {
	case "/":
		fmt.Fprintf(w, "Hello, world")
	default:
		fmt.Fprintf(w, "Not found")
	}
}

// ====================================================================================
// リクエストボディの処理関数
// ====================================================================================
// リクエストボディを構造体に変換する関数
// func decodeRequestBody(r *http.Request, v interface{}) error {
// 	// リクエストボディを読み込む
// 	err := r.ParseForm()
// 	if err != nil {
// 		return err
// 	}

// 	// リクエストボディをJSONに変換
// 	body := r.Form.Get("body")
// 	err = json.Unmarshal([]byte(body), v)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// ====================================================================================
// レスポンスの処理関数
// ====================================================================================
// レスポンスを返す関数
func sendResponse(w http.ResponseWriter, v interface{}) {
	// レスポンスをJSONに変換
	res, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
