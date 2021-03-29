package thesisSearch

import (
	"backend/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type AnalyzedThesis struct {
	ID int			`json:"id"`
	Source string	`json:"source"`
	Year int		`json:"year"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Keyword string	`json:"keyword"`
	Abstract string	`json:"abstract"`
}


// GetThesisList 模糊搜索、查找相应文章
func GetThesisList(c *gin.Context) {
	source := c.Query("source")
	year := c.Query("year")  //类型是否有问题？
	keyword := c.Query("keyword")
	pageStr := c.Query("page")

	keyword = strings.ReplaceAll(keyword, "\n", "%")
	keyword = strings.ReplaceAll(keyword, "\r", "%")
	keyword = strings.ReplaceAll(keyword, "\t", "%")
	keyword = strings.ReplaceAll(keyword, " ", "%")  //problem:正则表达式无法处理，交前端限制？
	keyword = "%" + keyword + "%"

	var selectStr string = "select * from analyzed_thesis where "
	if source != "" {
		selectStr += "Source = '" + source + "' "
	}
	if year != "" {
		if source != "" {
			selectStr += "and "
		}
		selectStr += "Year = '" + year + "' "
	}
	if keyword != "" {
		if !(source == "" && year == "") {
			selectStr += "and "
		}
		selectStr += "Keyword like '" + keyword + "' "
	}
	if source == "" && year == "" && keyword == "" {
		length := len(selectStr) - 6
		selectStr = selectStr[0:length]
	}
	if pageStr != "" {
		page, _ :=strconv.Atoi(pageStr)
		if page >= 1 {
			page --
		} else {
			page = 0
		}
		page *= 4
		pageStr = strconv.Itoa(page)
		selectStr += " limit " + pageStr + " , 4"
	}
	selectStr += " ;"

	rows, _ := database.DB.Raw(selectStr).Rows()
	var ID int
	var Source string
	var Year int
	var Title string
	var Author string
	var Keyword string
	var Abstract string
	var thesisArr []AnalyzedThesis
	for rows.Next() {
		temp := AnalyzedThesis{}
		_ = rows.Scan(&ID, &Source, &Year, &Title, &Author, &Keyword, &Abstract)
		temp.ID = ID
		temp.Source = Source
		temp.Year = Year
		temp.Title = Title
		temp.Author = Author
		temp.Keyword = Keyword
		temp.Abstract = Abstract
		thesisArr = append(thesisArr,temp)
	}

	c.JSON(http.StatusOK, thesisArr)
}


