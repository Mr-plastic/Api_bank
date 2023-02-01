package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var account = 10000

type Result struct {
	Money  int    `json:"money"`
	Status string `json:"status"`
	Msg    string `json:"message"`
}

var result = Result{}

func main() {
	router := gin.Default()
	router.GET("/account", getBalance)
	router.GET("/deposit/:input", deposit)
	router.GET("/withdraw/:input", withdraw)
	router.Run(":80")
}

func getBalance(c *gin.Context) {
	result.Money = account
	result.Status = "OK"
	result.Msg = ""
	c.JSON(http.StatusOK, result)
}

func deposit(context *gin.Context) {
	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	fmt.Println(err)
	fmt.Println(amount)
	if err == nil {
		if amount <= 0 {
			result.Money = 0
			result.Status = "failed"
			result.Msg = "操作失敗，存款需要大於0元"
		} else {
			account += amount
			result.Status = "ok"
			result.Msg = "已成功存款" + strconv.Itoa(amount) + "元"
		}
	} else {
		result.Money = 0
		result.Status = "failed"
		result.Msg = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, result)
}
func withdraw(c *gin.Context) {
	input := c.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount < 0 {
			result.Money = 0
			result.Msg = "金額不可以小於0，請重新輸入"
			result.Status = "false"
		} else {
			if account-amount < 0 {
				result.Money = 0
				result.Status = "failed"
				result.Msg = "錢不夠!"
			} else {
				account = account - amount
				result.Status = "ok"
				result.Msg = "成功提款" + strconv.Itoa(amount) + "元"
			}
		}
	} else {
		result.Money = 0
		result.Status = "false"
		result.Msg = "請重新輸入"
	}
	c.JSON(http.StatusOK, result)
}
