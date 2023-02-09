module github.com/dayvillefire/tenders

go 1.13

replace (
	github.com/dayvillefire/tenders/api => ./api
	github.com/dayvillefire/tenders/auth => ./auth
	github.com/dayvillefire/tenders/common => ./common
	github.com/dayvillefire/tenders/config => ./config
	github.com/dayvillefire/tenders/models => ./models
)

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/dayvillefire/tenders/api v0.0.0-00010101000000-000000000000
	github.com/dayvillefire/tenders/auth v0.0.0-00010101000000-000000000000
	github.com/dayvillefire/tenders/common v0.0.0-00010101000000-000000000000
	github.com/dayvillefire/tenders/config v0.0.0-00010101000000-000000000000
	github.com/dayvillefire/tenders/models v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/contrib v0.0.0-20181101072842-54170a7b0b4b
	github.com/gin-gonic/gin v1.7.7
	github.com/gobuffalo/uuid v2.0.5+incompatible
	github.com/natefinch/lumberjack v2.0.0+incompatible
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
