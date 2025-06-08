package main

import "github.com/isaafisyah/studi-kasus-multifinance/app"

func main()  {
	server := app.NewServer()
	server.Run()
}