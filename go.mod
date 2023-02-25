module github.com/wappgo

go 1.13

require (
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/sessions v1.2.0
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.1 // indirect
	wappgo v0.0.0
)

replace wappgo => ./
