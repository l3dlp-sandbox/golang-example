package urlParsing

import (
	"fmt"
	"net"
	"net/url"
)

func UrlParsingDemo() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	//
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
}
