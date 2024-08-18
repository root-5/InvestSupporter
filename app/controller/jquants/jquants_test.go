// JQuants API を利用するための関数をまとめたパッケージ
package jquants

import (
	log "app/controller/log"
	"fmt"
)

// Test は JQuants API 用のテスト関数
func Test() {
	fmt.Println("Exec jquants.Test")

	var err error

	// getRefreshToken
	fmt.Println("Test getRefreshToken")
	err = getRefreshToken()
	if err != nil {
		fmt.Println(">> NG")
		log.Error(err)
	} else {
		fmt.Println(">> OK")
	}

	// getIdToken
	fmt.Println("Test getIdToken")
	err = getIdToken(refreshToken)
	if err != nil {
		fmt.Println(">> NG")
		log.Error(err)
	} else {
		fmt.Println(">> OK")
	}

	// setIdToken
	fmt.Println("Test setIdToken")
	err = setIdToken()
	if err != nil {
		fmt.Println(">> NG")
		log.Error(err)
	} else {
		fmt.Println(">> OK")
	}

	// GetFinancialInfo
	fmt.Println("Test GetFinancialInfo")
	financialInfo, err := GetFinancialInfo("2023-01-30")
	if err != nil {
		fmt.Println(">> NG")
		log.Error(err)
	} else {
		fmt.Println(">> OK")
		fmt.Println(financialInfo)
	}
}