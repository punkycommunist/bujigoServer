package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting listening on port", 4200, "...")
	listener, err := net.Listen("tcp", ":4200")
	ch := make(chan string)
	var str string
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		str = handleConnection(c, ch)
		str = <-ch
		if str != "nil" {
			writeToFile(str)
		}
	}
}

func handleConnection(c net.Conn, cha chan string) (data string) {
	log.Printf("Serving [%s]\n", c.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println("Errore: [", err.Error(), "]")
			fmt.Fprintf(c, "errore 1")
			break
		}
		c.Write([]byte(string("420!\n")))
		c.Close()
		fmt.Print(data)
		return data
	}
	c.Close()
	return "nil"
}

func writeToFile(line string) {
	f, err := os.OpenFile("archive.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	line = line + "\n" //newline removed from ReadString delimitor
	f.WriteString(line)
}
