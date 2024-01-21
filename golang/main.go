package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"encoding/base64"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
	"github.com/xlzd/gotp"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"errors"

)


type Login struct {
	Username     string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	RePassword string `form:"repassword" json:"repassword" xml:"repassword"`
	YourOTP string `form:"yourotp" json:"yourotp" xml:"yourotp"`
}

type Userdata struct {
	ID            	int
	Username      	string
	UserPassword    string
	UserSecret    	string
	UserOTPURL     	string
}


func main() {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/myotpdb")
	defer db.Close()
	if err != nil {
		fmt.Println("connect fail")
	} else {
		fmt.Println("connect success")
	}


	router := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())


	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/login", func(c *gin.Context) {

		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userData, err := getUser(db, json.Username, json.Password)
		fmt.Println(userData)

		if(err != nil){
			fmt.Println("error username/password")
			c.JSON(http.StatusForbidden, gin.H{"status": "Cannot Login username/password WRONG"})
			return
		}

		//serverotp := gotp.NewDefaultTOTP(userData.UserSecret).Now()
		//fmt.Println("Server OTP generated %s:%s", serverotp, len(json.YourOTP))
	
		result := gotp.NewDefaultTOTP(userData.UserSecret).Verify(json.YourOTP, time.Now().Unix())

		fmt.Printf("Verify result:%v", result )

		//if(serverotp != json.YourOTP){
		if(!result){	
			fmt.Println("error OTP not match %s", json.YourOTP)
			c.JSON(http.StatusForbidden, gin.H{"status": "Cannot Login because OTP mismatch"})
			return			
		}


		c.JSON(http.StatusOK, gin.H{"status": "Successfully Logged In"})



	})



	router.POST("/register", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		secretLength := 16
		mysecret := gotp.RandomSecret(secretLength) 
	
	
		fmt.Println(mysecret)
	
		totp := gotp.NewDefaultTOTP(mysecret)
		myotpstr := totp.ProvisioningUri(json.Username,"My App OTP")
	
		fmt.Println(myotpstr)
	
		var png []byte
		png, err := qrcode.Encode(myotpstr, qrcode.Medium, 256)
		if( err != nil){
			fmt.Println("Error")
		}

		var base64Encoding string
		base64Encoding += "data:image/png;base64,"

		base64Encoding += base64.StdEncoding.EncodeToString(png)
		fmt.Println(base64Encoding)

		fmt.Println(addUser(db, json.Username, json.Password, mysecret, myotpstr))

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in","png":base64Encoding})
	})

	router.Run(":8080")




}


func getUser(db *sql.DB, username string, userpassword string) (Userdata, error) {
	rows, err := db.Query("SELECT * FROM otpuser WHERE username=? AND userpassword=?",username, userpassword)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var userData Userdata
	if rows.Next() {
		err := rows.Scan(
			&userData.ID,
			&userData.Username,
			&userData.UserPassword,
			&userData.UserSecret,
			&userData.UserOTPURL,

		)
		if err != nil {
			panic(err.Error())
		}

	}else{
		return userData, errors.New("Not found")
	}
	return userData, nil
}


func addUser(db *sql.DB, username string, password string, usersecret string, userotpurl string) bool {
	statement, _ := db.Prepare(`INSERT INTO otpuser(username, userpassword,
			usersecret, userotpurl
		)VALUES (?,?,?,?)`)

	_, err := statement.Exec(username, password, usersecret, userotpurl)

	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

