package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/vmdt/gostorage/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandShakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tcpTransport := p2p.NewTCPTranport(tcpOpts)
	prefix, _ := strings.CutPrefix(tcpTransport.ListenAddr, ":")
	fileOpts := FileServerOpts{
		StorageRoot:       "go_network_" + prefix,
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	server := NewFileServer(fileOpts)
	tcpTransport.OnPeer = server.OnPeer

	return server
}

func main() {
	server1 := makeServer(":3000", "")
	// server1.Start()

	server2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(server1.Start())
	}()

	time.Sleep(time.Second * 1)

	go server2.Start()

	time.Sleep(time.Second * 1)

	// for i := 0; i < 1; i++ {
	// 	data := bytes.NewReader([]byte("my big data file here!"))
	// 	if err := server2.Store(fmt.Sprintf("myprivatedata"), data); err != nil {
	// 		log.Fatalln("store data error: ", err)
	// 	}
	// 	time.Sleep(5 * time.Microsecond)
	// }

	r, err := server2.Get("myprivatedata")
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(r)
	if (err) != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	select {}
}
