package adapter

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	NotFoundConfigMongo = "not found config mongo "
)

// Mongos ..
type Mongos map[string]*Mongo

// Get Mongo
func (adapters Mongos) Get(name string) (result *Mongo) {
	if adapter, ok := adapters[name]; ok {
		result = adapter
	} else {
		panic(NotFoundConfigMongo + name)
	}
	return
}

// Mongo struct
type Mongo struct {
	Name       string        `mapstructure:"name"`
	Address    []string      `mapstructure:"address"`
	RepSetName *string       `mapstructure:"repset_name"`
	DBAuthen   string        `mapstructure:"dbauthen"`
	User       string        `mapstructure:"user"`
	Pass       string        `mapstructure:"pass"`
	Timeout    time.Duration `mapstructure:"timeout"`
	IsSSL      bool          `mapstructure:"is_ssl"`
	DBName     string        `mapstructure:"dbname"`
	PoolLimit  *uint64       `mapstructure:"pool_limit"`
	ReadPref   string        `mapstructure:"read_pref"`
	ConClient  *mongo.Client
}

var (
	onceMongo      map[string]*sync.Once
	onceMongoMutex = sync.RWMutex{}
)

func init() {
	onceMongo = make(map[string]*sync.Once)
}

func mapReadPrefV2(readPreference string) (readPrefMode *readpref.ReadPref) {

	mode, err := readpref.ModeFromString(readPreference)
	if err != nil {
		mode = readpref.PrimaryMode
	}

	readPrefMode, _ = readpref.New(mode)
	return
}

// Init func
func (config *Mongo) Init() {

	// ! loop until connect db success
	for {
		initProcessErr := func() error {
			onceMongoMutex.Lock()
			defer onceMongoMutex.Unlock()

			// each name key init by onetime by singletone
			if onceMongo[config.Name] == nil {
				onceMongo[config.Name] = &sync.Once{}
			}

			var connectError error

			onceMongo[config.Name].Do(func() {
				log.Printf("[%s][%s] mongo [connecting]\n", config.Name, config.Address)

				// create client option
				socketTimeOut := config.Timeout * time.Second
				clientOption := &options.ClientOptions{
					Hosts:                  config.Address,
					ReplicaSet:             config.RepSetName,
					MaxPoolSize:            config.PoolLimit,
					SocketTimeout:          &socketTimeOut,
					ServerSelectionTimeout: &socketTimeOut,
				}

				// set Authen
				if config.User != "" && config.Pass != "" {
					clientAuth := options.Credential{
						AuthSource: config.DBAuthen,
						Username:   config.User,
						Password:   config.Pass,
					}

					clientOption.SetAuth(clientAuth)
				}

				// set readRef
				clientOption.SetReadPreference(mapReadPrefV2(config.ReadPref))

				// set SSL
				if config.IsSSL {
					tlsConfig := new(tls.Config)
					clientOption.SetTLSConfig(tlsConfig)
				}

				// connect to DB
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				myConClient, err := mongo.Connect(ctx, clientOption)
				if err != nil {
					connectError = err
					return
				}

				// Check the connection
				errCheckConnect := myConClient.Ping(ctx, nil)
				if errCheckConnect != nil {
					connectError = errCheckConnect
				} else {
					// don't disconnect client for using connection pool (like: using session.copy in mgo driver)
					config.ConClient = myConClient

					log.Printf("[%s][%s] mongo [connected]\n", config.Name, config.Address)
				}
			})

			return connectError
		}()

		if initProcessErr == nil {
			break
		} else {
			log.Printf("[%s][%s] mongo [error]: %v \n", config.Name, config.Address, initProcessErr)
			onceMongo[config.Name] = &sync.Once{}
			time.Sleep(1 * time.Second)
		}
	}
}

func (config *Mongo) GetClient() *mongo.Client {
	return config.ConClient
}

//CollectionV2 ..
type CollectionV2 struct {
	*mongo.Collection
	Mongo Mongo
}

// GetCollectionV2 func
func (config Mongo) GetCollectionV2(collectionName string, dbName ...string) (collection CollectionV2) {
	if config.ConClient != nil {
		db := config.DBName

		if len(dbName) == 1 {
			db = dbName[0]
		}

		collection = CollectionV2{config.ConClient.Database(db).Collection(collectionName), config}
	} else {
		panic(fmt.Errorf("[%s] Chưa init Mongo", config.Name))
	}
	return
}

// GetChangeStream func
func (config Mongo) GetChangeStream(dbName string, collectionName string, cb func(bson.M)) (err error) {

	ctx := context.Background()
	cur := &mongo.ChangeStream{}
	opts := options.ChangeStream()
	pinelineData := []bson.D{}

	// Set up type look up
	opts.SetFullDocument(options.UpdateLookup)

	if dbName != "" && collectionName != "" { // Watching a collection
		fmt.Println("Watching", dbName+"."+collectionName)

		coll := config.ConClient.Database(dbName).Collection(collectionName)
		cur, err = coll.Watch(ctx, pinelineData, opts)

	} else if dbName != "" { // Watching a database

		fmt.Println("Watching", dbName)
		db := config.ConClient.Database(dbName)
		cur, err = db.Watch(ctx, pinelineData, opts)

	} else { // Watching all

		fmt.Println("Watching all")
		cur, err = config.ConClient.Watch(ctx, pinelineData, opts)
	}

	if err != nil {
		return
	}

	defer cur.Close(ctx)

	// loop forever look change data
	for cur.Next(ctx) {
		data := bson.M{}
		cur.Decode(&data)
		cb(data)
	}

	return
}

// NextID ..
func (config Mongo) NextID(collectionName string, dbName ...string) (id int) {

	type myCounter struct {
		ID  string `json:"_id"`
		Seq int    `json:"seq"`
	}

	counter := myCounter{}

	if config.ConClient != nil {
		db := config.DBName

		if len(dbName) == 1 {
			db = dbName[0]
		}

		collection := config.ConClient.Database(db).Collection("counters")

		err := collection.FindOneAndUpdate(context.Background(),
			bson.M{"_id": collectionName},
			bson.M{"$set": bson.M{"_id": collectionName}, "$inc": bson.M{"seq": 1}},
			options.FindOneAndUpdate().SetUpsert(true)).Decode(&counter)

		if err == mongo.ErrNoDocuments { // nếu như collection mới
			err = nil
			id = 1
			return
		}

		if err != nil {
			panic(err)
		}

		id = counter.Seq + 1
	} else {
		panic(fmt.Errorf("[%s] Chưa init Mongo", config.Name))
	}

	return
}
