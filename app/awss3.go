package main

import (
	"fmt"
	"os"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"
)


func main() {

	var (
		file *os.File
		err error
		awsS3 *infrastructure.AwsS3
		url string
	)

	// このファイルと同じ階層に images/sample.png がある想定
	file, err = os.Open("images/sample.png")

	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer file.Close()

	awsS3 = infrastructure.NewAwsS3()

	/**
	 * ex) gin でのファイル操作
	 *
	 * package gin
	 * type Context
	 * func (c *gin.Context) FormFile(name string) (*multipart.FileHeader, error)
	 *
	 *
	 * import "mime/multipart"
	 * type FileHeader
	 * func (fh *FileHeader) Open() (multipart.File, error)
	 *
	 *
	 * fileHeader, err := context.FormFile("uploadImage")
	 * if err != nil { ... }
	 *
	 * file, err := fileHeader.Open()
	 * if err != nil { ... }
	 *
	 * url, err = awsS3.UploadTest(file, "test", "png")
	 */

	// multipart.File と os.File は同じように扱える
	url, err = awsS3.UploadTest(file, "test", "png")

	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(url)
}
