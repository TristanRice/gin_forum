package main

import (
    "github.com/gin-gonic/gin"
)

func create_static( r *gin.Engine ) {
    //static for the CSS
    r.Static("/css", "assets/css/")

    //static for the JS
    r.Static("/js", "assets/js")

    //static for images
    r.Static("/img", "assets/img")
}
