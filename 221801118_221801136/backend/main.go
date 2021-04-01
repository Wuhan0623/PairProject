package main

import (
	_ "backend/database"
	"backend/thesisSearch"
	"backend/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {

	r := gin.Default()
	r.Use(cors.Default())
	// username=?&&password=?
	r.POST("/user/login/check", user.UserLogin)

	// username=?&&password=?
	r.POST("/user/register/check", user.UserRegister)

	// username=?&&thesis_id=?
	r.POST("/user/like/delete", user.UserLikeDelete)

	// username=?
	r.POST("/user/like", user.UserLikeShow)

	// username=?&&thesis_id=?
	r.POST("/list/like/add", user.UserLikeAdd)

	// source=?&&keyword=?&&year=?&&page=?
	r.POST("/list/search", thesisSearch.GetThesisList)

	//
	r.POST("/list", thesisSearch.GetThesisList)

	r.POST("/")
	r.Run(":8080")
}

