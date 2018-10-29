package routes

import (
    "github.com/gin-gonic/gin";
    "net/http"
)

func forum(c *gin.Context) {
    /*session := sessions.Default( c )
    if flash := session.Get("flash"); flash != nil {
    	c.HTML()
    }
    */
    //var posts []map[string]string
    posts := []map[string]string {
    	{
    		"title": "title1",
    		"content":"content1",
    	},
    	{
    		"title": "title1",
    		"content":"content1",
    	},
    }

    c.HTML(http.StatusOK, "forum.html", gin.H {
    	"title":"Forum",
    	"loggedin":"true",
    	"posts": posts,
    })
}