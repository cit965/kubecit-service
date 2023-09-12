package gin

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"kubecit-service/ent"
	"kubecit-service/ent/account"
	"kubecit-service/ent/user"
	"kubecit-service/internal/pkg/jwt"
	"net/http"
	"net/url"
	"sort"
)

import "embed"

//go:embed all:dist
var StaticFiles embed.FS

type GinService struct {
	*gin.Engine
}

const (
	TOKEN     = "111"
	AppId     = "wx4edbdc4895597796"
	AppSecret = "cea927f6ae8a5974a0ce3e266a4cad0b"
)

func NewGinService() *GinService {
	r := gin.Default()
	r.Use(Cors())
	// static files

	//staticFS, _ := fs.Sub(StaticFiles, "dist")
	//r.StaticFS("web/", http.FS(staticFS))
	// 签名校验
	r.GET("/web/wechat/check", CheckSignature)
	//二维码生成
	r.GET("/web/wechat/login", Redirect)
	// 回调地址
	r.GET("/web/wechat/callback", Callback)
	return &GinService{r}
}

func CheckSignature(c *gin.Context) {
	// 获取查询参数中的签名、时间戳和随机数
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	fmt.Println(signature)
	// 创建包含令牌、时间戳和随机数的字符串切片
	tmpArr := []string{TOKEN, timestamp, nonce}
	// 对切片进行字典排序
	sort.Strings(tmpArr)
	// 将排序后的元素拼接成单个字符串
	tmpStr := ""
	for _, v := range tmpArr {
		tmpStr += v
	}
	// 对字符串进行SHA-1哈希计算
	tmpHash := sha1.New()
	tmpHash.Write([]byte(tmpStr))
	tmpStr = fmt.Sprintf("%x", tmpHash.Sum(nil))
	fmt.Println(tmpStr)
	fmt.Println(signature)
	// 将计算得到的签名与请求中提供的签名进行比较，并根据结果发送相应的响应
	if tmpStr == signature {
		c.String(200, echostr)
	} else {
		c.String(403, "签名验证失败 "+timestamp)
	}
}

func Redirect(c *gin.Context) {
	path := c.Query("path")
	state := GetRandomString(5)                                               //防止跨站请求伪造攻击 增加安全性
	redirectURL := url.QueryEscape("http://" + path + "/web/wechat/callback") //userinfo,
	wechatLoginURL := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&state=%s&scope=snsapi_userinfo#wechat_redirect", AppId, redirectURL, state)
	wechatLoginURL, _ = url.QueryUnescape(wechatLoginURL)
	// 生成二维码
	qrCode, err := qrcode.Encode(wechatLoginURL, qrcode.Medium, 256)
	if err != nil {
		// 错误处理
		c.String(http.StatusInternalServerError, "Error generating QR code")
		return
	}
	// 将二维码图片作为响应返回给用户
	c.Header("Content-Type", "image/png")
	c.Writer.Write(qrCode)
}

func Callback(ctx *gin.Context) {
	// 获取微信返回的授权码
	code := ctx.Query("code")
	// 向微信服务器发送请求，获取access_token和openid
	tokenResp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", AppId, AppSecret, code))
	if err != nil {
		FailWithMessage("获取token失败", ctx)
		return
	}
	//// 解析响应中的access_token和openid
	var tokenData struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenID       string `json:"openid"`
	}
	if err1 := json.NewDecoder(tokenResp.Body).Decode(&tokenData); err1 != nil {
		FailWithMessage("获取token失败", ctx)
		return
	}
	userInfoURL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", tokenData.AccessToken, tokenData.OpenID)
	userInfoResp, err := http.Get(userInfoURL)
	if err != nil {
		FailWithMessage("获取用户信息失败", ctx)
		return
	}
	defer userInfoResp.Body.Close()

	var userData struct {
		OpenID   string `json:"openid"`
		Nickname string `json:"nickname"`
	}
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userData); err != nil {
		FailWithMessage("解析微信用户信息失败", ctx)
		return
	}
	entClient, err := ent.Open("sqlite3", "./test.db?_fk=1")
	if err != nil {
		FailWithMessage("连接数据库失败", ctx)
		return
	}
	// 查询账号是否存在
	ac, err := entClient.Account.Query().Where(account.OpenidEQ(userData.OpenID)).First(ctx)
	// 账号不存在,则创建用户和账号。 用户是唯一的，只有创建了用户才能有账号
	if ent.IsNotFound(err) {
		err := WithTx(ctx, entClient, func(tx *ent.Tx) error {
			u, err := entClient.User.Create().SetUsername(userData.OpenID).SetRoleID(2).SetChannel("").Save(ctx)
			if err != nil {
				return err
			}
			_, err = entClient.Account.Create().SetOpenid(userData.OpenID).SetPassword("********").SetMethod("wechat").SetUserID(uint64(u.ID)).Save(ctx)
			if err != nil {
				return err
			}
			token, _ := jwt.GenerateToken(uint64(u.ID), u.RoleID)
			Result(http.StatusOK, gin.H{
				"access_token": token}, "成功！", ctx)
			return nil
		})
		if err != nil {
			FailWithMessage("账号创建失败", ctx)
			return
		}
	} else {
		// 账号存在查询用户信息
		u, err := entClient.User.Query().Where(user.IDEQ(int(ac.UserID))).First(ctx)
		if ent.IsNotFound(err) {
			_, err := entClient.User.Create().SetUsername(userData.OpenID).SetRoleID(2).SetChannel("").Save(ctx)
			if err != nil {
				FailWithMessage("用户创建失败", ctx)
				return
			}
			return
		} else {
			token, _ := jwt.GenerateToken(uint64(u.ID), u.RoleID)
			Result(http.StatusOK, gin.H{
				"access_token": token}, "成功！", ctx)
			return
		}

	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//origin := c.Request.Header.Get("Origin") //请求头部

		//接收客户端发送的origin （重要！）
		c.Header("Access-Control-Allow-Origin", "*")
		//服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允许跨域设置可以返回其他子段，可以自定义字段
		c.Header("Access-Control-Allow-Headers", "*")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		//设置缓存时间
		c.Header("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		c.Header("Access-Control-Allow-Credentials", "true")

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
