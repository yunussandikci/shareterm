module github.com/yunussandikci/shareterm/server

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0
	github.com/sethvargo/go-password v0.1.3
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/valyala/fasttemplate v1.1.1 // indirect
	github.com/yunussandikci/shareterm/common v0.0.3
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
)

//replace github.com/yunussandikci/shareterm/common => ../common
