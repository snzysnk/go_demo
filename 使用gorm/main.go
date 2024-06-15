package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type class struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getConnect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("xieruixiang:Iam@123456@tcp(139.196.105.2:3306)/xim"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", db) //&{Config:0xc0001a6630 Error:<nil> RowsAffected:0 Statement:0xc0001d0000 clone:1}
	return db
}

func set(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	s.SetMaxOpenConns(10)
	s.SetMaxIdleConns(5)
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&class{})
	if err != nil {
		log.Fatal(err)
	}
}

func create(db *gorm.DB) {
	data := class{
		Name: "Anna",
		Age:  17,
	}
	tx := db.Create(&data)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}
	fmt.Println(data.Id) //1
}

func update(db *gorm.DB) {
	data := class{
		Id:   1,
		Name: "Anna",
	}
	tx := db.Model(&data).Update("Name", "Anna_1")
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}
	if affected := tx.RowsAffected; affected <= 0 {
		log.Fatal("update fail")
	}
	fmt.Println(data.Name, data.Age) // Anna_1 100
}

func update2(db *gorm.DB) {
	data := class{}
	// UPDATE `classes` SET `name`='math',`age`=99 WHERE id = 1
	tx := db.Model(&data).Where("id = ?", 1).Updates(class{
		Name: "math",
		Age:  99,
	})
	if err := tx.Error; err != nil {
		log.Fatal(err)
	}
	if affected := tx.RowsAffected; affected <= 0 {
		log.Fatal("update fail")
	}
	fmt.Println(data.Id, data.Name, data.Age) //
}

func delete(db *gorm.DB) {
	data := class{
		Id:   1,
		Name: "degree",
	}
	//DELETE FROM `classes` WHERE `classes`.`id` = 1
	db.Delete(data)
	//DELETE FROM `classes` WHERE `classes`.`id` = 1 AND `classes`.`name` = 'degree'
	db.Delete(class{}, data)
	//DELETE FROM `classes` WHERE id = 1
	db.Where("id = ?", 1).Delete(class{})
	//DELETE FROM `classes` WHERE name = 'Anna'
	tx := db.Where("name = ?", "Anna").Delete(class{})
	if affected := tx.RowsAffected; affected <= 0 {
		log.Fatal("delete fail")
	}
}

func affair(db *gorm.DB) {
	data := class{
		Id:   996,
		Name: "hello world",
		Age:  996,
	}
	begin := db.Begin()
	begin.Create(&data)
	fmt.Println(data.Id) //996
	//begin.Commit()
	begin.Rollback()

	newData := class{}
	tx := db.First(&newData, data)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("not found") //NotFound
	}
}

func First(db *gorm.DB) {
	var data class
	//SELECT * FROM `classes` WHERE id = 997 ORDER BY `classes`.`id` LIMIT 1
	first := db.Where("id = ?", 996).First(&data)
	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		tx := db.Create(class{
			Id:   996,
			Name: "996_name",
			Age:  996,
		})
		if tx.Error != nil {
			log.Fatal("insert 996 fail")
		}
	}

	db.Where("id = ?", 996).First(&data)
	fmt.Println(data) //{996 996_name 996}

	tx := db.Where("id = ?", 997).First(&data)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		log.Fatal("not found") //not found
	}
	fmt.Println(data)
}

func Find(db *gorm.DB) {
	var classes []class
	db.Create(class{
		Id:   1,
		Name: "1",
		Age:  1,
	})
	db.Create(class{
		Id:   2,
		Name: "2",
		Age:  2,
	})
	// SELECT * FROM `classes` ORDER BY id desc LIMIT 2
	db.Offset(0).Limit(2).Order("id desc").Find(&classes)
	fmt.Println(classes) //[{2 2 2} {1 1 1}]
}

func monitoring(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	stats := s.Stats()
	fmt.Println("因为超过最大存活时间被关闭的数量", stats.MaxLifetimeClosed)
	fmt.Println("因为超过最大空闲连接被关闭的数量", stats.MaxIdleClosed)
	fmt.Println("空闲连接因为超过空闲的最大生存周期被关闭的数量", stats.MaxIdleTimeClosed)
	fmt.Println("等待数量", stats.WaitCount)
	fmt.Println("等待时间", stats.WaitDuration)
	fmt.Println("打开的连接数", stats.OpenConnections)
}

func main() {
	connect := getConnect()
	monitoring(connect)
}
