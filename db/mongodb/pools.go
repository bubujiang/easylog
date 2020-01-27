package mongodb

import (
	"context"
	pools "github.com/jolestar/go-commons-pool/v2"
	"log-server/config"
	"strconv"
	"sync/atomic"
)

type MongoFactory struct {
}

func (f *MongoFactory) MakeObject(ctx context.Context) (*pools.PooledObject, error) {
	
	o := &Mongo{
		DSN:      config.Cnf.DB.DSN,
		Database: config.Cnf.DB.DSN,
	}

	err := o.Connect()
	if err != nil {
		return &pools.PooledObject{}, err
	}
	
	return pools.NewPooledObject(
			o.Operate),
		nil
}

func (f *MongoFactory) DestroyObject(ctx context.Context, object *pools.PooledObject) error {
	// do destroy
	return nil
}

func (f *MongoFactory) ValidateObject(ctx context.Context, object *pools.PooledObject) bool {
	// do validate
	return true
}

func (f *MongoFactory) ActivateObject(ctx context.Context, object *pools.PooledObject) error {
	// do activate
	return nil
}

func (f *MongoFactory) PassivateObject(ctx context.Context, object *pools.PooledObject) error {
	// do passivate
	return nil
}

func InitPool() *pools.ObjectPool {
	ctx := context.Background()
	p := pools.NewObjectPoolWithDefaultConfig(ctx, &MongoFactory{})
	p.Config.MaxTotal = 100

	return p
}
