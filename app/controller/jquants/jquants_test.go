// JQuants API を利用するための関数をまとめたパッケージ
package jquants_test

import (
	jquants "app/controller/jquants"
	"fmt"
	"testing"
)

// Test は JQuants API 用のテスト関数
func TestJQuants(t *testing.T) {
	fmt.Println("Exec jquants.Test")
	var err error

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
