package main

//Settings ...
type Settings struct {
	MongoConnectionString string `env:"MONGO_CONNECTION_STRIG,required"`
}
