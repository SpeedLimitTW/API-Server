package main

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

var beep_state = false

func hello(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	fmt.Println(req.URL.Path)
	fmt.Fprint(res, "Hello world")
}

func main() {

	// 建立 Gin 引擎。
	router := gin.Default()

	// 監聽 GET /isBeep 路徑。
	router.GET("/isBeep", func(c *gin.Context) {
		c.String(200, strconv.FormatBool(beep_state))
	})

	router.GET("/setBeep", func(c *gin.Context) {
		beep_state = false
		c.String(200, "OK")
	})

	router.GET("/questBeep", func(c *gin.Context) {
		beep_state = true
		c.String(200, "setBeepRequestPendding")
	})

	// 在 :8080 啟動伺服器。
	router.Run(":8080")
}
