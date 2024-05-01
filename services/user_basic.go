package services

import (
	"fmt"
	"go_chat/helper"
	"go_chat/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// ! Don't use c.Query()
	account := c.PostForm("account")
	password := c.PostForm("password")
	fmt.Println("!!!!!!!!!!!!", account, password)
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "account or password is null.",
		})
		return
	}
	/* If you want to enable md5, uncomment under */
	// ub, err := model.GetAccountAndPassword(account, helper.GetMd5(password))
	/* If you want to enable md5, uncomment upside*/
	ub, err := model.GetAccountAndPassword(account, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "No this account or password.",
		})
		return
	}
	token, err := helper.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "System error,gennerate token failed.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "success",
		"token": token,
	})
}

func UserDetail(c *gin.Context) {
	u, _ := c.Get("user_claims")
	uc := u.(*helper.UserClaims)
	userBasic, err := model.GetUserDetail(uc.Identity)
	if err != nil {
		log.Println("[DB ERROR]", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Database query exception",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "query userdetail success",
		"data": userBasic,
	})
}

func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	fmt.Println("[test]:", email)
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "empty email!",
		})
		return
	}
	cnt, err := model.GetEmailCnt(email)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "The current mailbox has been registered.",
		})
		return
	}
	err = helper.SendCode(email, "666666")
	if err != nil {
		fmt.Println("use sendcode:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "[ERROR] Send code",
		})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "[SUCCESS] Send code success.",
	})
}
