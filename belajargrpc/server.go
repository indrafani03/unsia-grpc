package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"unsia/controllers"
	"unsia/pb/cities"
	"unsia/pkg/database"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":7080")

	if err != nil {
		fmt.Println("failed to listen %v", err)
		return
	}

	log := log.New(os.Stdout, "Esensial :", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	db, err := database.OpenDB()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()

	cityServer := controllers.City{DB: db}
	cities.RegisterCitiesServiceServer(grpcServer, &cityServer)

	fmt.Println("running server grpc")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve : %v", err)
		return
	}

}
