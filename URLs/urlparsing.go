package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	pl := fmt.Println
	sUrl := "https://dean:wrris@nowhere.net:3939/path?x=test"

	u, err := url.Parse(sUrl)

	if err != nil {
		panic(nil)
	}

	pl("Scheme", u.Scheme)
	pl("User", u.User.Username())
	pl("Password", u.User.Username())
	pl("Host", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	pl("Domain", host)
	pl("Port", port)
	pl("Path", u.Path)
	pl("Fragment", u.Fragment)
	pl("Raw Query", u.RawQuery)
	q, _ := url.ParseQuery(u.RawQuery)
	pl("X", q["x"][0]) // treated as an array

}
