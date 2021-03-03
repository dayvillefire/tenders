module github.com/dayvillefire/tenders/models

go 1.16

replace (
	github.com/dayvillefire/tenders/common => ../common
	github.com/dayvillefire/tenders/config => ../config
)

require (
	github.com/dayvillefire/tenders/common v0.0.0-00010101000000-000000000000
	github.com/gobuffalo/buffalo v0.16.21 // indirect
	github.com/gobuffalo/envy v1.9.0
	github.com/gobuffalo/mw-csrf v1.0.0 // indirect
	github.com/gobuffalo/packr/v2 v2.8.1
	github.com/gobuffalo/pop v4.13.1+incompatible
	github.com/gobuffalo/suite v2.8.2+incompatible
	github.com/gobuffalo/validate v2.0.4+incompatible
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/gorilla/sessions v1.2.1 // indirect
)
