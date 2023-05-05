module github.com/dayvillefire/tenders/auth

go 1.19

replace (
	github.com/dayvillefire/tenders/common => ../common
	github.com/dayvillefire/tenders/config => ../config
	github.com/dayvillefire/tenders/models => ../models
)

require (
	github.com/appleboy/gin-jwt/v2 v2.9.1
	github.com/dayvillefire/tenders/config v0.0.0-20210303212310-ab20a0dbcce2
	github.com/dayvillefire/tenders/models v0.0.0-20210303212310-ab20a0dbcce2
	github.com/gin-gonic/gin v1.9.0
)

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cockroachdb/cockroach-go v2.0.1+incompatible // indirect
	github.com/dayvillefire/tenders/common v0.0.0-20210303212310-ab20a0dbcce2 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gobuffalo/envy v1.10.2 // indirect
	github.com/gobuffalo/events v1.4.3 // indirect
	github.com/gobuffalo/fizz v1.14.4 // indirect
	github.com/gobuffalo/flect v1.0.0 // indirect
	github.com/gobuffalo/genny v0.6.0 // indirect
	github.com/gobuffalo/github_flavored_markdown v1.1.3 // indirect
	github.com/gobuffalo/helpers v0.6.7 // indirect
	github.com/gobuffalo/httptest v1.5.2 // indirect
	github.com/gobuffalo/logger v1.0.7 // indirect
	github.com/gobuffalo/meta v0.3.3 // indirect
	github.com/gobuffalo/nulls v0.4.2 // indirect
	github.com/gobuffalo/packd v1.0.2 // indirect
	github.com/gobuffalo/plush/v4 v4.1.18 // indirect
	github.com/gobuffalo/pop v4.13.1+incompatible // indirect
	github.com/gobuffalo/tags/v3 v3.1.4 // indirect
	github.com/gobuffalo/validate v2.0.4+incompatible // indirect
	github.com/gobuffalo/validate/v3 v3.3.3 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jackc/pgconn v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.1 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/markbates/errx v1.1.0 // indirect
	github.com/markbates/grift v1.5.0 // indirect
	github.com/markbates/oncer v1.0.0 // indirect
	github.com/markbates/refresh v1.12.0 // indirect
	github.com/markbates/safe v1.0.1 // indirect
	github.com/markbates/sigtx v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/microcosm-cc/bluemonday v1.0.22 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monoculum/formam v3.5.5+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/serenize/snaker v0.0.0-20201027110005-a7ad2135616e // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
