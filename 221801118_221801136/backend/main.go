package main

import (
	_ "backend/database"
	"backend/thesisSearch"
	"backend/user"
	"github.com/gin-gonic/gin"
)


func main() {

	r := gin.Default()

	// username=?&&password=?
	r.GET("/user/login", user.UserLogin)

	// username=?&&password=?
	r.GET("/user/register", user.UserRegister)

	// username=?&&thesis_id=?
	r.GET("/user/add/like", user.UserAddLike)

	// username=?
	r.GET("/user/show/like", user.UserShowLike)

	// source=?&&keyword=?&&year=?
	r.GET("/search/list", thesisSearch.GetThesisList)

	//r.POST("/user", user.HandleList)
	r.GET("/")
	r.Run(":8080")
}

