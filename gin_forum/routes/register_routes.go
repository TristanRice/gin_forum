package routes

import (
    "github.com/gin-contrib/sessions";
    "github.com/gin-gonic/gin";
    "net/http";
    passwords "gin_forum/passwords";
    "gin_forum/models"
)

func RegisterPost( c *gin.Context ) {

    const (
        TEMPLATE = "register.html"
    )

    username, exists_username := c.GetPostFormArray("username")
	password, exists_password := c.GetPostFormArray("password_1")
    password2, exists_password2 := c.GetPostFormArray("password_2")
	email, exists_email := c.GetPostFormArray("email")

	//If the post data does not exist, then send an error
	if !exists_username || !exists_password || !exists_email || !exists_password2 {
		send_authenticate_with_error( TEMPLATE, "Please enter all fields", c)
		return
	}

    if password[0] != password2[0] {
        send_authenticate_with_error( TEMPLATE, "Passwords do not match", c)
        return
    }

    var error_message string
    switch (true) {
        //check that the usernames are valid
        case !passwords.Validate_username(username[0]): error_message = "Your username is not valid"; break;
        case !passwords.Validate_password(password[0]): error_message = "Your password is not valid"; break;
        case !passwords.Validate_email(email[0]): error_message = "Your email is not valid"; break;
    }

    if error_message!="" {
        send_authenticate_with_error( TEMPLATE, error_message, c)
    }

    db := models.Make_connection( )
    defer db.Close( )

    var user models.User
    //Check whether a user with the email or username exists
    if !db.Debug( ).Where( "username = ? OR email = ?", username[0], email[0]).First(&user).RecordNotFound( ) {
        send_authenticate_with_error( TEMPLATE , "A user with that username or email already exists", c)
        return
    }

    User := &models.User {
        Username: username[0],
        Email: email[0],
        Password: passwords.Password_hash(password[0]),
    }

    if err := db.Debug( ).Create( User ).Error; err != nil {
        panic(err)
        send_authenticate_with_error( TEMPLATE, "There was an error creating your account, please try again later", c)
        return
    }

    session := sessions.Default( c )
    session.Set("username", username[0])
    session.Set("flash", "Account created succesfully")
    session.Save( )
    c.Redirect(http.StatusMovedPermanently, "/forum")
}

func RegisterGet( c *gin.Context ) {
    c.HTML(http.StatusOK, "register.html", gin.H{
        "title":"Register",
        "BodyId":"grad",
    })
}
