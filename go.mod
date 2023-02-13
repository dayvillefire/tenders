module github.com/dayvillefire/tenders

go 1.19

replace (
	github.com/dayvillefire/tenders/api => ./api
	github.com/dayvillefire/tenders/auth => ./auth
	github.com/dayvillefire/tenders/common => ./common
	github.com/dayvillefire/tenders/config => ./config
	github.com/dayvillefire/tenders/models => ./models
)

require (
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/appleboy/gin-jwt/v2 v2.9.1
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/dayvillefire/tenders/api v0.0.0-20210303212310-ab20a0dbcce2
	github.com/dayvillefire/tenders/auth v0.0.0-20210303212310-ab20a0dbcce2
	github.com/dayvillefire/tenders/common v0.0.0-20210303212310-ab20a0dbcce2
	github.com/dayvillefire/tenders/config v0.0.0-20210303212310-ab20a0dbcce2
	github.com/dayvillefire/tenders/models v0.0.0-20210303212310-ab20a0dbcce2
	github.com/fatih/color v1.14.1 // indirect
	github.com/gin-gonic/contrib v0.0.0-20221130124618-7e01895a63f2
	github.com/gin-gonic/gin v1.8.2
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gobuffalo/envy v1.10.2 // indirect
	github.com/gobuffalo/fizz v1.14.4 // indirect
	github.com/gobuffalo/flect v1.0.0 // indirect
	github.com/gobuffalo/logger v1.0.7 // indirect
	github.com/gobuffalo/nulls v0.4.2 // indirect
	github.com/gobuffalo/packd v1.0.2 // indirect
	github.com/gobuffalo/plush/v4 v4.1.18 // indirect
	github.com/gobuffalo/uuid v2.0.5+incompatible
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/microcosm-cc/bluemonday v1.0.22 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/serenize/snaker v0.0.0-20201027110005-a7ad2135616e // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
