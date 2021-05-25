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

var client = &http.Client{}

func main() {
	for {
		var reply = make([]byte, 1024)
		var rStdin = bufio.NewReader(os.Stdin)

		fmt.Println("1. Add user")
		fmt.Println("2. Show users")
		fmt.Println("3. Exit")
		var n, err = rStdin.Read(reply)
		if err != nil {
			log.Fatal(err)
		}

		reply = reply[:n]
		switch string(reply) {
		case "1\n":
			reply = make([]byte, 1024)

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

			var response, body = createNewRequest("http://localhost:8080/user/add", bytes.NewReader(data))
			defer response.Body.Close()
			fmt.Printf("%s\n", body)
		case "2\n":
			var response, body = createNewRequest("http://localhost:8080/user/show", nil)
			defer response.Body.Close()
			fmt.Printf("%s\n", body)
		case "3\n":
			fmt.Println("E X I T I N G . . .")
			os.Exit(0)
		default:
			var response, body = createNewRequest("http://localhost:8080/user/default", nil)
			defer response.Body.Close()
			fmt.Printf("%s\n", body)
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
