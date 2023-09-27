package biz

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/errors"
	"strconv"
	"time"

	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

// VideoPlayerUsecase is a VideoPlayer usecase.
type VideoPlayerUsecase struct {
	conf       *conf.TencentCloudVideoPlayer
	courseRepo CourseRepo
	log        *log.Helper
}

// NewVideoPlayerUsecase new a VideoPlayer usecase.
func NewVideoPlayerUsecase(conf *conf.TencentCloudVideoPlayer, courseRepo CourseRepo, logger log.Logger) *VideoPlayerUsecase {
	return &VideoPlayerUsecase{conf: conf, courseRepo: courseRepo, log: log.NewHelper(logger)}
}
func (usecase *VideoPlayerUsecase) GetPlayerParam(ctx context.Context, req *pb.VideoPlayerGetParamReq) (*pb.VideoPlayerGetParamReply, error) {
	userRow := ctx.Value("user_id")
	userId, ok := userRow.(uint64)

	if !ok {
		return nil, errors.BadRequest("解析不出用户ID", "用户未登录")
	}

	lesson, err := usecase.courseRepo.GetLesson(ctx, int(req.LessonId))
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, errors.BadRequest("课程不存在", "课程不存在")
	}

	// todo 判断用户是否有这节课的权限
	if lesson.StoragePath == "" {
		return nil, errors.BadRequest("音视频不存在", "课程小结中没有音视频资源")
	}

	sign := usecase.buildSign(ctx, lesson.StoragePath, userId)

	return &pb.VideoPlayerGetParamReply{
		Appid:   uint64(usecase.conf.GetAppid()),
		FieldID: lesson.StoragePath,
		PSign:   sign,
	}, nil
}
func (usecase *VideoPlayerUsecase) buildSign(ctx context.Context, fileID string, userID uint64) (sign string) {
	appId := usecase.conf.GetAppid() //  appid
	fileId := fileID                 // 目标 FileId
	audioVideoType := "Original"     // 播放的音视频类型 原视频

	currentTime := time.Now().Unix()
	psignExpire := currentTime + 3600
	urlTimeExpire := strconv.FormatInt(psignExpire, 16) // 可任意设置过期时间，16进制字符串形式，示例1h
	playKey := []byte(usecase.conf.GetAppSecret())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"appId":  appId,
		"fileId": fileId,
		"contentInfo": jwt.MapClaims{
			"audioVideoType": audioVideoType,
		},
		"currentTimeStamp": currentTime,
		"expireTimeStamp":  psignExpire,
		"urlAccessInfo": map[string]string{
			"t":      urlTimeExpire,
			"uv":     strconv.FormatUint(userID, 16), // 水印
			"domain": usecase.conf.Domain,
			"scheme": usecase.conf.Scheme,
		},
	})

	tokenString, _ := token.SignedString(playKey)
	return tokenString

}
