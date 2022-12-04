package main

import (
	"goCSA/Week_04/api"
	"goCSA/Week_04/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
