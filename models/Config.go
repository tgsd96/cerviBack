package models

type DBConfig struct {
	FirebaseConfig string
	Adapter        string
	DbName         string
	DbHost         string
	DbPass         string
	DbUser         string
}
type Config struct {
	APPName string `default:"cerviBack"`
	Version string
	Aws     struct {
		QueueUrl string
		Bucket   string
	}
	Dev  DBConfig
	Prod DBConfig
}
