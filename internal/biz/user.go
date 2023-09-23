package biz

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"strconv"
	"sync"
	"time"

	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/ent"
	"kubecit-service/internal/pkg/jwt"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	AccountMethodUsername = "username"
	AccountMethodWeChat   = "wechat"
	AccountMethodEmail    = "email"
)
const (
	UserRoleInvalid uint8 = iota
	UserRoleGuest
	UserRoleRegisterUser
	UserRoleLecturer
	UserRoleSuperAdmin
)
const (
	ServiceUserErrorCode = 100000
)
const (
	TypeSystemErrorCode = 1000 + iota*1000
	TypeDatabaseErrorCode
	TypeUserParamErrorCode
	UserUsernameOrEmailExistsErrorCode
)

// ServiceUserErrorCode+TypeSystemErrorCode

const (
	_ = ServiceUserErrorCode + TypeSystemErrorCode + iota
)

// ServiceUserErrorCode+TypeDatabaseErrorCode

const (
	_ = ServiceUserErrorCode + TypeDatabaseErrorCode + iota
	UserSaveDatabaseErrorCode
	UserFindDatabaseErrorCode
	InvalidEmailVerificationCodeErrorCode
)

// ServiceUserErrorCode+TypeUserParamErrorCode

const (
	_ = ServiceUserErrorCode + TypeUserParamErrorCode + iota
	UserUsernameIsExists
	UserUsernameNotExists
	UserPasswordNotMatch
)

var emailVerificationCodes sync.Map

type AccountPO struct {
	Id       uint64
	UserId   uint64
	Openid   string
	Password string
	Method   string
}

type UserPO struct {
	Id       uint64
	Username string
	Channel  string
	RoleId   uint8
}

type AccountRepo interface {
	FindByOpenidAndMethod(ctx context.Context, openid string, method string) (po *AccountPO, err error)
	Save(ctx context.Context, accountPO *AccountPO) error
}

type UserRepo interface {
	FindById(ctx context.Context, id uint64) (po *UserPO, err error)
	Save(ctx context.Context, po *UserPO) error
	SaveAccountAndUserTx(ctx context.Context, accountPO *AccountPO, userPO *UserPO) error
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	accountRepo AccountRepo
	userRepo    UserRepo
	log         *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(accountRepo AccountRepo, userRepo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		accountRepo: accountRepo,
		userRepo:    userRepo,
		log:         log.NewHelper(logger)}
}

