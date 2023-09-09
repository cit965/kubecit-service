package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kubecit-service/ent"
	"kubecit-service/ent/slider"
	"kubecit-service/internal/biz"
)

type sliderRepo struct {
	data *Data
	log  *log.Helper
}

// NewSliderRepo 用户数据仓库构造方法
func NewSliderRepo(data *Data, logger log.Logger) biz.SliderRepo {
	return &sliderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *sliderRepo) Create(ctx context.Context, slider *biz.Slider) (*biz.Slider, error) {
	res, err := c.data.db.Slider.Create().SetTitle(slider.Title).SetContent(slider.Content).SetImageLink(slider.ImageLink).
		SetIsValid(slider.IsValid).SetPriority(int(slider.Priority)).Save(ctx)
	if err != nil {
		c.log.Errorf("slider repo create error: %v\n", err)
		return nil, err
	}
	return &biz.Slider{
		Id:        res.ID,
		Title:     res.Title,
		Content:   res.Content,
		ImageLink: res.ImageLink,
		CreateAt:  res.CreateAt,
		UpdateAt:  res.UpdateAt,
		IsValid:   res.IsValid,
		Priority:  res.Priority,
	}, nil
}

func (c *sliderRepo) Get(ctx context.Context, id int) (*biz.Slider, error) {
	res, err := c.data.db.Slider.Query().Where(slider.IDEQ(id)).Only(ctx)
	if err != nil {
		c.log.Errorf("slider repo get error: %v\n", err)
		return nil, err
	}
	return &biz.Slider{
		Id:        res.ID,
		Title:     res.Title,
		Content:   res.Content,
		ImageLink: res.ImageLink,
		CreateAt:  res.CreateAt,
		UpdateAt:  res.UpdateAt,
		IsValid:   res.IsValid,
		Priority:  res.Priority,
	}, nil
}

func (c *sliderRepo) List(ctx context.Context) ([]*biz.Slider, error) {
	res, err := c.data.db.Slider.Query().All(ctx)
	if err != nil {
		c.log.Errorf("slider repo list error: %v\n", err)
		return nil, err
	}
	sliders := make([]*biz.Slider, 0)
	for _, value := range res {
		sliderIns := &biz.Slider{
			Id:        value.ID,
			Title:     value.Title,
			Content:   value.Content,
			ImageLink: value.ImageLink,
			CreateAt:  value.CreateAt,
			UpdateAt:  value.UpdateAt,
			IsValid:   value.IsValid,
			Priority:  value.Priority,
		}
		sliders = append(sliders, sliderIns)
	}
	return sliders, nil
}

func (c *sliderRepo) Delete(ctx context.Context, id int) (int, error) {
	res, err := c.data.db.Slider.Delete().Where(slider.IDEQ(id)).Exec(ctx)
	if err != nil {
		c.log.Errorf("slider repo delete error: %v\n", err)
		return -1, err
	}
	return res, nil
}

func (c *sliderRepo) Update(ctx context.Context, id int, ins *biz.Slider) (*biz.Slider, error) {
	res, err := c.data.db.Slider.UpdateOneID(id).SetTitle(ins.Title).SetContent(ins.Content).
		SetImageLink(ins.ImageLink).SetIsValid(ins.IsValid).SetPriority(ins.Priority).Save(ctx)
	if err != nil {
		c.log.Errorf("slider repo update error: %v\n", err)
		return nil, err
	}

	return &biz.Slider{
		Id:        res.ID,
		Title:     res.Title,
		Content:   res.Content,
		ImageLink: res.ImageLink,
		CreateAt:  res.CreateAt,
		UpdateAt:  res.UpdateAt,
		IsValid:   res.IsValid,
		Priority:  res.Priority,
	}, nil
}

func (c *sliderRepo) ListByPriority(ctx context.Context, count int) ([]*biz.Slider, error) {
	res, err := c.data.db.Slider.Query().Order(ent.Asc(slider.FieldPriority)).
		Limit(count).All(ctx)
	if err != nil {
		c.log.Errorf("slider repo list by priority error: %v\n", err)
		return nil, err
	}
	sliders := make([]*biz.Slider, 0)
	for _, value := range res {
		sliderIns := &biz.Slider{
			Id:        value.ID,
			Title:     value.Title,
			Content:   value.Content,
			ImageLink: value.ImageLink,
			CreateAt:  value.CreateAt,
			UpdateAt:  value.UpdateAt,
			IsValid:   value.IsValid,
			Priority:  value.Priority,
		}
		sliders = append(sliders, sliderIns)
	}
	return sliders, nil
}
