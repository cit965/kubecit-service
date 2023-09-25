package common

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
)

func GetUserFromCtx(ctx context.Context) (userID uint64, err error) {
	userID, ok := ctx.Value("user_id").(uint64)
	if !ok {
		return 0, errors.New(400, "用户ID不存在", "从token中解析不出用户ID")
	}
	return userID, nil
}
