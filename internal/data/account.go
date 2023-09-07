package data

import (
	"context"

	"kubecit-service/ent/account"
	"kubecit-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo 用户账号数据仓库构造方法
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (repo *accountRepo) FindByOpenidAndMethod(ctx context.Context, openid string, method string) (po *biz.AccountPO, err error) {
	account, err := repo.data.db.Account.Query().Where(account.And(account.OpenidEQ(openid), account.MethodEQ(method))).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.AccountPO{
		Id:       uint64(account.ID),
		UserId:   account.UserID,
		Openid:   account.Openid,
		Password: account.Password,
		Method:   account.Method,
	}, nil

}

func (repo *accountRepo) Save(ctx context.Context, accountPO *biz.AccountPO) error {

	if accountPO.Id == 0 {
		_, err := repo.data.db.Account.Create().
			SetOpenid(accountPO.Openid).
			SetPassword(accountPO.Password).
			SetUserID(accountPO.UserId).
			SetMethod(accountPO.Method).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err := repo.data.db.Account.Update().
			SetOpenid(accountPO.Openid).
			SetPassword(accountPO.Password).
			SetUserID(accountPO.UserId).
			SetMethod(accountPO.Method).
			Where(account.IDEQ(int(accountPO.Id))).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}

}
