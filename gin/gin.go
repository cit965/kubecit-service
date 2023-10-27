package gin

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/skip2/go-qrcode"
	"kubecit-service/ent"
	"kubecit-service/ent/account"
	"kubecit-service/ent/chapter"
	"kubecit-service/ent/course"
	"kubecit-service/ent/lesson"
	"kubecit-service/ent/user"
	"kubecit-service/internal/conf"
	"kubecit-service/internal/pkg/jwt"
	"kubecit-service/internal/pkg/provider/oss/qiniucloud"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

import "embed"

//go:embed all:dist
var StaticFiles embed.FS

type GinService struct {
	*gin.Engine
}

type WxConfig struct {
	Token     string
	AppId     string
	AppSecret string
}

type Datastore struct {
	Driver string
	Source string
}

type OssConfig struct {
	Bucket    string
	AccessKey string
	SecretKey string
	Domain    string
}

func initConfig(g *conf.Gin, d *conf.Data) (*WxConfig, *Datastore, *OssConfig) {
	return &WxConfig{
			Token:     g.Wechat.Token,
			AppId:     g.Wechat.Appid,
			AppSecret: g.Wechat.AppSecret,
		}, &Datastore{
			Driver: d.Database.Driver,
			Source: d.Database.Source,
		}, &OssConfig{
			Bucket:    g.Oss.Bucket,
			AccessKey: g.Oss.AccessKey,
			SecretKey: g.Oss.SecretKey,
			Domain:    g.Oss.Domain,
		}

}

var (
	WX  *WxConfig
	DB  *Datastore
	OSS *OssConfig
)

var (
	entClient *ent.Client
)

func initEntClient() *ent.Client {
	if entClient == nil {
		client, err := ent.Open(DB.Driver, DB.Source)
		if err != nil {
			log.Fatal("连接数据库失败")
		}
		entClient = client
	}
	return entClient
}

func NewGinService(g *conf.Gin, d *conf.Data) *GinService {
	WX, DB, OSS = initConfig(g, d)
	r := gin.Default()
	r.Use(Cors())
	// static files

	//staticFS, _ := fs.Sub(StaticFiles, "dist")
	//r.StaticFS("web/", http.FS(staticFS))
	// 签名校验
	r.GET("/gin/wechat/check", CheckSignature)
	//二维码生成
	r.GET("/gin/wechat/login", Redirect)
	// 回调地址
	r.GET("/gin/wechat/callback", Callback)

	// 上传文件
	r.POST("/gin/upload/:courseId/:chapterId/:lessonId", UploadHandler)
	return &GinService{r}
}

func CheckSignature(c *gin.Context) {
	s := fmt.Sprintf("token:%s, appid:%s, app_secret:%s", WX.Token, WX.AppId, WX.AppSecret)
	sprintf := fmt.Sprintf("driver:%s, source:%s", DB.Driver, DB.Source)
	fmt.Println(s, sprintf)

	// 获取查询参数中的签名、时间戳和随机数
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	// 创建包含令牌、时间戳和随机数的字符串切片
	tmpArr := []string{WX.Token, timestamp, nonce}
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
	wechatLoginURL := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&state=%s&scope=snsapi_userinfo#wechat_redirect", WX.AppId, redirectURL, state)
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
	tokenResp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", WX.AppId, WX.AppSecret, code))
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

	entClient = initEntClient()

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

func UploadHandler(c *gin.Context) {
	ossClient := qiniucloud.NewossQiniuCloud(OSS.Bucket, OSS.AccessKey, OSS.SecretKey)
	file, header, errFile := c.Request.FormFile("file")
	if errFile != nil {
		c.String(http.StatusBadRequest, "form file error")
		return
	}
	defer file.Close()

	courseId, _ := strconv.Atoi(c.Param("courseId"))
	chapterId, _ := strconv.Atoi(c.Param("chapterId"))
	lessonId, _ := strconv.Atoi(c.Param("lessonId"))

	filepath := fmt.Sprintf("%v/%v/%v/%v", c.Param("courseId"), c.Param("chapterId"), c.Param("lessonId"), header.Filename)

	_, err := ossClient.UploadFile(file, filepath, header.Size)
	if err != nil {
		c.String(http.StatusInternalServerError, "upload file failed")
		return
	}
	entClient = initEntClient()
	lessonInfo, err := entClient.Course.Query().Where(course.IDEQ(courseId)).
		QueryChapters().Where(chapter.IDEQ(chapterId)).
		QueryLessons().Where(lesson.IDEQ(lessonId)).Only(context.TODO())
	if err != nil {
		c.String(http.StatusInternalServerError, "query lesson error")
		return
	}
	_, err = lessonInfo.Update().SetStoragePath(filepath).Save(context.TODO())
	if err != nil {
		c.String(http.StatusInternalServerError, "update lesson error")
		return
	}
	c.String(http.StatusOK, "upload file %v success", filepath)
}
