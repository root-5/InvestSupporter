package jquants_test

import (
	"app/controller/jquants"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

// Test は JQuants API 用のテスト関数
func TestJQuants(t *testing.T) {
	fmt.Println("")
	fmt.Println("============================================================")
	fmt.Println("TEST JQUANTS")
	fmt.Println("============================================================")
	fmt.Println("")

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
		fmt.Println("")
		return
	}

	/*
		Init
	*/
	fmt.Println("Test SchedulerStart")
	jquants.SchedulerStart()
	if os.Getenv("JQUANTS_API_KEY") == "" {
		t.Errorf("JQUANTS_API_KEY is not set")
		return
	} else {
		fmt.Println(">> OK")
		fmt.Println("")
	}

	/*
		GetStocksInfo
	*/
	fmt.Println("Test GetStocksInfo")
	stocksInfo, err := jquants.FetchStocksInfo()
	if err != nil {
		t.Errorf("GetStocksInfo failed: %v", err)
		return
	} else {
		fmt.Println(">> len(stocksInfo) = ", len(stocksInfo))
		fmt.Println(">> OK")
		fmt.Println("")
	}

	/*
		GetStatementInfo
	*/
	fmt.Println("Test GetStatementInfo (7203)")
	StatementInfo, err := jquants.FetchStatementsInfo("7203")
	if err != nil {
		t.Errorf("GetStatementInfo failed: %v", err)
		return
	} else {
		// StatementInfo[0] の構造体に含まれる sql.Null~ 型の変数のうち、Valid が false のものを表示
		pattern, _ := regexp.Compile("sql.Null.*")
		m := reflect.ValueOf(StatementInfo[0])
		for i := 0; i < m.NumField(); i++ {
			if pattern.MatchString(reflect.TypeOf(m.Field(i).Interface()).String()) {
				if !reflect.ValueOf(m.Field(i).Interface()).FieldByName("Valid").Bool() {
					fmt.Println(">> ", m.Type().Field(i).Name, " = ", m.Field(i).Interface())
				}
			}
		}
		fmt.Println(">> OK")
		fmt.Println("")
	}

	fmt.Println("Test GetStatementInfo (2024-8-30)")
	StatementInfo, err = jquants.FetchStatementsInfo("2024-08-30")
	if err != nil {
		t.Errorf("GetStatementInfo failed: %v", err)
		return
	} else {
		// StatementInfo[0] の構造体に含まれる sql.Null~ 型の変数のうち、Valid が false のものを表示
		pattern, _ := regexp.Compile("sql.Null.*")
		m := reflect.ValueOf(StatementInfo[0])
		for i := 0; i < m.NumField(); i++ {
			if pattern.MatchString(reflect.TypeOf(m.Field(i).Interface()).String()) {
				if !reflect.ValueOf(m.Field(i).Interface()).FieldByName("Valid").Bool() {
					fmt.Println(">> ", m.Type().Field(i).Name, " = ", m.Field(i).Interface())
				}
			}
		}
		fmt.Println(">> len(StatementInfo) = ", len(StatementInfo))
		fmt.Println(">> OK")
		fmt.Println("")
	}
}
