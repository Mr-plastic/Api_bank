package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var account = 10000

func main() {
	router := gin.Default()
	router.GET("/account", getBalance)
	router.GET("/deposit/:input", deposit)
	router.GET("/withdraw/:input", withdraw)
	router.Run(":80")
}

func getBalance(c *gin.Context) {
	var msg = "您的帳戶有:" + strconv.Itoa(account) + "元"
	c.JSON(http.StatusOK, gin.H{
		"amount":  account,
		"status":  "OK",
		"message": msg,
	})
}

func deposit(context *gin.Context) {
	var status string
	var msg string

	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	fmt.Println(err)
	fmt.Println(amount)
	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "操作失敗，存款金額需大於0元！"
		} else {
			account += amount
			status = "ok"
			msg = "已成功存款" + strconv.Itoa(amount) + "元"
		}
	} else {
		amount = 0
		status = "failed"
		msg = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}
func withdraw(c *gin.Context) {
	var msg string
	var status string

	input := c.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount < 0 {
			amount = 0
			msg = "金額不可以小於0，請重新輸入"
			status = "false"
		} else {
			if account-amount < 0 {
				amount = 0
				status = "failed"
				msg = "錢不夠!"
			} else {
				account = account - amount
				status = "ok"
				msg = "成功提款" + strconv.Itoa(amount) + "元"
			}
		}
	} else {
		amount = 0
		status = "false"
		msg = "請重新輸入"
	}
	c.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}
