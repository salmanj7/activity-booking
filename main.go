package main

import(
	"github.com/gin-gonic/gin"
	"github.com/salmanj7/activity-booking/config"
	"github.com/salmanj7/activity-booking/route"
)

func main(){
	r := gin.Default()
    config.ConnectDatabase()
    route.RegisterRoutes(r)
    r.Run(":8080")
}