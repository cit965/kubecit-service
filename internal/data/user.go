package data

import (
	"context"

	"kubecit-service/ent/user"
	"kubecit-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 用户数据仓库构造方法
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (repo *userRepo) FindById(ctx context.Context, id uint64) (po *biz.UserPO, err error) {
	user, err := repo.data.db.User.Query().Where(user.IDEQ(int(id))).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.UserPO{
		Id:       uint64(user.ID),
		Username: user.Username,
		Channel:  user.Channel,
		RoleId:   user.RoleID,
	}, nil

}

func (repo *userRepo) Save(ctx context.Context, userPO *biz.UserPO) error {
	if userPO.Id == 0 {
		nUser, err := repo.data.db.User.Create().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Save(ctx)
		if err != nil {
			return err
		}
		userPO.Id = uint64(nUser.ID)
		return nil
	} else {
		_, err := repo.data.db.User.Update().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Where(user.IDEQ(int(userPO.Id))).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}

}
