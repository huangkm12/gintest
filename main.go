package main

import (
	"bubble/dao"
	"bubble/routers"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)





func main() {
	// 创建数据库

	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("open mysql failed,err:%v", err)
		return
	}
	defer dao.DB.Close()



	// 注册路由
	r := routers.SetupRouter()

	r.Run()
}
