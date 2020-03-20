package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"cumtBook/models"
	"cumtBook/pkg/e"
	"cumtBook/pkg/logging"
	"cumtBook/pkg/setting"
	"cumtBook/pkg/utils"
)

// Book POST请求json对象接受体
type Book struct {
	BookID    int    `json:"book_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Org       string `json:"org"`
	Phone     string `json:"phone" binding:"required"`
	Date      string `json:"date"`
	Last      int    `json:"last"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	CreatedBy string `json:"created_by"`
	State     int    `json:"state"`
}

// GetBook 获取单个预约 此处的id指的是book_id
func GetBook(c *gin.Context) {
	bookID := com.StrTo(c.Param("book_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(bookID, 1, "book_id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistBookById(bookID) {
			data = models.GetBook(bookID)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_BOOK
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetBooks 获取多个预约
func GetBooks(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var bookID int = -1
	if arg := c.Query("book_id"); arg != "" {
		bookID = com.StrTo(arg).MustInt()
		maps["book_id"] = bookID

		valid.Min(bookID, 1, "book_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetBooks(utils.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetBookTotal(maps)

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddBook 新增预约
func AddBook(c *gin.Context) {
	var book Book
	data := make(map[string]interface{})
	if c.BindJSON(&book) == nil {
		bookID := book.BookID
		name := book.Name
		org := book.Org
		phone := book.Phone
		date := book.Date
		last := book.Last
		desc := book.Desc
		content := book.Content
		createdBy := book.CreatedBy
		state := book.State

		// 数据校验
		valid := validation.Validation{}
		valid.Min(bookID, 1, "book_id").Message("预约ID必须大于0")
		valid.Required(name, "name").Message("负责人姓名不能为空")
		valid.Required(org, "org").Message("预约人单位不能为空")
		valid.Required(phone, "phone").Message("负责人电话不能为空")
		valid.Required(date, "date").Message("预约日期不能为空")
		valid.Required(last, "last").Message("请填写预约时长")
		valid.Required(desc, "desc").Message("请填写预约信息简述")
		valid.Required(content, "content").Message("内容不能为空")
		valid.Required(createdBy, "created_by").Message("创建人不能为空")
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

		code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			if !models.ExistBookById(bookID) {
				data["book_id"] = bookID
				data["name"] = name
				data["org"] = org
				data["phone"] = phone
				data["date"] = date
				data["last"] = last
				data["desc"] = desc
				data["content"] = content
				data["created_by"] = createdBy
				data["state"] = state

				models.AddBook(data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_EXISTED_BOOK
			}
		} else {
			for _, err := range valid.Errors {
				logging.Info(err.Key, err.Message)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	} else {
		code := e.INVALID_PARAMS
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}
}

// 删除预约
func DelBook(c *gin.Context) {
	bookID := com.StrTo(c.Param("book_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(bookID, 1, "book_id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistBookById(bookID) {
			models.DeleteBook(bookID)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_BOOK
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改预约
func EditBook(c *gin.Context) {

}
