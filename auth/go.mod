module github.com/dayvillefire/tenders/auth

go 1.16

replace (
	github.com/dayvillefire/tenders/common => ../common
	github.com/dayvillefire/tenders/config => ../config
	github.com/dayvillefire/tenders/models => ../models
)

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/dayvillefire/tenders/config v0.0.0-00010101000000-000000000000
	github.com/dayvillefire/tenders/models v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
)
