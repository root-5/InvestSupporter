// 各コントローラーへの処理をまとめ、動作単位にまとめた関数を定義するパッケージ
package usecase

import (
	jquants "app/controller/jquants"
	postgres "app/controller/postgres"
	"fmt"
)

/* DB の初期化を行う関数 */
func InitDB() (err error) {
	fmt.Println("InitDB")
	err = postgres.InitDB()

	return err
}

/* リフレッシュトークンを取得した上でIDトークンを取得する関数
	> err	エラー
*/
func SetIdToken() (err error) {
	// リフレッシュトークンを取得
	refreshToken, err := jquants.GetRefreshToken()
	if err != nil {
		return err
	}

	// ID トークンを取得
	err = jquants.GetIdToken(refreshToken)
	if err != nil {
		return err
	}

	return nil
}

/* Jquants API から上場銘柄一覧を取得し、DB に保存する関数 */
func GetAndUpdateStocksInfo() (err error) {
	fmt.Println("GetAndUpdateStocksInfo")
	// 上場銘柄一覧を取得
	stocks, err := jquants.GetStocksInfo()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 取得した上場銘柄を DB に保存
	err = postgres.UpdateStocksInfo(stocks)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
