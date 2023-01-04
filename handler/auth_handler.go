package handler

import (
	"context"
	_ "database/sql"
	"github.com/goccy/go-json"
	"go-blog/errno"
	"go-blog/model"
	"go-blog/service"
	"go-blog/store"
	"go-blog/utils"
	"go-blog/vo"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_Token      = "生成token失败"
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
			"code":    1,
			"message": "参数解析异常",
		})
		return
	}
	result := store.DB.Debug().Where("username = ? AND password = ? AND valid = 0", user.Username, user.Password).First(&user)
	if result.Error != nil {
		//用户不存在
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{
				"code":    errno.ERROR_USERNAME_NOT_EXIST,
				"message": errno.GetErrmsg(errno.ERROR_USERNAME_NOT_EXIST),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    1,
				"message": errno.GetErrmsg(errno.ERROR),
			})
			return
		}
		return
	}
	userInfo := &vo.User{
		Id:           user.Id,
		CreateTime:   user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:   user.UpdateTime.Format("2006-01-02 15:04:05"),
		Username:     user.Username,
		HeadPortrait: user.HeadPortrait,
		Password:     user.Password,
	}
	//获取token 存入redis
	token := utils.CreateToken() + "#" + strconv.Itoa(int(userInfo.Id))
	userInfoJSON, err := json.Marshal(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp("1", err.Error()))
		return
	}
	store.RedisClient.Set(store.Ctx, token, userInfoJSON, time.Hour*12)
	m := make(map[string]interface{})
	m["token"] = token
	m["user_Info"] = userInfo
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": m,
	})
}

/**
 * @Description 注册
 * @Param
 * @return
 **/
func Register(c *gin.Context) {
	var req vo.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"message": "参数解析异常",
		})
	}

	if err := service.CreateUser(context.Background(), req.UserName, req.PassWord); err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.ConstructResp("0", "注册成功", nil))
	return
}

/**
 * @Description 退出登录
 * @Param
 * @return
 **/
func Logout(c *gin.Context) {
	token := c.GetHeader("token")
	err := service.Logout(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructResp("1", "退出登录失败!", nil))
	}
	c.JSON(http.StatusOK, errno.ConstructResp("0", "退出登录成功", nil))
	return
}

/**
 * @Description 身份鉴权
 * @Param
 * @return
 **/
//func Verify(c *gin.Context) (result bool, userName string, err error) {
//	strToken := c.GetHeader("token")
//	claim, err := VerifyAction(strToken)
//	if err != nil {
//		result = false
//		return result, "", nil
//	}
//	result = true
//	userName = claim.Username
//	return result, userName, nil
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
//	return claims, nil
//}
//
//func GetToken(claims *JWTClaims) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	signedToken, err := token.SignedString([]byte(Secret))
//	if err != nil {
//		return "", errors.New(ErrorReason_Token)
//	}
//	return signedToken, nil
//}

type Time time.Time
