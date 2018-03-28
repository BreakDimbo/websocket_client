package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

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
	}

	api.Execute("login", c)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter api: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		api.Execute(text, c)
	}
}
