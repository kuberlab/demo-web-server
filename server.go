package main

import (
	"net/http"
	"os"
	"flag"
	"fmt"
)
var vc,dv,iv,name,release string
func main() {
	flag.StringVar(&vc,"v","","Container version")
	flag.Parse()
	dv = os.Getenv("SERVER_VERSION")
	iv = os.Getenv("USED_IMAGE")
	name = os.Getenv("VNAME")
	release = os.Getenv("VRELEASE")
	http.HandleFunc("/", HelloServer)
	fmt.Printf("Hello World!!! Run version='%s' App version='%s' image='%s'. Name='%s' Release='%s'\n",vc,dv,iv,name,release)
	http.ListenAndServe(":8082",nil)
}

func HelloServer(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello World!!! Run version='%s' App version='%s' image='%s'. Name='%s' Release='%s'",vc,dv,iv,name,release)))
}