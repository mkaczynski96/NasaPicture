package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Resp struct {
	*http.Response
	err error
}

// Get method
func MakeGetRequest(url string, reqs int, rc chan Resp, sem chan bool) {
	defer close(rc)
	defer close(sem)
	for reqs > 0 {
		select {
		case sem <- true:
			req, _ := http.NewRequest("GET", url, nil)
			transport := &http.Transport{}
			resp, err := transport.RoundTrip(req)
			r := Resp{resp, err}
			rc <- r
			reqs--
		default:
			<-sem
		}
	}
}

func CheckResponses(rc chan Resp) []byte {
	conns := 0
	for {
		select {
		case r, ok := <-rc:
			if ok {
				conns++
				if r.err != nil {
					fmt.Println(r.err)
				} else {
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						fmt.Println(r.err)
					}
					if err := r.Body.Close(); err != nil {
						fmt.Println(r.err)
					}
					return body
				}
			} else {
				return nil
			}
		}
	}
}
