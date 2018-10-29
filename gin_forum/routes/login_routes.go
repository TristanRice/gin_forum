package routes

import (
    "github.com/gin-contrib/sessions";
    "github.com/gin-gonic/gin";
    "net/http";
    "gin_forum/passwords";
    "gin_forum/models"
)

const (
    FORUM = "/forum"
)

func Logout( c *gin.Context ) {
    session := sessions.Default( c )
    session.Options(sessions.Options{MaxAge: -1})
    session.Clear( ) // or whatever changes session.written to true
    session.Save( )
    c.Redirect(http.StatusMovedPermanently, FORUM)
}

func LoginGet( c *gin.Context ) {
    c.HTML(http.StatusOK, "login.html", gin.H{
        "title":"Register",
        "BodyId":"grad",
    })
}

func LoginPost(c *gin.Context) {

    const (
        TEMPLATE = "login.html"
    )

    username, exists_username := c.GetPostFormArray("username")
    password, exists_password := c.GetPostFormArray("password")

    if !exists_username || !exists_password {
        send_authenticate_with_error(TEMPLATE, "Please enter all fields", c)
        return
    }

    db := models.Make_connection( )
    defer db.Close( )

    var User models.User
    if err := db.Debug( ).Where("username = ?", username[0]).First(&User).Error; err != nil {
        send_authenticate_with_error(TEMPLATE, "There was an error while loggin you in", c)
        return
    }

    if !passwords.Password_verify(User.Password, password[0]) {
        send_authenticate_with_error(TEMPLATE, "Incorrect credentials", c)
        return
    }

    session := sessions.Default( c )
    session.Set("username", username[0])
    session.Set("flash", "Logged in succesfully")
    session.Save( )
    c.Redirect(http.StatusMovedPermanently, FORUM)
}
