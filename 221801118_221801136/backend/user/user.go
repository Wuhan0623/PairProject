package user

import (
	"backend/database"
	_ "backend/database"
	"backend/thesisSearch"
	_ "database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strings"
)

type loginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type userLikeParam struct {
	Username string `json:"username" binding:"required"`
	ThesisID int `json:"thesis_id" binding:"required"`
}

type showUserLikeParam struct {
	Username string `json:"username" binding:"required"`
}

// UserLogin 用户登录
func UserLogin(c *gin.Context){
	param := loginParam{}
	_ = c.ShouldBindJSON(&param)
	username := param.Username
	password := param.Password

	ifExist, _ := database.DB.Raw("select * from user where username = ? and password = ?", username, password).Rows()
	var i int
	if ifExist.Next() {
		i++
	}

	if (i == 1) {
		c.JSON(http.StatusOK, gin.H{"ifCorrect":true})
	} else {
		c.JSON(http.StatusOK, gin.H{"ifCorrect":false})
	}

}


// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	param := loginParam{}
	_ = c.ShouldBindJSON(&param)
	username := param.Username
	password := param.Password

	rows, _ := database.DB.Raw("select * from user where username = ? ", username).Rows()
	var i int
	if rows.Next() {
		i++
	}

	if (i == 0 && len(username) <= 60 && len(password) >= 6) {
		affected := database.DB.Exec("insert into user values(?, ?)", username, password).RowsAffected
		if(affected == 1){
			c.JSON(http.StatusOK, gin.H{"ifSucceed":true})
		} else {
			c.JSON(http.StatusOK, gin.H{"ifSucceed":false})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"ifSucceed":false})
	}
}


// UserAddLike 用户增加收藏
func UserLikeAdd(c *gin.Context) {
	param := userLikeParam{}
	_ = c.ShouldBindJSON(&param)
	username := param.Username
	thesisID := param.ThesisID

	//先判断用户和论文是否都存在
	user, _ := database.DB.Raw("select * from user where username = ? ;", username).Rows()
	var i int
	if user.Next() {
		i++
	}
	thesis, _ := database.DB.Raw("select * from analyzed_thesis where ID = ? ;", thesisID).Rows()
	var j int
	if thesis.Next() {
		j++
	}
	//判断该用户是否已收藏该文章
	rows, _ := database.DB.Raw("select * from user_like where username = ? and thesis_id = ? ;", username, thesisID).Rows()
	var k int
	if rows.Next() {
		k++
	}

	if i>0 && j>0 && k==0 {
		x := database.DB.Exec("insert into user_like(username, thesis_id) values(?,?)", username, thesisID).RowsAffected
		if( x==1 ) {
			c.JSON(http.StatusOK, gin.H{"ifSucceed":true})
		} else {
			c.JSON(http.StatusOK, gin.H{"ifSucceed":false})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"ifSucceed":false})
	}

}


// UserDeleteLike 用户删除收藏
func UserLikeDelete(c *gin.Context) {
	param := userLikeParam{}
	_ = c.ShouldBindJSON(&param)
	username := param.Username
	thesisID := param.ThesisID

	//先判断用户和论文是否都存在
	user, _ := database.DB.Raw("select * from user where username = ? ", username).Rows()
	var i int
	if user.Next() {
		i++
	}
	thesis, _ := database.DB.Raw("select * from analyzed_thesis where ID = ? ", thesisID).Rows()
	var j int
	if thesis.Next() {
		j++
	}
	//判断该用户是否已收藏该文章
	rows, _ := database.DB.Raw("select * from user_like where username = ? and thesis_id = ? ", username, thesisID).Rows()
	var k int
	if rows.Next() {
		k++
	}

	if i>0 && j>0 && k==1 {
		x := database.DB.Exec("delete from user_like where username = ? and thesis_id = ?", username, thesisID).RowsAffected
		if( x==1 ) {
			c.JSON(http.StatusOK, gin.H{"ifSucceed":true})
		} else {
			c.JSON(http.StatusOK, gin.H{"ifSucceed":false})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"ifSucceed":false})
	}

}


// UserShowLike 返回用户收藏列表
func UserLikeShow(c *gin.Context) {
	param := showUserLikeParam{}
	_ = c.ShouldBindJSON(&param)
	username := param.Username

	//取出该用户所有收藏文章的ID号
	idRows ,_ := database.DB.Raw("select thesis_id from user_like where username = ?", username).Rows()
	var thesisIdArr []int
	var tempID int
	for idRows.Next() {
		_ = idRows.Scan(&tempID)
		thesisIdArr = append(thesisIdArr, tempID)
	}

	//根据收藏ID数组查找论文

	var ID int
	var Source string
	var Year int
	var Title string
	var Author string
	var Keyword string
	var Abstract string
	var Link string
	var thesisArr []thesisSearch.AnalyzedThesis
	for _, eachID := range thesisIdArr {
		database.DB.Raw("select * from analyzed_thesis where ID = ? ;", eachID).Row().Scan(&ID, &Source, &Year, &Title, &Author, &Keyword, &Abstract, &Link)
		temp := thesisSearch.AnalyzedThesis{}
		temp.ID = ID
		temp.Source = Source
		temp.Year = Year
		temp.Title = Title
		temp.Author = Author
		temp.Keyword = Keyword
		temp.Abstract = Abstract
		temp.Link = Link
		thesisArr = append(thesisArr,temp)
	}

	c.JSON(http.StatusOK, thesisArr)
}
