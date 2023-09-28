package common

import (
	"errors"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func Copy(a interface{}, b interface{}) (err error) {
	return copier.CopyWithOption(a, b, copier.Option{IgnoreEmpty: false})
}

func DeepCopy(a interface{}, b interface{}) (err error) {
	return copier.CopyWithOption(a, b, copier.Option{IgnoreEmpty: false, DeepCopy: true})
}

func CopyConvertType(a interface{}, b interface{}) (err error) {
	return copier.CopyWithOption(a, b, copier.Option{
		IgnoreEmpty: false,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: &timestamppb.Timestamp{},
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(time.Time)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return timestamppb.New(s), nil
				},
			},
		}})
}

func DeepCopyConvertType(a interface{}, b interface{}) (err error) {
	return copier.CopyWithOption(a, b, copier.Option{
		DeepCopy:    true,
		IgnoreEmpty: false,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: &timestamppb.Timestamp{},
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(time.Time)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return timestamppb.New(s), nil
				},
			},
		}})
}
