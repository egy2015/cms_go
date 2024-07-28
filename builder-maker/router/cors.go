package router

import "github.com/gin-contrib/cors"

var corsConfig = cors.Config{
	AllowAllOrigins: true,
	AllowHeaders:    []string{"Authorization", "Content-Type"},
	AllowMethods:    []string{"DELETE", "GET", "POST", "PUT"},
}
