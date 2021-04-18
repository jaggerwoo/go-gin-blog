module go-gin-blog

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/go-ini/ini v1.62.0 // indirect
	github.com/go-playground/validator/v10 v10.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jaggerwoo/go-gin-blog/pkg/setting v0.0.0-00010101000000-000000000000
	github.com/jaggerwoo/go-gin-blog/routers v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.5 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc // indirect
	golang.org/x/sys v0.0.0-20210415045647-66c3f260301c // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/jaggerwoo/go-gin-blog/conf => ./conf
	github.com/jaggerwoo/go-gin-blog/middleware => ./middleware
	github.com/jaggerwoo/go-gin-blog/models => ./models
	github.com/jaggerwoo/go-gin-blog/pkg/setting => ./pkg/setting
	github.com/jaggerwoo/go-gin-blog/routers => ./routers
	github.com/jaggerwoo/go-gin-blog/routers/api/v1 => ./routers/api/v1
)
