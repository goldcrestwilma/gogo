package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"../contract"
)

const port = 8888

func main() {
	log.Printf("Server starting on port : $v", port)
	StartServer()
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	//rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	//log.Printf("Server starting on port %v\n", port)
	//http.Serve(l, nil)
	defer l.close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
