package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func server(connection *net.UDPConn, i int, wg *sync.WaitGroup){
	defer wg.Done()
	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())
	host, _ := os.Hostname()
	fmt.Println("host ", host)
	addrs, _ := net.LookupIP(host)
	var ipaddr string = "  "
	for _, addr := range addrs {
		fmt.Println("addr ", addr)
		if ipv4 := addr.To4(); ipv4 != nil {
			ipaddr += ipv4.String()
			break
		}
	}
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))



		data := []byte(strconv.Itoa(i) + " " + ipaddr + "    v4"  )
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}


func main() {

	PORT := ":1234"

	s, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp", s)

	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < 4; i++{
		wg.Add(1)
		go server(connection, i, &wg)
	}
	wg.Wait()

}
