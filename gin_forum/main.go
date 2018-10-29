package main

import (
    "log";
    "github.com/gin-gonic/gin";
    "gin_forum/routes";
    "gin_forum/models";
    "flag"
)

const (
    SERVER_HOST = "localhost"
    SERVER_PORT = "8080"
)

func main( ) {

    //Parse the command line arguments, so we can migrate easily. 
    migrate := flag.Bool("migrate", false, "When migrating to a new server, or changing the models")
    flag.Parse( )

    //create the router
    router := gin.Default( )

    //load the templates
    router.LoadHTMLGlob("templates/*")

    //create the static files
    create_static(router)

    //create the routes
    routes.Create(router)

    //alert the admin that we're starting the server
    log.Printf("Listening on: %s:%s", SERVER_HOST, SERVER_PORT)

    //Migrate the models
    if *migrate { models.Migrate( ) } //Remember to uncomment this when migrating to a new datbase, with new fields and shit
    
    //NOw start the server
    router.Run(SERVER_HOST+":"+SERVER_PORT)
}
