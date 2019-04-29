package settings

//Settings ...
type Settings struct {
	MongoConnectionString string `env:"MONGO_CONNECTION_STRING,required"`
}
