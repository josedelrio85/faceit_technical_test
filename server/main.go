package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	faceit_cc "github.com/josedelrio85/faceit_technical_test/server/pkg"
)

func main() {

	dt := faceit_cc.NewDatabase()
	defer dt.Database.Db.Close()

	kafka := faceit_cc.KafkaInstance{}
	err := kafka.Initialize()
	if err != nil {
		log.Fatalf(err.Error())
	}
	dt.Kafka = kafka

	server := rpc.NewServer()
	faceit_cc.RegisterUser(server, *dt)

	imp := faceit_cc.BasicImplementation{}
	faceit_cc.RegisterStatus(server, imp)

	// Register a HTTP handler
	server.HandleHTTP("/", "/debug")

	listener, e := net.Listen("tcp", ":9001")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Printf("Serving RPC server on port %d", 9001)

	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
