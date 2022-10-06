package handler

import (
	_ "database/sql"
	"errors"
	"go-blog/model"
	"go-blog/store"
	"go-blog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

//func SayHello(c *gin.Context) {
//	strToken := c.Param("token")
//	claim, err := verifyAction(strToken)
//	if err != nil {
//		c.String(http.StatusNotFound, err.Error())
//		return
//	}
//	c.String(http.StatusOK, "hello,", claim.Username)
//}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID   int64  `json:"user_id"`
	Password string `json:"password"`
	Username string `json:"username"`
	RoleId   int64  `json:"role_id"`
}

/**
 * @Description 后台系统登录
 * @Param
 * @return
 **/
func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"message": "参数解析异常",
		})
	}
	result := store.DB.Where("username = ? AND password = ? AND delete_flag = 0", user.Username, user.Password).First(&user)
	if result.Error != nil {
		//用户不存在
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{
				"code":    errmsg.ERROR_USERNAME_NOT_EXIST,
				"message": errmsg.GetErrmsg(errmsg.ERROR_USERNAME_NOT_EXIST),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    0,
				"message": errmsg.GetErrmsg(errmsg.ERROR),
			})
		}
		return
	}
	claims := &JWTClaims{
		UserID:   user.Id,
		Username: user.Username,
		Password: user.Password,
		RoleId:   user.RoleId,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	//获取token
	signedToken, err := GetToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	m := make(map[string]interface{})
	m["token"] = signedToken
	m["user_Info"] = user
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": m,
	})
}

//func Verify(c *gin.Context) {
//	strToken := c.Param("token")
//	claim, err := verifyAction(strToken)
//	if err != nil {
//		c.String(http.StatusNotFound, err.Error())
//		return
//	}
//	c.String(http.StatusOK, "verify,", claim.Username)
//}

//func Refresh(c *gin.Context) {
//	strToken := c.Param("token")
//	claims, err := verifyAction(strToken)
//	if err != nil {
//		c.String(http.StatusNotFound, err.Error())
//		return
//	}
//	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
//	signedToken, err := getToken(claims)
//	if err != nil {
//		c.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.String(http.StatusOK, signedToken)
//}

//func VerifyAction(strToken string) (*JWTClaims, error) {
//	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(Secret), nil
//	})
//	if err != nil {
//		return nil, errors.New(ErrorReason_ServerBusy)
//	}
//	claims, ok := token.Claims.(*JWTClaims)
//	if !ok {
//		return nil, errors.New(ErrorReason_ReLogin)
//	}
//	if err := token.Claims.Valid(); err != nil {
//		return nil, errors.New(ErrorReason_ReLogin)
//	}
//	fmt.Println("verify")
//	return claims, nil
//}

func GetToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	split := strings.Split(signedToken, ".")
	signedToken = split[2]
	return signedToken, nil
}
