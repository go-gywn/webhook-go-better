package helpers

import (
	"context"
	"gorm.io/gorm"
	"sync"
)

const ContextDBKey = "DB"
const ContextValidTokenKey = "validToken"

type contextHelper struct {
}

var (
	contextHelperOnce     sync.Once
	contextHelperInstance *contextHelper
)

func ContextHelper() *contextHelper {
	contextHelperOnce.Do(func() {
		contextHelperInstance = &contextHelper{}
	})

	return contextHelperInstance
}

func (contextHelper) GetDB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ContextDBKey)
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*gorm.DB); ok {
		return db
	}
	panic("DB is not exist")
}

func (contextHelper) SetDB(ctx context.Context, gormDB *gorm.DB) context.Context {
	return context.WithValue(ctx, ContextDBKey, gormDB)
}

func (contextHelper) SetValidToken(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextValidTokenKey, true)
}

func (contextHelper) IsValidToken(ctx context.Context) bool {
	v := ctx.Value(ContextValidTokenKey)
	if v == nil{
		return false
	}
	return ctx.Value(ContextValidTokenKey).(bool)
}