func (usecase *UserUsecase) LoginByJson(ctx context.Context, request *pb.LoginByJsonRequest) (*pb.LoginByJsonReply, error) {

	accountPO, err := usecase.accountRepo.FindByOpenidAndMethod(ctx, request.Username, AccountMethodUsername)
	if accountPO == nil {
		if _, isEmpty := err.(*ent.NotFoundError); isEmpty {
			return &pb.LoginByJsonReply{
				Meta: usecase.errorMeta("用户名不存在", UserUsernameNotExists),
			}, nil
		} else {
			return &pb.LoginByJsonReply{
				Meta: usecase.errorMeta("数据库发生错误", UserFindDatabaseErrorCode),
			}, nil
		}
	}
	if accountPO.Password != usecase.md5(request.Password) {
		return &pb.LoginByJsonReply{
			Meta: usecase.errorMeta("密码错误", UserPasswordNotMatch),
		}, nil

	}

	userPO, err := usecase.userRepo.FindById(ctx, accountPO.UserId)
	if err != nil {
		return &pb.LoginByJsonReply{
			Meta: usecase.errorMeta("数据库发生错误", UserFindDatabaseErrorCode),
		}, nil
	}

	token, _ := jwt.GenerateToken(userPO.Id, userPO.RoleId)
	return &pb.LoginByJsonReply{
		Meta: &pb.Metadata{
			Code:    "0",
			Success: true,
		},
		Data: &pb.LoginByJsonReplyData{AccessToken: token},
	}, nil

}
func (usecase *UserUsecase) RegisterUsername(ctx context.Context, request *pb.RegisterUsernameRequest) (*pb.RegisterUsernameReply, error) {
	// 验证邮箱验证码
	storedCode, ok := emailVerificationCodes.Load(request.Email)
	if !ok || storedCode != request.EmailVerificationCode {
		return &pb.RegisterUsernameReply{
			Meta: usecase.errorMeta("无效的邮箱验证码", InvalidEmailVerificationCodeErrorCode),
		}, nil
	}
	// 检查用户名是否已存在
	accountByUsername, err := usecase.accountRepo.FindByOpenidAndMethod(ctx, request.Username, AccountMethodUsername)
	if err != nil {
		if _, isEmpty := err.(*ent.NotFoundError); !isEmpty {
			usecase.log.Errorf("注册用户名错误：%v", err.Error())
			return &pb.RegisterUsernameReply{
				Meta: usecase.errorMeta("数据库发生错误", UserFindDatabaseErrorCode),
			}, nil
		}
	}

	// 检查电子邮件是否已存在
	accountByEmail, err := usecase.accountRepo.FindByOpenidAndMethod(ctx, request.Email, AccountMethodEmail)
	if err != nil {
		if _, isEmpty := err.(*ent.NotFoundError); !isEmpty {
			usecase.log.Errorf("注册电子邮件错误：%v", err.Error())
			return &pb.RegisterUsernameReply{
				Meta: usecase.errorMeta("数据库发生错误", UserFindDatabaseErrorCode),
			}, nil
		}
	}

	// 如果用户名或电子邮件已存在，返回错误
	if accountByUsername != nil || accountByEmail != nil {
		return &pb.RegisterUsernameReply{
			Meta: usecase.errorMeta("用户名或电子邮件已存在", UserUsernameOrEmailExistsErrorCode),
		}, nil
	}

	// 创建用户账户和信息
	userPO := &UserPO{
		Username: request.Username,
		Channel:  "",
		RoleId:   UserRoleRegisterUser,
	}
	accountPO := &AccountPO{
		Openid:   request.Email, // 使用电子邮件作为 OpenID
		Password: usecase.md5(request.Password),
		Method:   AccountMethodEmail, // 使用电子邮件注册
	}
	err = usecase.userRepo.SaveAccountAndUserTx(ctx, accountPO, userPO)
	if err != nil {
		usecase.log.Errorf("register username err: %v", err.Error())
		return &pb.RegisterUsernameReply{
			Meta: usecase.errorMeta("数据库发生错误", UserSaveDatabaseErrorCode),
		}, nil
	}

	token, _ := jwt.GenerateToken(userPO.Id, userPO.RoleId)
	return &pb.RegisterUsernameReply{
		Meta: &pb.Metadata{
			Code:    "0",
			Success: true,
		},
		Data: &pb.LoginByJsonReplyData{AccessToken: token},
	}, nil
}

func (usecase *UserUsecase) SendEmailVerificationCode(ctx context.Context, request *pb.SendEmailVerificationCodeRequest) (*pb.SendEmailVerificationCodeReply, error) {
	// 生成随机的验证码，这里假设你已经有一个生成验证码的函数
	verificationCode, err := generateVerificationCode(4)
	if err != nil {
		return nil, err
	}

	// 发送电子邮件
	err = sendEmail(request.Email, verificationCode)
	if err != nil {
		return nil, err
	}

	// 如果需要，可以在这里记录已发送验证码的信息
	emailVerificationCodes.Store(request.Email, verificationCode)

	// 返回响应
	return &pb.SendEmailVerificationCodeReply{
		// 可以添加其他响应字段，根据需要
	}, nil
}

func (usecase *UserUsecase) md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str

}
func (usecase *UserUsecase) errorMeta(msg string, code int) *pb.Metadata {
	return &pb.Metadata{
		Msg:       msg,
		Code:      strconv.FormatInt(int64(code), 10),
		Success:   false,
		Version:   "",
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
	}

}

func (usecase *UserUsecase) CurrentUserInfo(ctx context.Context) (*pb.UserInfoReply, error) {
	UserId := ctx.Value("user_id").(uint64)
	userPO, err := usecase.userRepo.FindById(ctx, UserId)
	if err != nil {
		return &pb.UserInfoReply{}, err
	}
	return &pb.UserInfoReply{
		Username: userPO.Username,
		Channel:  userPO.Channel,
		RoleId:   uint32(int32(userPO.RoleId)),
	}, nil
}

// 生成随机验证码
func generateVerificationCode(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes)[:length], nil
}

// 发送包含验证码的电子邮件
func sendEmail(email, code string) error {

	from := "your_email@example.com"
	password := "your_email_password"
	to := []string{email}
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "注册验证码"
	message := "您的验证码是：" + code

	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + message)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
