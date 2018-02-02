package main

import (
	"io/ioutil"
	"log"
	"github.com/michaelbironneau/asbclient"
	"fmt"
	"time"
    "strings"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main(){

	log.Println("Starting")

	path := "colomanager-to-rp"

	dat, err := ioutil.ReadFile("/etc/password/queuePassword")
	//dat, err := ioutil.ReadFile("C:\\temp\\password.txt")

	check(err)
	password := string(dat)
    password = strings.TrimSuffix(password, "\n")

	sbqClient := asbclient.New(asbclient.Queue, "ucs-service-bus", "sendlisten", password)
	go func() {
		for i:= 0; i < 10; i++{
			i := 1
			err := sbqClient.Send(path, &asbclient.Message{
				Body: []byte(fmt.Sprintf("Hello World from GoLang %d", i)),
			})

			if err != nil {
				log.Printf("Send Error!!!! %s", err)
			} else {
				log.Printf("Sent: %d", i)
			}
		}
	}()

	for {
		log.Printf("Peeking...")
		msg, err := sbqClient.PeekLockMessage(path, 30)

		if err != nil {
			log.Printf("Peek error: %s", err)
		} else {
			log.Printf("Peeked message: '%s'", string(msg.Body))
			err = sbqClient.DeleteMessage(msg)
			if err != nil {
				log.Printf("Delete error: %s", err)
			} else {
				log.Printf("Deleted message")
			}
		}

		time.Sleep(time.Millisecond * 200)
	}
}
