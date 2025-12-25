// API を提供するパッケージ
package api

import (
	"app/controller/log"
	"app/controller/postgres"
	"app/usecase/usecase"
	"fmt"
	"net/http"
	"strings"
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

	// 基本情報を取得
	case "/financial":
		// postges から基本情報を取得
		data, err := postgres.GetBasicInfoForApi()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendCsvResponse(w, data)

		// すべての財務情報を取得
	case "/statement":
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
		data, err := postgres.GetStatementsInfo(code)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendCsvResponse(w, data)

	// 株価終値情報を取得
	case "/closeprice":
		// コードと日付を取得
		code := r.URL.Query().Get("code")
		ymd := r.URL.Query().Get("ymd")
		// コードが指定されていない場合はエラー
		if code == "" {
			http.Error(w, "codes is required", http.StatusBadRequest)
			return
		}
		// 日付フォーマットをチェック
		if ymd != "" {
			isDateRange := strings.Contains(ymd, "~")
			if isDateRange {
				if len(ymd) != 21 { // "YYYY-MM-DD~YYYY-MM-DD"
					http.Error(w, "ymd range format is invalid. use YYYY-MM-DD~YYYY-MM-DD", http.StatusBadRequest)
					return
				}
			} else {
				if len(ymd) != 10 { // "YYYY-MM-DD"
					http.Error(w, "ymd format is invalid. use YYYY-MM-DD", http.StatusBadRequest)
					return
				}
			}
		}
		// カンマ区切りのコードをスライスに変換
		codes := strings.Split(code, ",")
		// コードが4桁の場合は5桁に変換
		for i := range codes {
			if len(codes[i]) == 4 {
				codes[i] = codes[i] + "0"
			}
		}
		// DB から株価情報を取得
		data, err := usecase.GetClosePricesInfo(codes, ymd)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println(len(data))
		// fmt.Println(data)
		sendCsvResponse(w, data)

	// 株価情報を取得
	case "/price":
		// コードと日付を取得
		code := r.URL.Query().Get("code")
		ymd := r.URL.Query().Get("ymd")
		// 日付フォーマットをチェック
		if ymd != "" {
			isDateRange := strings.Contains(ymd, "~")
			if isDateRange {
				if len(ymd) != 21 { // "YYYY-MM-DD~YYYY-MM-DD"
					http.Error(w, "ymd range format is invalid. use YYYY-MM-DD~YYYY-MM-DD", http.StatusBadRequest)
					return
				}
			} else {
				if len(ymd) != 10 { // "YYYY-MM-DD"
					http.Error(w, "ymd format is invalid. use YYYY-MM-DD", http.StatusBadRequest)
					return
				}
			}
		} else if code == "" {
			http.Error(w, "code or ymd is required", http.StatusBadRequest)
			return
		}
		// コードが4桁の場合は5桁に変換
		if len(code) == 4 {
			code = code + "0"
		}
		// コードが指定されているときはスライスに変換、そうでないときは空のスライス
		var codes []string
		if code != "" {
			codes = []string{code}
		} else {
			codes = []string{}
		}
		// DB から株価情報を取得
		data, err := postgres.GetPricesInfo(codes, ymd)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendCsvResponse(w, data)

	// 株価情報データを削除し、再取得
	case "/admin/reset/price":
		err := usecase.ResetPriceInfoAll()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	// 財務情報データを削除し、再取得
	case "/admin/reset/statement":
		err := usecase.ResetStatementInfoAll()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	// 全データを削除し、再取得
	case "/admin/reset/all":
		err := usecase.ResetDataAll()
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	// 説明用 html
	case "/howto":
		http.ServeFile(w, r, "controller/api/index.html")

	// ファビコン
	case "/favicon.ico":
		http.ServeFile(w, r, "controller/api/smile.ico")

	// 説明用
	case "/":
		resText := "<a href=\"./howto\">./howto</a>"
		w.Write([]byte(resText))

	default:
		// アクセス元のIPアドレスを取得
		// 将来的には複数回アクセスがあった場合に、そのIPアドレスをブロックするようにする
		ip := r.RemoteAddr
		log.Info(path)
		log.Info("Not found: " + ip)
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
// 構造体やスライスをjson形式の文字列に変換してレスポンスを返す関数
func sendCsvResponse(w http.ResponseWriter, data interface{}) {
	dataSlice, ok := data.([][]string)
	if ok {
		csvString := ""
		for i := range dataSlice {
			for j := range dataSlice[i] {
				csvString += dataSlice[i][j]
				if j < len(dataSlice[i])-1 {
					csvString += ","
				}
			}
			csvString += "\n"
		}
		// w.Header().Set("Content-Type", "text/csv") // これを有効にすると、CSV がダウンロードされる
		w.Write([]byte(csvString))
	} else {
		// CSV形式の文字列に変換
		csvString, err := structToCSV(data)
		if err != nil {
			log.Error(err)
			http.Error(w, "no data", http.StatusInternalServerError)
			return
		}
		// w.Header().Set("Content-Type", "application/json") // これを有効にすると、CSV がダウンロードされる
		w.Write([]byte(csvString))
	}
}
