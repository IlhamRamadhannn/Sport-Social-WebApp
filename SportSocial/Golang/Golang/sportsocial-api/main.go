package main

import (
	// "fmt"
	"log"
	"sportsocial-api/handler"
	"sportsocial-api/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sportsocial-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&user.UserInput{})

	userRespository := user.NewRepository(db)
	userService := user.NewService(userRespository)
	userHandler := handler.NewUserHandler(userService)



	// users, err := userRespository.FindAll()

	// for _, user := range users {
	// 	fmt.Println("username :", user.Username)
	// }

	// user := user.UserInput{}
	// user.Username = "Calvin winatha"
	// user.Age = 19
	// user.Email = "calvin10@gmail.com"
	// user.PhoneNumber = "+627457458811"

	// err = db.Create(&user).Error
	// if err != nil{
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Creating User record")
	// 	fmt.Println("==========================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users", userHandler.GetUsers)
	v1.GET("/users/:id", userHandler.GetUser)
	v1.POST("/users", userHandler.CreateUser)
	v1.PUT("/users/:id", userHandler.UpdateUser)
	v1.DELETE("/users/:id", userHandler.DeleteUser)

	router.Run()
}
