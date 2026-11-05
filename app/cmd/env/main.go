
package main


import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)



type SampleENV struct {
	Test string `required:"false"`
	EchoHello string `required:"false"`
	ErrorTest string `required:"false"`
	Hoge string `required:"true"`
	CamelCase string `required:"true"`
	SplitUnderBar string `required:"true" split_words:"true"`
	SomethingID string `required:"true" split_words:"true"`
	Number123 string `required:"true" split_words:"true"`
	Number456 string `envconfig:"number_456" required:"true" split_words:"true"`
	Number78 string `envconfig:"number7_8" required:"true" split_words:"true"`
	URLTest string `required:"true" split_words:"true"`
	DiffText string `envconfig:"this_is_text" required:"true" split_words:"true"`
}


// 環境変数の出力
//
func getenv(key string) {
	// if os.Getenv(key) == "" {
	// 	panic(key + " is empty.\n")
	// }
	fmt.Printf("[%s] %s\n", key, os.Getenv(key))
}

func main() {
	// echo $PATH
	getenv("PATH")

	// echo $HOME
	getenv("HOME")

	// export TEST=test123
	// echo $TEST
	getenv("TEST")

	// "github.com/joho/godotenv"
	//
	// go run main.go した階層の .env を読み込む。
	// go build main.go はまだ試してない。
	//
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file\n")
	}
	// if err := godotenv.Load("test.env"); err != nil {
	// 	panic("Error loading test.env file\n")
	// }

	// .env
	getenv("TEST_VALUE")


	// "github.com/kelseyhightower/envconfig"
	//
	// export SAMPLE_TEST=test
	// export SAMPLE_ECHOHELLO=World
	//
	// export SAMPLE_ERRORTEST=error/err:hoge
	//
	// or
	//
	// source .envconfig
	//
	var s SampleENV
	if err := envconfig.Process("sample", &s); err != nil {
		panic(err.Error())
	}

	fmt.Println(s)
	fmt.Print(s.Hoge, "\n")
	fmt.Print(s.CamelCase, "\n")
	fmt.Print(s.SplitUnderBar, "\n")
	fmt.Print(s.SomethingID, "\n")
	fmt.Print(s.Number123, "\n")
	fmt.Print(s.Number456, "\n")
	fmt.Print(s.Number78, "\n")
	fmt.Print(s.URLTest, "\n")
	fmt.Print(s.DiffText, "\n")


	// getenv("TEST")
	// getenv("ECHO_HELLO")
	// getenv("ERROR_TEST")
}
