package mongodb

import (
	"context"
	pools "github.com/jolestar/go-commons-pool/v2"
	"log-server/config"
	"sync"
)
var (
	initialized sync.Once
	P *pools.ObjectPool
)

type MongoFactory struct {
}

func (f *MongoFactory) MakeObject(ctx context.Context) (*pools.PooledObject, error) {
	
	o := &Mongo{
		DSN:      config.Cnf.DB.DSN,
		Database: config.Cnf.DB.Database,
	}

	err := o.Connect()
	if err != nil {
		return &pools.PooledObject{}, err
	}
	
	return pools.NewPooledObject(
			o),
		nil
}

func (f *MongoFactory) DestroyObject(ctx context.Context, object *pools.PooledObject) error {
	// do destroy
	o := object.Object.(*Mongo)
	o.Close()
	return nil
}

func (f *MongoFactory) ValidateObject(ctx context.Context, object *pools.PooledObject) bool {
	// do validate
	o := object.Object.(*Mongo)
	return o.Validate()
}

func (f *MongoFactory) ActivateObject(ctx context.Context, object *pools.PooledObject) error {
	// do activate
	return nil
}

func (f *MongoFactory) PassivateObject(ctx context.Context, object *pools.PooledObject) error {
	// do passivate
	return nil
}

func Init() {
	initialized.Do(func() {
		ctx := context.Background()
		P = pools.NewObjectPoolWithDefaultConfig(ctx, &MongoFactory{})
		P.Config.MaxTotal = config.Cnf.DB.Max
	})
}
