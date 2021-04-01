package thesisSearch

import (
	"backend/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AnalyzedThesis struct {
	ID int			`json:"id"`
	Source string	`json:"source"`
	Year int		`json:"year"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Keyword string	`json:"keyword"`
	Abstract string	`json:"abstract"`
	Link string		`json:"link"`
}

type thesisSearchParam struct {
	Keyword string	`json:"keyword" binding:"required"`
	Page int		`json:"page" binding:"required"`
}


// GetThesisList 模糊搜索、查找相应文章
func GetThesisList(c *gin.Context) {
	param := thesisSearchParam{}
	_ = c.ShouldBindJSON(&param)

	keyword := param.Keyword
	page := param.Page


	var selectStr string = "select * from analyzed_thesis "

	if keyword != "" {
		keyword = "%" + keyword + "%"
		selectStr += " where Keyword like '" + keyword + "' "
	}
	if page >= 1 {
		page --
	} else {
		page = 0
	}
	page *= 4
	pageStr := strconv.Itoa(page)
	selectStr += " limit " + pageStr + " , 4"
	selectStr += " ;"


	//var Keys []string

	rows, _ := database.DB.Raw(selectStr).Rows()
	var ID int
	var Source string
	var Year int
	var Title string
	var Author string
	var Keyword string
	var Abstract string
	var Link string
	var thesisArr []AnalyzedThesis
	for rows.Next() {
		temp := AnalyzedThesis{}
		_ = rows.Scan(&ID, &Source, &Year, &Title, &Author, &Keyword, &Abstract, &Link)
		temp.ID = ID
		temp.Source = Source
		temp.Year = Year
		temp.Title = Title
		temp.Author = Author
		temp.Keyword = Keyword
		temp.Abstract = Abstract
		temp.Link = Link
		thesisArr = append(thesisArr,temp)

		//Keys = append(Keys, Keyword)
	}

	//c.JSON(http.StatusOK, Keys)
	c.JSON(http.StatusOK, thesisArr)
}


