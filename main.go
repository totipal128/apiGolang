package main

import "golang3/routes"

func main(){
	route := routes.GetRoutes()

	route.Run()
}