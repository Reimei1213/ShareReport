package user

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "share-report/proto/user"
	"share-report/user/usecase"
	"share-report/user/db"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("USER_API_PORT"))
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, usecase.NewUserService(db))

	reflection.Register(server)
	server.Serve(listenPort)
}
