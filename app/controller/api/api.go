// 主にスプレッドシートからの利用を想定したAPIを提供する
package api

import (
	"app/controller/log"
	"app/controller/postgres"
	"app/usecase/usecase"
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
		data, err := postgres.GetFinancialInfoForApi("52530")
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// data を []interface{} に変換する
		var interfaceSlice []interface{} = make([]interface{}, len(data))
		for i, d := range data {
			interfaceSlice[i] = d
		}

		dataJson := convertToCsv(interfaceSlice)
		// dataJson := convertToJson(interfaceSlice)

		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(dataJson))

	// 上場銘柄一覧を取得
	case "/rebuild_data":
		// 全データを削除し、再取得
		err := usecase.RebuildData()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
// 構造体をjson形式の文字列に変換してレスポンスを返す関数
// func sendResponse(w http.ResponseWriter, v interface{}) {
// 	// レスポンスをJSONに変換
// 	res, err := json.Marshal(v)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// レスポンスを返す
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(res)
// }
