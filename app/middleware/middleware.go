package middleware

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
)

// ================== Middleware start ========================

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

// ================== Middleware end ========================
