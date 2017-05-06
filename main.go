package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var beep_state = false

func hello(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	fmt.Println(req.URL.Path)
	fmt.Fprint(res, "Hello world")

}

func main() {
	beep_state = true

	// 建立 Gin 引擎。
	router := gin.Default()

	// 監聽 GET /path 路徑。
	router.GET("/path", func(c *gin.Context) {
		// 回傳 200 OK 的 JSON 資料：{ "test": "wow!!" }。
		c.JSON(200, gin.H{
			"test": "wow!!",
		})
	})

	// 在 :8080 啟動伺服器。
	router.Run(":8080")
}
