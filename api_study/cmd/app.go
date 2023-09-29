package main

import (
	"fmt"
	"hello/pkg/api"
	"hello/pkg/common"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	v := route.Group("/v1")

	api.UserApi(v)		// ユーザーAPI
	api.TaskApi(v)		// タスクAPI

	// http://localhost:8080/v1/user?category=aaa
	// http://localhost:8080/v1/task/all
	route.Run(":8080")

	// ===============================
	// 確認用
	// ===============================
	var s int = common.Add(1,2);
	d := common.Hiku(8,1);
	fmt.Println(s);
	fmt.Println(d);
}

