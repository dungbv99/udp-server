package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func client(c *net.UDPConn, i int, wg *sync.WaitGroup){
	defer wg.Done()
	defer c.Close()
	for {
		text := "data " + strconv.Itoa(i)
		data := []byte(text + "\n")
		_, err := c.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Reply %v : %s\n", i , string(buffer[0:n]))
	}
}

func main() {

	CONNECT := "127.0.0.1:1234"

	s, err := net.ResolveUDPAddr("udp", CONNECT)
	c, err := net.DialUDP("udp", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	var wg sync.WaitGroup

	for i:=0; i< 8; i++{
		wg.Add(1)
		go client(c,i, &wg)
	}
	wg.Wait()

}