package routes

import (
    "github.com/gin-contrib/sessions";
    "github.com/gin-contrib/sessions/cookie";
    "github.com/gin-gonic/gin";
    "net/http"
)

func Create( router *gin.Engine ) {
    //Create the cookie
    store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

    //Logon routes
    router.GET("/login", LoginGet)
    router.POST("/login", LoginPost)
    router.GET("/logout", Logout)

    //Register routes
    router.GET("/register", RegisterGet)
    router.POST("/register", RegisterPost)

    //forum routes
    router.GET("/forum", forum)

    //making post routes
    create = router.Group("/make")
    create.GET("/post")
}

func send_authenticate_with_error( templates string, error string, c *gin.Context ) {
    c.HTML(http.StatusOK, templates, gin.H {
        "title":"register",
        "BodyId":"grad",
        "error":error,
    })
}
