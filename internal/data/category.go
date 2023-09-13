package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/ent/category"
	"kubecit-service/internal/biz"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 用户数据仓库构造方法
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *categoryRepo) Delete(ctx context.Context, id int32) error {
	return c.data.db.Category.DeleteOneID(int(id)).Exec(ctx)
}

func (c *categoryRepo) Update(ctx context.Context, id int, name string) error {
	return c.data.db.Category.UpdateOneID(id).SetName(name).Exec(ctx)
}

func (c *categoryRepo) Create(ctx context.Context, category *biz.Category) error {
	create := c.data.db.Category.Create().SetName(category.CategoryName).SetLevel(category.Level)
	if category.ParentId != 0 {
		create.SetParentID(int(category.ParentId))
	}
	_, err := create.Save(ctx)
	return err
}

func (c *categoryRepo) ListAll(ctx context.Context) ([]*biz.Category, error) {
	categories, err := c.data.db.Category.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var categoryResult []*biz.Category

	for _, v := range categories {
		categoryResult = append(categoryResult, &biz.Category{
			CategoryName: v.Name,
			Id:           int32(v.ID),
			ParentId:     int32(v.ParentID),
			Level:        v.Level,
		})
	}
	return categoryResult, nil
}

func (c *categoryRepo) ListByLevel(ctx context.Context, level, categoryId *int32) ([]*biz.Category, error) {
	query := c.data.db.Category.Query()
	if level != nil {
		query.Where(category.Level(int(*level)))
	}

	if categoryId != nil {
		ca, err := c.data.db.Category.Query().Where(category.ID(int(*categoryId))).First(ctx)
		if err != nil {
			return nil, err
		}
		categories, err := c.data.db.Category.QueryChildren(ca).All(context.Background())
		if err != nil {
			return nil, err
		}

		var categoryResult []*biz.Category

		for _, v := range categories {
			categoryResult = append(categoryResult, &biz.Category{
				CategoryName: v.Name,
				Id:           int32(v.ID),
				ParentId:     int32(v.ParentID),
				Level:        v.Level,
			})
		}
		return categoryResult, nil
	}

	categories, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	var categoryResult []*biz.Category

	for _, v := range categories {
		categoryResult = append(categoryResult, &biz.Category{
			CategoryName: v.Name,
			Id:           int32(v.ID),
			ParentId:     int32(v.ParentID),
			Level:        v.Level,
		})
	}
	return categoryResult, nil

}

func (c *categoryRepo) ListSubCategories(ctx context.Context, parentID int32) ([]*biz.Category, error) {

	ca, err := c.data.db.Category.Query().Where(category.ID(int(parentID))).First(ctx)
	if err != nil {
		return nil, err
	}
	categories, err := c.data.db.Category.QueryChildren(ca).All(context.Background())
	if err != nil {
		return nil, err
	}
	var categoryResult []*biz.Category

	for _, v := range categories {
		categoryResult = append(categoryResult, &biz.Category{
			CategoryName: v.Name,
			Id:           int32(v.ID),
			ParentId:     int32(v.ParentID),
			Level:        v.Level,
		})
	}
	return categoryResult, nil
}
