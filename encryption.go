package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stringchan := make(chan string)
	tower1chan := make(chan string)
	tower2chan := make(chan string)
	var offset int32 = 3
	go tower1(stringchan, tower1chan, offset)
	go tower2(stringchan, tower2chan, offset)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-tower1chan:
			fmt.Printf("Control message from tower1 - %v", msg)
		case msg := <-tower2chan:
			fmt.Printf("Control message from tower2 - %v", msg)
		}
	}
}

func tower1(s chan string, t1 chan string, offset int32) {
	inputStream := bufio.NewReader(os.Stdin)
	fmt.Println("\nTower1: Enter your message for tower2")
	userInput, _ := inputStream.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	fmt.Printf("\nTower1: Original string: %s", userInput)
	var secretString string
	for _, c := range userInput {
		secretString += string(c + offset)
	}
	fmt.Printf("\nTower1 encrypted string: %s", secretString)

	s <- secretString
	t1 <- "Msg send to tower 2"
}

func tower2(s chan string, t2 chan string, offset int32) {
	secretString := <-s
	var orgString string
	for _, c := range secretString {
		orgString += string(c - offset)
	}
	fmt.Printf("\nTower2 decrypted message : %s", orgString)
	t2 <- "Msg received from tower 1"

}
