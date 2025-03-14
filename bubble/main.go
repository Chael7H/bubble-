package main

import (
	"bubble/databases"
	"bubble/routers"
)

func main() {
	//连接数据库
	err := databases.InitMysql()
	if err != nil {
		panic(err)
	}

	r := routers.SetUpRoutes()

	r.Run(":8080")
}
