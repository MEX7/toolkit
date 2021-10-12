package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Subdomains map[string]http.Handler

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if mux := subdomains[r.Host]; mux != nil {
		mux.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not found other", 404)
	}
}

func main() {
	r := gin.Default()
	r2 := gin.Default()
	// r.RunTLS(":81",
	// 	"/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls.pem",
	// 	"/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls.key")
	// r2.RunTLS(":82",
	// 	"/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls.pem",
	// 	"/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls.key")

	p := "/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls.pem"
	k := "/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls.key"

	p2 := "/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls2.com.pem"
	k2 := "/Users/duminxiang/cosmos/go/src/github.com/kl7sn/toolkit/subds/cert/tls2.com.key"

	r.GET("/ping", adminHandlerOne)
	r2.GET("/ping", adminHandlerOne)

	s := &http.Server{Handler: r}
	s2 := &http.Server{Handler: r2}

	// certificates := make([]tls.Certificate, 2)
	// certificates[1], _ = tls.LoadX509KeyPair(p, k)
	// certificates[0], _ = tls.LoadX509KeyPair(p2, k2)
	//
	// s.TLSConfig = &tls.Config{}
	// s.TLSConfig.Certificates = certificates
	// s.ListenAndServeTLS(p, k)

	ln, _ := net.Listen("tcp", ":443")
	go s.ServeTLS(ln, p, k)
	s2.ServeTLS(ln, p2, k2)

}

func adminHandlerOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
