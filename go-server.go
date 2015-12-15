package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
        
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		fmt.Println(n)
		if err != nil {
			return
		}	
		req := string(buf[0:])
		fmt.Println(req)
		header := strings.Split(req,"\n")
		parts := strings.Split(header[0], " ")
		method := parts[0]
//		path := strings.TrimSpace(parts[1])
		version := strings.TrimSpace(parts[2])
		if version != "HTTP/1.1"{
			fmt.Println("version is, ", version)
			conn.Write([]byte("HTTP/1.1 505 HTTP Version Not Supported\n\r\n"))
			fmt.Println("Response: HTTP/1.1 505 HTTP Version Not Supported")
		} else if method != "GET"{
			 conn.Write([]byte("HTTP/1.1 405 Method Not Allowed\n\r\n"))
			fmt.Println("Response: HTTP/1.1 405 Method Not Allowed")	
		} else{
			conn.Write([]byte("HTTP/1.1 200 OK\n\r\n"))
			dat, err := ioutil.ReadFile("./index.html")
			check(err)
			conn.Write([]byte(string(dat)))
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
