
package infrastructure

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// type Config struct {
// 	Aws struct {
// 		S3 struct {
// 			Region string
// 			Bucket string
// 			AccessKeyID string
// 			SecretAccessKey string
// 		}
// 	}
// 	DB struct {
// 		Production struct {
// 			Host string
// 			Username string
// 			Password string
// 			DBName string
// 		}
// 	}
// 	Slack struct {
// 		APIToken string
// 	}
// 	Mail struct {
// 		Gmail struct {
// 			Addr string
// 			Host string
// 			Address string
// 			Name string
// 			Password string
// 			To string
// 		}
// 	}
// }


type Config struct {
	AwsS3Region string `required:"true" split_words:"true"`
	AwsS3Bucket string `required:"true" split_words:"true"`
	AwsS3AccessKeyID string `required:"true" split_words:"true"`
	AwsS3SecretAccessKey string `required:"true" split_words:"true"`
	DBProductionHost string `required:"true" split_words:"true"`
	DBProductionUsername string `required:"true" split_words:"true"`
	DBProductionPassword string `required:"true" split_words:"true"`
	DBProductionDBName string `required:"true" split_words:"true"`
	SlackAPIToken string `required:"true" split_words:"true"`
	MailGmailAddr string `required:"true" split_words:"true"`
	MailGmailHost string `required:"true" split_words:"true"`
	MailGmailAddress string `required:"true" split_words:"true"`
	MailGmailName string `required:"true" split_words:"true"`
	MailGmailPassword string `required:"true" split_words:"true"`
	MailGmailTo string `required:"true" split_words:"true"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file\n")
	}
	c := &Config{}
	if err := envconfig.Process("gosample", c); err != nil {
		panic(err.Error())
	}
	return c
}
