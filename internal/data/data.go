package data

import (
	"context"

	"kubecit-service/ent"
	"kubecit-service/internal/biz"
	"kubecit-service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCategoryRepo, NewSliderRepo, NewAccountRepo, NewUserRepo, NewUserAggregateRepo)

// Data .
type Data struct {
	conf *conf.Data
	db   *ent.Client
}

// NewData 构造方法，初始化了数据库 client
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	entClient, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("fail to open connection to db,%s", err)
	}
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("fail to create schema,%s", err)
	}
	return &Data{
		conf: c,
		db:   entClient,
	}, cleanup, nil
}

// example code
type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
