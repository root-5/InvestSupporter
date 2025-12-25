package api

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"reflect"
	"time"
)

/*
sql.Null~型などを持つ構造体を受け取り、json形式の文字列に変換して返す関数
  - arg) sql.Null~型などを持つ構造体
  - return) json形式の文字列
*/
func structToCSV(data interface{}) (string, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return "", fmt.Errorf("expected a slice, but got %s", v.Kind())
	}

	if v.Len() == 0 {
		return "", fmt.Errorf("slice is empty")
	}

	// スライスの最初の要素の型を取得
	elemType := v.Index(0).Type()

	// ヘッダー行を作成
	var headers []string
	for i := 0; i < elemType.NumField(); i++ {
		headers = append(headers, elemType.Field(i).Tag.Get("json"))
	}
	if err := writer.Write(headers); err != nil {
		return "", err
	}

	// データ行を作成
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		var record []string
		for j := 0; j < elem.NumField(); j++ {
			field := elem.Field(j)
			switch field.Kind() {
			case reflect.String:
				record = append(record, field.String())
			case reflect.Int64:
				record = append(record, fmt.Sprintf("%d", field.Int()))
			case reflect.Float64:
				record = append(record, fmt.Sprintf("%f", field.Float()))
			case reflect.Struct:
				switch field.Interface().(type) {
				case sql.NullInt64:
					if field.FieldByName("Valid").Bool() {
						record = append(record, fmt.Sprintf("%d", field.FieldByName("Int64").Int()))
					} else {
						record = append(record, "")
					}
				case sql.NullFloat64:
					if field.FieldByName("Valid").Bool() {
						record = append(record, fmt.Sprintf("%f", field.FieldByName("Float64").Float()))
					} else {
						record = append(record, "")
					}
				case sql.NullString:
					if field.FieldByName("Valid").Bool() {
						record = append(record, field.FieldByName("String").String())
					} else {
						record = append(record, "")
					}
				case sql.NullTime:
					if field.FieldByName("Valid").Bool() {
						timeValue := field.FieldByName("Time").Interface().(time.Time)
						record = append(record, timeValue.Format(time.RFC3339))
					} else {
						record = append(record, "")
					}
				default:
					return "", fmt.Errorf("unsupported struct field type: %s", field.Type())
				}
			default:
				return "", fmt.Errorf("unsupported field type: %s", field.Kind())
			}
		}
		if err := writer.Write(record); err != nil {
			return "", err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return buf.String(), nil
}
