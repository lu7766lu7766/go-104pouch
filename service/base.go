package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Base(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "result": "success"})
}
