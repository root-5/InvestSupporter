package main

import (
	jquants "app/controller"

	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
  }

func main() {
	fmt.Println("Program started")

	// 環境変数からメールアドレスとパスワードを取得
	email := os.Getenv("JQUANTS_EMAIL")
	pass := os.Getenv("JQUANTS_PASS")



	// DB の初期化
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
	fmt.Println(dsn)

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&Product{})
  
	// Create
	db.Create(&Product{Code: "D42", Price: 100})
  
	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42
  
	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  
	// Delete - delete product
	// db.Delete(&product, 1)

	// Read
	var productAlfa Product
	db.First(&productAlfa, 1) // find product with integer primary key
	fmt.Printf("ProductCode: %v\n", productAlfa.Code)
	fmt.Printf("ProductPrice: %v\n", productAlfa.Price)

	// ID トークンをセット
	idToken, err := jquants.SetIdToken(email, pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = idToken
}