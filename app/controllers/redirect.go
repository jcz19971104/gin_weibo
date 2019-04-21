package controllers

import (
	"gin_weibo/app/models"
	"gin_weibo/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Redirect : 路由重定向
func Redirect(c *gin.Context, redirectRoute string) {
	c.Redirect(http.StatusMovedPermanently, config.AppConfig.URL+redirectRoute)
}

// 重定向到用户展示页面
func RedirectToUserShowPage(c *gin.Context, u *models.User) {
	Redirect(c, "/users/show/"+u.GetIDstring())
}

// 重定向到登录页面
func RedirectToLoginPage(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		Redirect(c, "/login")
		return
	}

	Redirect(c, "/login?back="+c.Request.URL.Path)
}

// 重定向到用户创建页面 (注册页面)
func RedirectToUserCreatePage(c *gin.Context) {
	Redirect(c, "/users/create")
}

// 重定向到用户更新编辑页面
func RedirectToUserEditPage(c *gin.Context, idStr string) {
	Redirect(c, "/users/edit/"+idStr)
}

// 重定向到 root page
func RedirectToRootPage(c *gin.Context) {
	Redirect(c, "/")
}