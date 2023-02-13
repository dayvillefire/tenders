module github.com/dayvillefire/tenders/models

go 1.19

replace (
	github.com/dayvillefire/tenders/common => ../common
	github.com/dayvillefire/tenders/config => ../config
)

require (
	github.com/dayvillefire/tenders/common v0.0.0-20210303212310-ab20a0dbcce2
	github.com/gobuffalo/buffalo v0.16.21 // indirect
	github.com/gobuffalo/envy v1.10.2
	github.com/gobuffalo/mw-csrf v1.0.0 // indirect
	github.com/gobuffalo/packr/v2 v2.8.1
	github.com/gobuffalo/pop v4.13.1+incompatible
	github.com/gobuffalo/suite v2.8.2+incompatible
	github.com/gobuffalo/validate v2.0.4+incompatible
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/serenize/snaker v0.0.0-20201027110005-a7ad2135616e // indirect
)
