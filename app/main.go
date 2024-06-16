package main

import (
	"app/controller"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program started")

	// 環境変数からメールアドレスとパスワードを取得
	email := os.Getenv("JQUANTS_EMAIL")
	pass := os.Getenv("JQUANTS_PASS")

	// メールアドレスとパスワードを渡してリフレッシュトークンを取得
	refreshToken, err := jquants.GetRefreshToken(email, pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("RefreshToken: ", refreshToken)

	// リフレッシュトークンを渡して ID トークンを取得
	idToken, err := jquants.GetIdToken(refreshToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("IdToken: ", idToken)
	
	_ = idToken

}