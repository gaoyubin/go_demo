package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerNumber int64  `gorm:"column:customerNumber"`
	CustomerName   string `gorm:"column:customerName"`
	Country        string `gorm:"column:country"`
	//Phone          string `gorm:"column:phone"`
	Phone  string
	Orders []Order `gorm:"foreignKey:CustomerNumber;references:CustomerNumber"`
}

type Order struct {
	OrderNumber    int64         `gorm:"column:orderNumber"`
	Status         string        `gorm:"column:status"`
	CustomerNumber int64         `gorm:"column:customerNumber"`
	OrderDetails   []OrderDetail `gorm:"foreignkey:OrderNumber;references:OrderNumber"`
}

//	func (v Orderx) TableName() string {
//		return "orders"
//	}
type OrderDetail struct {
	OrderNumber     int64  `gorm:"column:orderNumber"`
	ProductCode     string `gorm:"column:productCode"`
	OrderLineNumber int64  `gorm:"column:orderLineNumber"`
}

func (o OrderDetail) TableName() string {
	return "orderdetails"
}

//func main() {
//	//gorm.Open
//	dsn := "root:172269QQWWpai@tcp(127.0.0.1:3306)/yiibaidb?charset=utf8mb4&parseTime=True&loc=Local"
//	dialog := mysql.Open(dsn)
//	db, err := gorm.Open(dialog, &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(db, err)

//var customer Customer
//db.Debug().First(&customer, 131)
//fmt.Printf("%#v", customer)
//
//db.Debug().Where("customerName=?", "Land of Toys Inc.").Find(&customer)
//fmt.Println(customer)

//order_list := []Orderx{}
//db.Debug().Where("customerNumber=?", 131).Find(&order_list)
//fmt.Println(order_list)

//db.Debug().Preload("orders").Where("customerNumber=?", 131).Find(&customer)
//db.Debug().Association("orders")

//db.Debug().Preload("orders").Where("customerNumber=?", 131).Find(&customer)
//db.Debug().Preload("orders").Where("customerNumber=?", 131).Find(&customer)
//db.Debug().Preload("Orders").Where("customerNumber=?", 131).Find(&customer)
//fmt.Println(customer)

//orderdetail_list := []OrderDetail{}

//orderlist := []Order{}
//db.Debug().Preload("OrderDetails").Where("orderNumber=?", 10329).Find(&orderlist)
//fmt.Println(orderlist)

// }

type User struct {
	//ID   uint `gorm:"primarykey"`
	gorm.Model
	Name string
	Age  uint
}

func main() {
	dsn := "root:172269QQWWpai@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dialog := mysql.Open(dsn)
	db, err := gorm.Open(dialog, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db, err)

	//err = db.Migrator().DropTable(&User{})
	//if err != nil {
	//	panic(err)
	//}
	//err = db.Migrator().CreateTable(&User{})
	//if err != nil {
	//	panic(err)
	//}

	//user := User{
	//	Name: "jinzhu",
	//	Age:  19,
	//}
	//db.Create(&user)
	//fmt.Println(user)

	//users := []User{}
	//var user1 User
	//db.Debug().Where("id=?", 1).Find(&user1)
	//fmt.Println(user1)
	//user1.Name = "low"
	//db.Debug().Save(&user1)
	//db.Debug().Model(&user1).Where("id=?", 4).Update("name", "apple")
	//user1.ID = 5
	////db.Debug().Model(&user1).Update("name", "gao")
	//db.Debug().Delete(&user1)

	//user_list := []User

	//var user_list []User
	//db.Debug().Unscoped().Find(&user_list)
	//fmt.Println(user_list[0])

	//user := User{}
	//user.ID = 1
	//user.Age = 30
	//err := conn.Model(activity).Clauses(clause.Returning{}).Update("detail", activity.Detail).Error

	//var users []User
	//result := db.Model(&users).Clauses(clause.Returning{}).Where("id = ?", 1).Update("age", 70)
	////result := db.Debug().Model(&user).Clauses(clause.Returning{}).Update("age", 30)
	//fmt.Println(result.RowsAffected, result.Error)
	//fmt.Println(users)

	//var res User
	//db.Debug().Raw("select * from users where id=2").Scan(&res)
	//fmt.Println(res)

	//rows, err := db.Raw("select * from users").Rows()
	//if err != nil {
	//	panic(err)
	//}
	//for rows.Next() {
	//	//var name string
	//	//var age int64
	//	//rows.Scan(&res)
	//	db.ScanRows(rows, &res)
	//	fmt.Println(res)
	//
	//}

	//id_list := []int64{1, 2, 3, 4, 5}
	//var users []User
	//result := db.Debug().Model(&users).Where("id in (?)", id_list).Find(&users)
	//fmt.Println(users, result.Error, result.RowsAffected)

	user := User{Age: 19}
	//user.Age = 0
	////user.Name = ""
	user_list := []User{}
	err = db.Debug().Model(User{}).Where(&user).Find(&user_list).Limit(10).Error
	fmt.Println(err, user_list)

	//user_filter := User{}
	//user_filter.ID = 2
	//user_update := User{Age: 0}
	//
	//result := db.Debug().Model(&user_filter).Updates(&user_update)
	//fmt.Println(result.Error, result.RowsAffected, user_update)

}
