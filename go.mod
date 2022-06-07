module github.com/a1comms/gcp-iap-auth

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/namsral/flag v1.7.4-pre
)

go 1.13

replace github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.4.1
