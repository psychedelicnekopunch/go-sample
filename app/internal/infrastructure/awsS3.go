
package infrastructure


import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


type AwsS3 struct {
	Config *Config
	Keys AwsS3URLs
	Uploader *s3manager.Uploader
}


type AwsS3URLs struct {
	Test string
}

/**
 * package: session
 *
 * func Must(sess *Session, err error) *Session
 *
 * ex)
 * sess := session.Must(session.NewSessionWithOptions(session.Options{}))
 *
 * func NewSessionWithOptions(opts session.Options) (*Session, error)
 *
 * type Options struct {
 *   Config aws.Config
 *   Profile string
 *   SharedConfigState SharedConfigState
 *   SharedConfigFiles []string
 *   AssumeRoleTokenProvider func() (string, error)
 *   CustomCABundle io.Reader
 * }
 *
 */

/**
 * package: aws
 *
 * type Config struct {
 *   CredentialsChainVerboseErrors *bool
 *   Credentials *credentials.Credentials
 *   Endpoint *string
 *   EndpointResolver endpoints.Resolver
 *   EnforceShouldRetryCheck *bool
 *   Region *string
 *   DisableSSL *bool
 *   HTTPClient *http.Client
 *   LogLevel *LogLevelType
 *   Logger Logger
 *   MaxRetries *int
 *   Retryer RequestRetryer
 *   DisableParamValidation *bool
 *   DisableComputeChecksums *bool
 *   S3ForcePathStyle *bool
 *   S3Disable100Continue *bool
 *   S3UseAccelerate *bool
 *   S3DisableContentMD5Validation *bool
 *   EC2MetadataDisableTimeoutOverride *bool
 *   UseDualStack *bool
 *   SleepDelay func(time.Duration)
 *   DisableRestProtocolURICleaning *bool
 *   EnableEndpointDiscovery *bool
 *   DisableEndpointHostPrefix *bool
 * }
 *
 */

/**
 * package: credentials
 *
 * func NewStaticCredentialsFromCreds(creds credentials.Value) *Credentials
 *
 * type Value struct {
 *   AccessKeyID string
 *   SecretAccessKey string
 *   SessionToken string
 *   ProviderName string
 * }
 *
 */

/**
 * package: s3manager
 *
 * func NewUploader(c client.ConfigProvider, options ...func(*Uploader)) *Uploader
 *
 * ex)
 * sess := session.Must(session.NewSessionWithOptions(session.Options{}))
 * uploader := s3manager.NewUploader(sess)
 *
 */
func NewAwsS3() *AwsS3 {

	config := NewConfig()

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			// Access key ID と Secret Access Key は IAM から作成する。
			// 一度作成した Access key ID の Secret Access Key は csv ファイルでダウンロードしていない場合は、
			// 確認する手段がないので新しくアクセスキーを作成する必要がある。
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID: config.Aws.S3.AccessKeyID,
				SecretAccessKey: config.Aws.S3.SecretAccessKey,
			}),
			// アジアパシフィック (東京): ap-northeast-1
			Region: aws.String(config.Aws.S3.Region),
		},
	}))

	return &AwsS3{
		Config: config,
		Keys: AwsS3URLs{
			Test: "test/images",
		},
		// Create an uploader with the session and default options
		Uploader: s3manager.NewUploader(sess),
	}
}


/**
 * package: s3manager
 *
 * func (u Uploader) Upload(input *s3manager.UploadInput, options ...func(*Uploader)) (*s3manager.UploadOutput, error)
 *
 * type UploadInput struct {
 *   ACL *string `location:"header" locationName:"x-amz-acl" type:"string" enum:"ObjectCannedACL"`
 *   Body io.Reader
 *   Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
 *   CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`
 *   ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`
 *   ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`
 *   ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`
 *   ContentMD5 *string `location:"header" locationName:"Content-MD5" type:"string"`
 *   ContentType *string `location:"header" locationName:"Content-Type" type:"string"`
 *   Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`
 *   GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`
 *   GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`
 *   GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`
 *   GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`
 *   Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`
 *   Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`
 *   ObjectLockLegalHoldStatus *string `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"ObjectLockLegalHoldStatus"`
 *   ObjectLockMode *string `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"ObjectLockMode"`
 *   ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`
 *   RequestPayer *string `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"RequestPayer"`
 *   SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`
 *   SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`
 *   SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`
 *   SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`
 *   ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`
 *   StorageClass *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`
 *   Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`
 *   WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
 * }
 *
 * type UploadOutput struct {
 *   Location string
 *   VersionID *string
 *   UploadID string
 * }
 *
 */
func (a *AwsS3) UploadTest(file multipart.File, fileName string, extension string) (url string, err error) {

	if fileName == "" {
		return "", errors.New("fileName is required")
	}

	var contentType string

	switch extension {
	case "jpg":
		contentType = "image/jpeg"
	case "jpeg":
		contentType = "image/jpeg"
	case "gif":
		contentType = "image/gif"
	case "png":
		contentType = "image/png"
	default:
		return "", errors.New("this extension is invalid")
	}

	// Upload the file to S3.
	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		// ACL の設定は重要
		ACL: aws.String("public-read"),
		Body: file,
		Bucket: aws.String(a.Config.Aws.S3.Bucket),
		// content-type の設定も忘れずに
		ContentType: aws.String(contentType),
		Key: aws.String(a.Keys.Test + "/" + fileName + "." + extension),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}
