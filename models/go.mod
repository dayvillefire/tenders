module github.com/dayvillefire/tenders/models

go 1.23.0

toolchain go1.24.0

replace (
	github.com/dayvillefire/tenders/common => ../common
	github.com/dayvillefire/tenders/config => ../config
)

require (
	github.com/dayvillefire/tenders/common v0.0.0-20240218192001-01cf76709c20
	github.com/gobuffalo/envy v1.10.2
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/gobuffalo/pop v4.13.1+incompatible
	github.com/gobuffalo/suite v2.8.2+incompatible
	github.com/gobuffalo/validate v2.0.4+incompatible
	github.com/gofrs/uuid v4.4.0+incompatible
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Masterminds/semver/v3 v3.3.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/bytedance/sonic v1.13.2 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/cockroachdb/cockroach-go v2.0.1+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dayvillefire/tenders/config v0.0.0-20240218192001-01cf76709c20 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.0.0 // indirect
	github.com/gin-gonic/gin v1.10.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.25.0 // indirect
	github.com/go-sql-driver/mysql v1.9.1 // indirect
	github.com/gobuffalo/buffalo v0.16.21 // indirect
	github.com/gobuffalo/events v1.4.1 // indirect
	github.com/gobuffalo/fizz v1.14.4 // indirect
	github.com/gobuffalo/flect v1.0.3 // indirect
	github.com/gobuffalo/genny v0.6.0 // indirect
	github.com/gobuffalo/github_flavored_markdown v1.1.4 // indirect
	github.com/gobuffalo/helpers v0.6.7 // indirect
	github.com/gobuffalo/httptest v1.5.0 // indirect
	github.com/gobuffalo/logger v1.0.7 // indirect
	github.com/gobuffalo/meta v0.3.0 // indirect
	github.com/gobuffalo/mw-csrf v1.0.0 // indirect
	github.com/gobuffalo/nulls v0.4.2 // indirect
	github.com/gobuffalo/packd v1.0.2 // indirect
	github.com/gobuffalo/plush v3.8.3+incompatible // indirect
	github.com/gobuffalo/plush/v4 v4.1.22 // indirect
	github.com/gobuffalo/pop/v5 v5.2.0 // indirect
	github.com/gobuffalo/tags/v3 v3.1.4 // indirect
	github.com/gobuffalo/validate/v3 v3.3.3 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/gorilla/handlers v1.4.2 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.6.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200307190119-3430c5407db8 // indirect
	github.com/jackc/pgtype v1.3.0 // indirect
	github.com/jackc/pgx/v4 v4.6.0 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/markbates/errx v1.1.0 // indirect
	github.com/markbates/grift v1.5.0 // indirect
	github.com/markbates/oncer v1.0.0 // indirect
	github.com/markbates/refresh v1.11.1 // indirect
	github.com/markbates/safe v1.0.1 // indirect
	github.com/markbates/sigtx v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/microcosm-cc/bluemonday v1.0.27 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monoculum/formam v0.0.0-20200905010316-d7a8fbd33677 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/serenize/snaker v0.0.0-20201027110005-a7ad2135616e // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.15.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/term v0.30.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
