package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/http2"
)

func main() {
	var srv http.Server
	srv.Addr = ":8080"
	//Enable http2
	http2.ConfigureServer(&srv, nil)

	http.HandleFunc("/", index_main)

	fmt.Printf("Listening on port %s\n", srv.Addr)
	srv.ListenAndServeTLS("certs/localhost.cert", "certs/localhost.key")	
}

func index_main(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(`
TLS handshake successfully completed: %t
TLS SNI: %s
TLS cipher suite: %d
TLS version: %d
`,
		r.TLS.HandshakeComplete,
		r.TLS.ServerName,
		r.TLS.CipherSuite,
		r.TLS.Version)
	fmt.Printf("\nFull TLS configuration: %+v\n", r.TLS)
	w.Write([]byte("<h1><center> Hello from Go! </h1></center>"))
}
