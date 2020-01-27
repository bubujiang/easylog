package mongodb

import (
	"context"
	pool "github.com/jolestar/go-commons-pool/v2"
	"strconv"
	"sync/atomic"
)

type MongoFactory struct {
}

func (f *MongoFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	return pool.NewPooledObject(
			&MyPoolObject{
				s: strconv.FormatUint(atomic.AddUint64(&f.v, 1), 10),
			}),
		nil
}

func (f *MongoFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	// do destroy
	return nil
}

func (f *MongoFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	// do validate
	return true
}

func (f *MongoFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do activate
	return nil
}

func (f *MongoFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do passivate
	return nil
}
