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

	// 全ての上場銘柄-財務情報を取得
	case "/financials":
		// postges から財務情報を取得
		data, err := postgres.GetFinancialsInfoForApi()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendCsvResponse(w, data)

	// コードを指定して上場銘柄-財務情報を取得
	case "/financial":
		// コードを取得
		code := r.URL.Query().Get("code")
		// コードが指定されていない場合はエラー
		if code == "" {
			http.Error(w, "code is required", http.StatusBadRequest)
			return
		}
		// コードが4桁の場合は5桁に変換
		if len(code) == 4 {
			code = code + "0"
		}
		// postges から財務情報を取得
		data, err := postgres.GetFinancialInfoForApi(code)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendCsvResponse(w, data)

	// 上場銘柄一覧を取得
	case "/admin/rebuild_data":
		// 全データを削除し、再取得
		err := usecase.RebuildData()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	// テスト用
	case "/":
		type explainStruct struct {
			Path  string `json:"このWEBサービスはアクセスするURLパスによって以下の機能を提供します"`
			Explain string `json:""`
		}
		data := []explainStruct{
			{Path: "GoogleスプレッドシートのIMPORTDATA関数の引数に以下のURLパスを指定してください", Explain: ""},
			{Path: "", Explain: ""},
			{Path: "URLパス", Explain: "使い方概要"},
			{Path: "/", Explain: "使い方概要"},
			{Path: "/financials", Explain: "財務情報（全上場銘柄）"},
			{Path: "/financial?code={{銘柄コード}}", Explain: "財務情報（単一） - {{銘柄コード}}は取得したい銘柄を4桁または5桁で指定"},
		}
		sendCsvResponse(w, data)

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
func sendCsvResponse(w http.ResponseWriter, data interface{}) {
	// CSV形式の文字列に変換
	csvString, err := structToCSV(data)
	if err != nil {
		log.Error(err)
		http.Error(w, "no data", http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	// w.Header().Set("Content-Type", "application/csv") // これをオンするとダウンロードされる
	w.Write([]byte(csvString))
}
