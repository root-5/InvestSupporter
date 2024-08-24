// JQuants API を利用するための関数をまとめたパッケージ
package jquants_test

import (
	jquants "app/controller/jquants"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Test は JQuants API 用のテスト関数
func TestJQuants(t *testing.T) {
	fmt.Println("==============================")
	fmt.Println("Exec jquants.Test")
	fmt.Println("==============================")

	var err error

	// .env ファイルを読み込む
	envText, err := os.Open("../../../infra/app/.env")
	if err != nil {
		t.Errorf("Open .env failed: %v", err)
		return
	}
	defer envText.Close()

	// envText を1行ずつ読んで環境変数にセット
	for {
		var envLine string
		_, err := fmt.Fscanln(envText, &envLine)
		if err != nil {
			break
		}
		// envLine を "=" で分割、「"」が含まれている場合は削除してから環境変数にセット
		envLineSlice := strings.Split(envLine, "=")
		envLineSlice[1] = strings.Replace(envLineSlice[1], "\"", "", -1)
		os.Setenv(envLineSlice[0], envLineSlice[1])
	}
	// 環境変数がセットされているか確認
	if os.Getenv("JQUANTS_EMAIL") == "" || os.Getenv("JQUANTS_PASS") == "" {
		t.Errorf("Set env failed: %v", err)
		return
	}

	/*
		Init
	*/
	fmt.Println("Test Init")
	jquants.Init()

	// IdTokenForTest が10文字以下ならNG
	if len(jquants.IdTokenForTest) < 10 {
		t.Errorf("Init failed: %v", err)
		return
	} else {
		fmt.Println(">> OK")
	}

	/*
		GetStocksInfo
	*/
	fmt.Println("Test GetStocksInfo")
	stocksInfo, err := jquants.GetStocksInfo()

	// GetStocksInfo がエラーならNG
	if err != nil {
		t.Errorf("GetStocksInfo failed: %v", err)
		return
	} else {
		fmt.Println(">> len(stocksInfo) = ", len(stocksInfo))
		fmt.Println(">> OK")
	}

	/*
		GetFinancialInfo
	*/
	fmt.Println("Test GetFinancialInfo")
	financialInfo, err := jquants.GetFinancialInfo("7203")

	// GetFinancialInfo がエラーならNG
	if err != nil {
		t.Errorf("GetFinancialInfo failed: %v", err)
		return
	} else {
		fmt.Println(">> len(financialInfo) = ", len(financialInfo))
		fmt.Println(">> OK")
	}
}
