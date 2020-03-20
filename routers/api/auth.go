package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"cumtBook/models"
	"cumtBook/pkg/e"
	"cumtBook/pkg/logging"
	"cumtBook/pkg/utils"
)

type auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

// GetAuth 获取认证权限
func GetAuth(c *gin.Context) {
	var au auth
	if c.BindJSON(&au) == nil {
		username := au.Username
		password := au.Password

		valid := validation.Validation{}

		a := auth{Username: username, Password: password}
		ok, _ := valid.Valid(&a)

		data := make(map[string]interface{})
		code := e.INVALID_PARAMS
		if ok {
			isExist := models.CheckAuth(username, password)
			if isExist {
				token, err := utils.GenerateToken(username, password)
				if err != nil {
					code = e.ERROR_AUTH_TOKEN
				} else {
					data["token"] = token

					code = e.SUCCESS
				}
			} else {
				code = e.ERROR_AUTH
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

}
