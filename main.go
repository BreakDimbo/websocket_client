package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/breakD/wb_client/api"
	"github.com/breakD/wb_client/wbclient"

	log "github.com/breakD/wb_client/logging"
)

/*
usage: ./breaksock [target address]
connecting....
connecte ok
receive <- :
send -> :
*/
func main() {
	tFlag := flag.String("t", "127.0.0.1:4022", "-t 127.0.0.0:4022")
	flag.Parse()
	log.Debugf("get flag from command line: %s\n", *tFlag)

	c := wbclient.New(*tFlag)
	c.Connect()

	for c.WsConn() == nil {
		// waiting for the ws connecting
		time.Sleep(time.Second)
	}

	api.Execute("login", c)
	fmt.Println("logging")
	time.Sleep(time.Second)

	reader := bufio.NewReader(os.Stdin)

	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Print("Enter api: ")
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		fmt.Println(text)
		api.Execute(text, c)
	}
}
