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

var client *http.Client = &http.Client{}

func main() {
	for {
		var reply []byte = make([]byte, 1024)
		var rStdin *bufio.Reader = bufio.NewReader(os.Stdin)

		fmt.Println("1. Add user")
		fmt.Println("2. Show users")
		fmt.Println("3. Exit")
		reply, _, err := rStdin.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		switch string(reply) {
		case "1":

			reply = make([]byte, 1024)
			var data []byte
			fmt.Println("Enter a name")
			n, err := rStdin.Read(reply)
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

			response, body := createNewRequest("POST", "http://localhost:8080/user/add", bytes.NewReader(data))
			var status int = response.StatusCode
			if status != 200 {
				log.Fatal(response.Status)
			}
			defer response.Body.Close()
			fmt.Printf("%s\n", body)

		case "2":

			response, body := createNewRequest("GET", "http://localhost:8080/user/show", nil)
			var status int = response.StatusCode
			if status != 200 {
				log.Fatal(response.Status)
			}
			defer response.Body.Close()
			fmt.Printf("%s\n", body)

		case "3":

			fmt.Println("E X I T I N G . . .")
			os.Exit(0)

		default:

			fmt.Println("404 page not found")

		}
	}
}

func createNewRequest(method string, url string, content io.Reader) (response *http.Response, body []byte) {
	request, err := http.NewRequest(method, url, content)
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
