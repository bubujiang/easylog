package db

import (
	"context"
	pool "github.com/jolestar/go-commons-pool/v2"
)

type Factory interface {
	MakeObject(ctx context.Context) (*pool.PooledObject, error)
	DestroyObject(ctx context.Context, object *pool.PooledObject) error
	ValidateObject(ctx context.Context, object *pool.PooledObject) bool
	ActivateObject(ctx context.Context, object *pool.PooledObject) error
	PassivateObject(ctx context.Context, object *pool.PooledObject) error
}

