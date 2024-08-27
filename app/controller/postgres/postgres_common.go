// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

// ポインタが nil でない場合に値を返し、nil の場合は nil を返すヘルパー関数
func getValueOrNil[T any](ptr *T) interface{} {
	if ptr != nil {
		return *ptr
	}
	return nil
}
