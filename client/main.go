package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "http://localhost:8080/user"

var client = &http.Client{}

func main() {
	var reply = make([]byte, 1024)
	var rStdin = bufio.NewReader(os.Stdin)

	for {
		//Show menu to server
		var response, body = createNewRequest(url, nil)
		defer response.Body.Close()
		fmt.Printf("%s", body)

		//Read option
		n, err := rStdin.Read(reply)
		if err != nil {
			log.Fatal(err)
		}

		//Send option
		response, body = createNewRequest(url, bytes.NewReader(reply[0:n]))
		defer response.Body.Close()

		body = body[:len(body)-1]
		body = body[bytes.LastIndex(body, []byte("\n"))+1:]

		switch string(body) {
		case "/add":
			var data []byte
			fmt.Println("Enter a name")
			n, err = rStdin.Read(reply)
			if err != nil {
				log.Fatal(err)
			}

			data = append(data, reply[:n]...)
			fmt.Println("Enter a password")
			n, err = rStdin.Read(reply)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, reply[:n]...)

			var response, body = createNewRequest(url+string(body), bytes.NewReader(data))
			defer response.Body.Close()
			fmt.Printf("%s", body)
		case "/show":
			var response, body = createNewRequest(url+string(body), nil)
			defer response.Body.Close()
			fmt.Printf("%s", body)
		default:
			var response, body = createNewRequest(url+string(body), nil)
			defer response.Body.Close()
			fmt.Printf("%s", body)
		}
	}
}

func createNewRequest(url string, content io.Reader) (response *http.Response, body []byte) {
	var request, err = http.NewRequest("POST", url, content)
	if err != nil {
		log.Fatal(err)
	}

	response, err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return
}
