package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {

	PORT := ":17777"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleTCPConn(c)
	}
}

func handleTCPConn(c net.Conn) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Printf("Got ERROR from %s  %s \n ", c.RemoteAddr().String(), err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			log.Printf("Got STOP from %s\n", c.RemoteAddr().String())
			break
		}

		random := rand.New(rand.New(rand.NewSource((time.Now().UnixMilli())))).Intn(100000)

		var result strings.Builder

		// result = strconv.Itoa(random) + "\n"

		result.WriteString(strconv.Itoa(random))
		result.WriteString(" ")
		result.WriteString(time.Now().Format(time.RFC850))


		c.Write([]byte(result.String()))

	}

	c.Close()

}
