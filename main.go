package main

import (
	"google.golang.org/grpc"
	"log"
)

const (
	dbServiceAddr = "localhost:50050"
	httpTCPAddr = ":8081"
)

var (
	dbServiceConn *grpc.ClientConn = nil
)

func connectDBService()  {
	var err error
	dbServiceConn , err = grpc.Dial(dbServiceAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not Connect: %v", err)
	}
}

func closeDatabaseService()  {
	if dbServiceConn != nil {
		func(dbServiceConn *grpc.ClientConn) {
			err := dbServiceConn.Close()
			if err != nil {
				log.Fatalf("Error while closing connection to DB @ %s : %v", dbServiceAddr, err)
			}
		}(dbServiceConn)
	}
}


func main()  {
	connectDBService()

	defer closeDatabaseService()

}