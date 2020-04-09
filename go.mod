module github.com/wappgo

go 1.13

require (
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/sessions v1.2.0
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71
	wappgo v0.0.0
)

replace wappgo => ./
