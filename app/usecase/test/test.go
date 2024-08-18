// テスト用のパッケージ
package test

import (
	jquants "app/controller/jquants"
	log "app/controller/log"
	postgres "app/controller/postgres"
	usecase "app/usecase/usecase"
	"fmt"
)

var isUnitTest = true
var isUsecaseTest = false

// Test はテスト用の関数
// 現状はユースケース単位でのテストを実行している
func Test() {
	fmt.Println("Exec Test")

	// Init
	fmt.Println("Test Init")
	usecase.Init()
	fmt.Println(">> OK")

	// UnitTest
	if isUnitTest {
		TestUnit()
	}

	// UsecaseTest
	if isUsecaseTest {
		TestUsecase()
	}
}

// TestUnit はユニットテスト用の関数
func TestUnit() {
	fmt.Println("Exec TestUnit")

	// GetStocksInfo
	fmt.Println("Test GetStocksInfo")
	stocks, err := jquants.GetStocksInfo()
	if err != nil {
		log.Error(err)
		fmt.Println(">> NG")
	} else {
		fmt.Println(">> OK")
		fmt.Println(stocks)
	}

	// GetStocksInfo
	fmt.Println("Test GetStocksInfo")
	stocks, err = postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		fmt.Println(">> NG")
	} else {
		fmt.Println(">> OK")
		fmt.Println(stocks)
	}
}

// TestUsecase はユースケーステスト用の関数
func TestUsecase() {
	fmt.Println("Exec TestUseCase")

	// GetAndSaveStocksInfo
	fmt.Println("Test GetAndSaveStocksInfo")
	err := usecase.GetAndSaveStocksInfo()
	if err != nil {
		log.Error(err)
		fmt.Println(">> NG")
	} else {
		fmt.Println(">> OK")
	}
}