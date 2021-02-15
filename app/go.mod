module github.com/psychedelicnekopunch/go-sample

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/aws/aws-sdk-go v1.19.18
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/jinzhu/gorm v1.9.4
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/nlopes/slack v0.6.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
