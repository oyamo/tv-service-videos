package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	pb "github.com/oyamoh-brian/tv-service-database/proto/database"
	"google.golang.org/grpc"
	"log"
)

const (
	dbServiceAddr = "localhost:50050"
	httpTCPAddr = ":8081"
)

type DatabaseClient struct {
	pb.DataBaseServiceClient
}



var (
	dbServiceConn         *grpc.ClientConn = nil
	dataBaseServiceClient *pb.DataBaseServiceClient
)

func GetDatabaseClient() (*pb.DataBaseServiceClient, error)  {
	if dbServiceConn == nil {
		connectDBService()
	}

	instance := pb.NewDataBaseServiceClient(dbServiceConn)
	return &instance , nil
}

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

//goland:noinspection ALL
func main()  {
	// Connect to the database
	var err error
	dataBaseServiceClient, err = GetDatabaseClient()

	if err != nil {

	}
	defer closeDatabaseService()

	// Declare new fibre app
	var app = fiber.New(fiber.Config{
		Concurrency:                  3,
		AppName:                      "tv-service-videos",
		ReduceMemoryUsage:            true,
	})

	// Declare Routes
	app.Get("/categories/", func(ctx *fiber.Ctx) error {
		var rpcResponse, err = (*(dataBaseServiceClient)).GetAllCategories(context.Background(), nil)

		if err != nil {
			resp := struct {
				Status     int32       `json:"status"`
				Categories []*pb.Category `json:"categories"`
				Message    string      `json:"message"`
			}{
				500,
				nil,
				err.Error(),
			}
			return ctx.JSON(resp)
		}

		return ctx.JSON(rpcResponse)
	})

	app.Get("/categories/:c", func(ctx *fiber.Ctx) error {
		var rpcResponse, err = (*(dataBaseServiceClient)).GetAllCategories(context.Background(), nil)

		if err != nil {
			resp := struct {
				Status     int32       `json:"status"`
				Categories []*pb.Category `json:"categories"`
				Message    string      `json:"message"`
			}{
				500,
				nil,
				err.Error(),
			}
			return ctx.JSON(resp)
		}

		return ctx.JSON(rpcResponse)
	})

	app.Listen(httpTCPAddr)

}