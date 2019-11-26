package actions

import (
	"github.com/gin-gonic/gin"
	"html_parse_api/logics"
	"html_parse_api/services/so_service"
	"html_parse_api/structs"
	"net/http"
)

func ParseSo(c *gin.Context) {
	parseRequest := &structs.ParseRequest{}
	if err := c.BindJSON(parseRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": logics.CodeParseRequestFormError,
			"msg":  "请求格式不正确",
			"err":  err.Error(),
		})
		return
	}

	soSearchInfo, err := so_service.ParseSoSearchInfoFromHtml(parseRequest.Html, parseRequest.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": logics.CodeParseError,
			"msg":  "解析失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": logics.CodeParseSuccess,
		"msg":  "解析成功",
		"data": soSearchInfo,
	})
}
