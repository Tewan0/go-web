package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "stock_db"
)

func main() {
	// สร้างตัวแปรเก็บ database connection
	// dsn := "root:@tcp(localhost:3306)/stock_db?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// สร้าง Log เพื่อเอาไว้ track ดู
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color
		},
	)

	// สร้าง database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	// ตรวจสอบว่า connect database ได้ไหม
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)
	fmt.Println("Database connection success... ^_^")

	// ทำ Migration ตัว Model กับ Database
	db.AutoMigrate(&Product{})
	fmt.Println("Migration success... ^_^")

	// ทดสอบเรียกใช้งานบันทึกข้อมูลลง table
	// newProduct := &Product{
	//  Name:     "iPhone 15",
	//  Price:    55000.75,
	//  Quantity: 9,
	// }
	// crateProductDB(db, newProduct)

	// ทดสอบเรียกใช้งานดึงข้อมูลทั้งหมดจาก table
	// getProductAllDB(db)

	// ทดสอบเรียกใช้งานดึงข้อมูลจาก table โดยใช้ id
	// getProductByIDDB(db, 4)

	// ทดสอบเรียกใช้งานดึงข้อมูลจาก table โดยใช้ name
	// getProductByNameDB(db, "no")

	// ทดสอบเรียกใช้งานแก้ไขข้อมูลใน table โดยใช้ id
	// updateProduct := &Product{
	//  Name:     "Readmi",
	//  Price:    55000.8,
	//  Quantity: 99,
	// }
	// updateProductByIDDB(db, 6, updateProduct)

	// ทดสอบเรียกใช้งานแก้ไขข้อมูลใน table โดยใช้ name
	// updateProduct := &Product{
	//  Price:    5000,
	//  Quantity: 15,
	// }
	// updateProductByNameDB(db, "Rea", updateProduct)

	// ทดสอบเรียกใช้งานลบข้อมูลใน table โดยใช้ id
	// deleteProductByIDDB(db, 5)

	// ทดสอบเรียกใช้งานลบข้อมูลใน table โดยใช้ name
	deleteProductByNameDB(db, "a")
}

// ฟังก์ชันบันทึกข้อมูลลง table (Create)
func crateProductDB(db *gorm.DB, product *Product) {
	// คำสั่งบันทึกข้อมูลลง table
	result := db.Create(product)

	// ตรวจสอบผลจากการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Create product success... ^_^")
}

// ฟังก์ชันเรียกดูข้อมูลทั้งหมดจาก table (Read)
func getProductAllDB(db *gorm.DB) {
	//สร้างตัวแปรเพื่อเก็บข้อมูลที่ได้จากการเรียกดู
	var products []Product

	// คำสั่งเรียกดูข้อมูลจาก table
	result := db.Find(&products)

	// ตรวจสอบผลจากการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get all product success... ^_^")
	fmt.Println(products)
	// return products
}

// ฟังก์ชันเรียกดูข้อมูลโดยมีเงื่อนไข ณ ที่นี้ใช้ id จาก table (Read)
func getProductByIDDB(db *gorm.DB, id int) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ได้จากการเรียกดู
	var product Product

	// คำสั่งเรียกดูข้อมูลจาก table
	result := db.First(&product, id)

	// ตรวจสอบผลจากการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product by id success... ^_^")
	fmt.Println(product)
	// return product
}

// ฟังก์ชันเรียกดูข้อมูลโดยมีเงื่อนไข ณ ที่นี้ใช้ name จาก table (Read)
func getProductByNameDB(db *gorm.DB, name string) {
	//สร้างตัวแปรเพื่อเก็บข้อมูลที่ได้จากการเรียกดู
	var products []Product

	// คำสั่งเรียกดูข้อมูลจาก table
	result := db.Where("name LIKE ? ", "%"+name+"%").Find(&products)

	// ตรวจสอบผลจากการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get all product success... ^_^")
	fmt.Println(products)
	// return products
}

// ฟังก์ชันอัปเดจข้อมูลโดยมีเงื่อนไข ณ ที่นี้ใช้ id จาก table (Update)
func updateProductByIDDB(db *gorm.DB, id int, product *Product) {
	// คำสั่งอัปเดตข้อมูล
	result := db.Model(&Product{}).Where("id = ?", id).Updates(product)

	// ตรวจสอบผลการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product success... ^_^")
}

// ฟังก์ชันอัปเดจข้อมูลโดยมีเงื่อนไข ณ ที่นี้ใช้ name จาก table (Update)
func updateProductByNameDB(db *gorm.DB, name string, product *Product) {
	// คำสั่งอัปเดตข้อมูล
	result := db.Model(&Product{}).Where("name Like ?", "%"+name+"%").Updates(product)

	// ตรวจสอบผลการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product success... ^_^")
}

// ฟังก์ชันลบข้อมูลโดยมีเงื่อนไข ณ ที่นี้ใช้ id จาก table (Delete)
func deleteProductByIDDB(db *gorm.DB, id int) {
	//่ คำสั่งลบข้อมูล (เป็นแบบ soft delete)
	// result := db.Delete(&Product{}, id)

	//่ คำสั่งลบข้อมูล (เป็นแบบ hard/premanent delete)
	result := db.Unscoped().Delete(&Product{}, id)

	// ตรวจสอบผลการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete product success... ^_^")
}

// ฟังก์ชันลบข้อมูลโดยมีเงื่อนไข ณ ที่นี้ใช้ name จาก table (Delete)
func deleteProductByNameDB(db *gorm.DB, name string) {
	//่ คำสั่งลบข้อมูล (เป็นแบบ soft delete)
	// result := db.Where("name Like ?", "%"+name+"%").Delete(&Product{})

	//่ คำสั่งลบข้อมูล (เป็นแบบ hard/premanent delete)
	result := db.Where("name Like ?", "%"+name+"%").Unscoped().Delete(&Product{})

	// ตรวจสอบผลการทำงาน
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete product success... ^_^")
}
