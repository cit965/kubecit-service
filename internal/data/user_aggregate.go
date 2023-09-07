package data

import (
	"context"

	"kubecit-service/ent/account"
	"kubecit-service/ent/user"
	"kubecit-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userAggregateRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserAggregateRepo 用户账号聚合数据仓库构造方法
func NewUserAggregateRepo(data *Data, logger log.Logger) biz.UserAggregateRepo {
	return &userAggregateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userAggregateRepo) SaveAccountAndUser(ctx context.Context, accountPO *biz.AccountPO, userPO *biz.UserPO) error {
	var err error
	tx, err := repo.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if userPO.Id == 0 {
		nUser, err := tx.User.Create().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Save(ctx)
		if err != nil {
			return err
		}
		userPO.Id = uint64(nUser.ID)
	} else {
		_, err := tx.User.Update().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Where(user.IDEQ(int(userPO.Id))).
			Save(ctx)
		if err != nil {
			return err
		}

	}
	if accountPO.Id == 0 {
		_, err := tx.Account.Create().
			SetOpenid(accountPO.Openid).
			SetPassword(accountPO.Password).
			SetUserID(userPO.Id).
			SetMethod(accountPO.Method).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := tx.Account.Update().
			SetOpenid(accountPO.Openid).
			SetPassword(accountPO.Password).
			SetUserID(accountPO.UserId).
			SetMethod(accountPO.Method).
			Where(account.IDEQ(int(accountPO.Id))).
			Save(ctx)
		if err != nil {
			return err
		}

	}
	return tx.Commit()
}
