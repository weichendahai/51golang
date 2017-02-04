package controllers

import (
	//"../models"
	"gopkg.in/gin-gonic/gin.v1"
	//"net/http"
	//"strconv"
)

// ================== login start ========================
func LoginPage(c *gin.Context) {
	//根据完整文件名渲染模板，并传递参数
	//c.HTML(http.StatusOK, "login.html", gin.H{"title": "Main web"})
	render(c, "login.html", gin.H{"title": "Login"})
}

func LoginAuth(c *gin.Context) {
	render(c, "login-successful.html", gin.H{"title": "Successful Login"})
}

func LoginOut(c *gin.Context) {
	render(c, "login-successful.html", gin.H{"title": "Successful Login"})
}

// ================== login end ========================

// ================== register start ========================
func RegisterPage(c *gin.Context) {
	render(c, "register.html", gin.H{"title": "Register"})
}

func Register(c *gin.Context) {
	render(c, "login-successful.html", gin.H{"title": "Successful registration & Login"})
}

// ================== register end ========================
