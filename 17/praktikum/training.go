package main

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	ID    int `json: "id" 		form :"id"`
// 	Name  string `json: "name"  form :"name "`
// 	Email string `json: "email" form :"email"`
// }

// var DB *gorm.DB	

// func initDB() {
// 	dsn := "root:@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if (err != nil) {
// 		panic(err.Error())
// 	}
// }

// func initMigrate(){
// 	DB.AutoMigrate(&User{ }) 
// }


// func main() {
// 	initDB()
// 	e := echo.New()

// 	// routing (request)
// 	e.GET("/users", GetUserController)
// 	e.POST("/users", CreateUserController )
	
	 
// 	e.Start(":8000")
// }

// // controller (respon)
// func GetUserController(e echo.Context) error {
// 	var user []User

// 	err := DB.Find(&user).Error
// 	if err != nil{
// 		return e.JSON(http.StatusOK, map[string] interface{} {
// 			"massages" : "success create user"
// 			"user", user
// 		})
		
// 	}
// 	return e.JSON(http.StatusOK, map[string] interface{} {
// 		"massages" : "success",
// 		"data", users
// 	})

// }
	
	

// func CreateUserController(c echo.Context) error{
// 	// email := c.FormValue("email")
// 	// name := c.FormValue("name")
	
// 	user := User{}
// 	c.Bind(&user)

// 	return e.JSON(http.StatusOK, map[string] interface{} {
// 		"massages" : "success create user"
// 		"user", user
// 	})
// }