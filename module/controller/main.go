package controller

import (
	"fmt"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/gin-gonic/gin"
	model "github.com/michaelchandrag/lemonilo-test/model"
)

func Find (c *gin.Context) {
	var opt model.UserProfileContractor
	var userProfile model.UserProfile
	var _userProfile model.UserProfileInterface = userProfile
	result, err := opt.FindContractor(_userProfile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": result,
	})
	return
}

func Add (c *gin.Context) {
	type ReqBody struct {
		Email 		string 		`json:"email"`
		Password 	string 		`json:"password"`
		Address 	string 		`json:"address"`
	}

	var payload ReqBody
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Body request is not correct.",
		})
		return
	}


	userProfile := model.UserProfile{
		Email: payload.Email,
		Password: MD5Hash(payload.Password),
		Address: payload.Address,
	}

	var opt model.UserProfileContractor
	
	var _userProfile model.UserProfileInterface
	_userProfile = userProfile

	result, err := opt.InsertContractor(_userProfile)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error when create user profile.",
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success create user profile.",
		"data": result,
	})
	return
}

func Update (c *gin.Context) {
	type ReqBody struct {
		Email 		string 		`json:"email"`
		Password 	string 		`json:"password"`
		Address 	string 		`json:"address"`
	}

	var payload ReqBody
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Body request is not correct.",
		})
	}

	userID := c.Param("user_id")
	pUserID, _ := strconv.Atoi(userID)
	userProfile := model.UserProfile{
		Email: payload.Email,
		Password: MD5Hash(payload.Password),
		Address: payload.Address,
	}

	var opt model.UserProfileContractor

	var _userProfile model.UserProfileInterface = userProfile

	err := opt.UpdateContractor(pUserID, _userProfile)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"success": false,
			"message": "Update user profile error.",
			"error": err,
		})
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Update user profile success.",
	})
	return
}

func Delete (c *gin.Context) {
	userId := c.Param("user_id")
	cUserId, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Requested parameters is incorrect.",
		})
		return
	}

	userProfile := model.UserProfile{
		UserID: cUserId,
	}

	var opt model.UserProfileContractor
	var _userProfile model.UserProfileInterface = userProfile
	err = opt.DeleteContractor(cUserId, _userProfile)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Delete user profile error",
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Delete user profile success",
	})
	return
}

func Login (c *gin.Context) {
	type ReqBody struct {
		Email 		string 		`json:"email"`
		Password 	string 		`json:"password"`
	}

	var payload ReqBody
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Body request is not correct.",
		})
		return
	}


	userProfile := model.UserProfile{
		Email: payload.Email,
		Password: MD5Hash(payload.Password),
	}

	var opt model.UserProfileContractor
	
	var _userProfile model.UserProfileInterface
	_userProfile = userProfile
	result, err := opt.LoginContractor(userProfile.Email, _userProfile)
	if err != nil || userProfile.Password != result.Password{
		c.JSON(400, gin.H{
			"success": false,
			"message": "Email or password is incorrect.",
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success login.",
		"data": result,
	})
	return
}



func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}