// 主にスプレッドシートからの利用を想定したAPIを提供する
package api

import (
	"fmt"
	"reflect"
	"time"
)

/*
sql.Null~型などを持つ構造体を受け取り、csv形式の文字列に変換して返す関数
  - arg) sql.Null~型などを持つ構造体
  - return) vsc形式の文字列
*/
func convertToCsv(rows []interface{}) (csvString string) {

	for _, row := range rows {
		// row の要素を一つずつ処理
		reflectValue := reflect.ValueOf(row)
		for i := 0; i < reflectValue.NumField(); i++ {

			// フィールドの値を取得
			fieldValue := reflectValue.Field(i).Interface()
			// フィールドの型を取得
			fieldType := reflectValue.Field(i).Type().String()

			// フィールドの値を追加
			switch fieldType {
			case "string":
				csvString += fmt.Sprintf("%s", fieldValue)
			case "int64":
				csvString += fmt.Sprintf("%v", fieldValue)
			case "float64":
				csvString += fmt.Sprintf("%v", fieldValue)
			case "bool":
				csvString += fmt.Sprintf("%v", fieldValue)
			case "time.Time":
				csvString += fmt.Sprintf("%s", fieldValue)
			case "sql.NullInt64":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					csvString += fmt.Sprintf("%v", reflectValue.Field(i).FieldByName("Int64").Int())
				} else {
					csvString += "NULL"
				}
			case "sql.NullFloat64":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					csvString += fmt.Sprintf("%v", reflectValue.Field(i).FieldByName("Float64").Float())
				} else {
					csvString += "NULL"
				}
			case "sql.NullString":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					csvString += reflectValue.Field(i).FieldByName("String").String()
				} else {
					csvString += "NULL"
				}
			case "sql.NullTime":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					timeValue := reflectValue.Field(i).FieldByName("Time").Interface().(time.Time)
					csvString += timeValue.Format("2006-01-02")
				} else {
					csvString += "NULL"
				}
			}
			csvString += ","
		}

		// 最後のカンマを削除
		csvString = csvString[:len(csvString)-1]
		csvString += "\n"
	}

	return csvString
}

/*
sql.Null~型などを持つ構造体を受け取り、json形式の文字列に変換して返す関数
  - arg) sql.Null~型などを持つ構造体
  - return) json形式の文字列
*/
func convertToJson(rows []interface{}) (jsonString string) {

	jsonString = "["

	for _, row := range rows {
		// row の要素を一つずつ処理
		jsonString += "{"

		// row の要素を一つずつ処理
		reflectValue := reflect.ValueOf(row)
		for i := 0; i < reflectValue.NumField(); i++ {

			// フィールド名を取得
			fieldName := reflectValue.Type().Field(i).Name
			// フィールドの値を取得
			fieldValue := reflectValue.Field(i).Interface()
			// フィールドの型を取得
			fieldType := reflectValue.Field(i).Type().String()

			// フィールド名を追加
			jsonString += fmt.Sprintf("%q:", fieldName)

			// フィールドの値を追加
			switch fieldType {
			case "string":
				jsonString += fmt.Sprintf("%q", fieldValue)
			case "int64":
				jsonString += fmt.Sprintf("%v", fieldValue)
			case "float64":
				jsonString += fmt.Sprintf("%v", fieldValue)
			case "bool":
				jsonString += fmt.Sprintf("%v", fieldValue)
			case "time.Time":
				jsonString += fmt.Sprintf("%q", fieldValue)
			case "sql.NullInt64":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					jsonString += fmt.Sprintf("%v", reflectValue.Field(i).FieldByName("Int64").Int())
				} else {
					jsonString += "NULL"
				}
			case "sql.NullFloat64":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					jsonString += fmt.Sprintf("%v", reflectValue.Field(i).FieldByName("Float64").Float())
				} else {
					jsonString += "NULL"
				}
			case "sql.NullString":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					jsonString += fmt.Sprintf("%q", reflectValue.Field(i).FieldByName("String").String())
				} else {
					jsonString += "NULL"
				}
			case "sql.NullTime":
				if reflectValue.Field(i).FieldByName("Valid").Bool() {
					timeValue := reflectValue.Field(i).FieldByName("Time").Interface().(time.Time)
					jsonString += timeValue.Format("2006-01-02")
				} else {
					jsonString += "NULL"
				}
			}
			jsonString += ","
		}

		// 最後のカンマを削除
		jsonString = jsonString[:len(jsonString)-1]
		jsonString += "},\n"
	}

	// 最後のカンマを削除
	jsonString = jsonString[:len(jsonString)-2] + "]"

	return jsonString
}
