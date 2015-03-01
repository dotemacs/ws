// simple web server that servers the contents of a given directory
// if no directory was given, then server the contents of the pwd

package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
)

func main() {
	var port string = ":8765"

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dirPtr := flag.String("d", pwd, "a name of the directory which to serve")
	flag.Parse()

	fmt.Println("serving contents of:", *dirPtr)
	fmt.Printf("open http://localhost%v\n", port)

	http.Handle("/", http.FileServer(http.Dir(*dirPtr)))
	http.ListenAndServe(port, nil)
}
