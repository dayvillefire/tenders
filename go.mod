module github.com/dayvillefire/tenders

go 1.13

require (
	github.com/appleboy/gin-jwt/v2 v2.6.2
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/cockroachdb/cockroach-go v0.0.0-20190925194419-606b3d062051 // indirect
	github.com/gin-gonic/contrib v0.0.0-20181101072842-54170a7b0b4b
	github.com/gin-gonic/gin v1.6.2
	github.com/gobuffalo/buffalo v0.15.3 // indirect
	github.com/gobuffalo/envy v1.8.1
	github.com/gobuffalo/packr/v2 v2.7.1
	github.com/gobuffalo/pop v4.13.0+incompatible
	github.com/gobuffalo/suite v2.8.2+incompatible
	github.com/gobuffalo/uuid v2.0.5+incompatible
	github.com/gobuffalo/validate v2.0.3+incompatible
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/jackc/pgconn v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	golang.org/x/oauth2 v0.0.0-20181203162652-d668ce993890
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/dayvillefire/tenders/models => ../models
