package errs

import "github.com/go-kratos/kratos/v2/errors"

func BadError(err error, message string) *errors.Error {
	return errors.BadRequest(err.Error(), message)
}

func NotFound(err error, message string) *errors.Error {
	return errors.NotFound(err.Error(), message)
}

func OrderNotFound(err error) *errors.Error {
	return NotFound(err, "订单不存在！")
}

func NotEnoughMoney(err error) *errors.Error {
	return BadError(err, "余额不足,请及时充值！")
}
