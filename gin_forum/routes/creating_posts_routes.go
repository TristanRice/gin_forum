package routes

import (
    "github.com/gin-gonic/gin";
    "net/http"
)

func makePostGet( c *gin.Context ) {
	c.HTML(http.StatusOK, "createpost.html", gin.H{
		"title":"Create post",
	})
}