package main

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"time"
)

type GPS_DATA struct {
	gorm.Model
	GPS       string
	TIME      int64
	DEVICE_ID int
}

var beep_state = false

func hello(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	fmt.Println(req.URL.Path)
	fmt.Fprint(res, "Hello world")
}

func main() {

	db, err := gorm.Open("mysql", "root:rootroot@/OLDMAN?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&GPS_DATA{})

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

	//The line after this comment is gps samulation feature.
	router.GET("/addDevice")

	router.GET("/addGPSData", func(c *gin.Context) {
		la := c.Param("la")
		lo := c.Param("lo")
		device_id := c.Param("device_id")

		deviceIdInteger, _ := strconv.Atoi(device_id)

		fmt.Print(deviceIdInteger)
		db.Create(&GPS_DATA{GPS: la + lo, TIME: time.Now().Unix(), DEVICE_ID: deviceIdInteger})

		c.String(200, "Received data")
	})

	// 在 :8080 啟動伺服器。
	router.Run(":8080")
}
