package gin

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"sort"
)

import "embed"

//go:embed all:dist
var StaticFiles embed.FS

type GinService struct {
	*gin.Engine
}

const TOKEN = "111"

func NewGinService() *GinService {
	r := gin.Default()
	r.Use(Cors())
	// static files

	staticFS, _ := fs.Sub(StaticFiles, "dist")
	r.StaticFS("web/", http.FS(staticFS))
	r.GET("wechat/check", CheckSignature)
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
