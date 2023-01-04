package main

import (
	"MyProjects/RentCar_gRPC/auth_rentcar_service/config"
	"MyProjects/RentCar_gRPC/auth_rentcar_service/protogen/authorization"
	serices	"MyProjects/RentCar_gRPC/auth_rentcar_service/services/authorization" //! services

	"MyProjects/RentCar_gRPC/auth_rentcar_service/storage"
	"MyProjects/RentCar_gRPC/auth_rentcar_service/storage/postgres"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// * @license.name  Apache 2.0
// * @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	cfg := config.Load()
	psqlConfString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	var err error
	var stg storage.StorageInter
	stg, err = postgres.InitDB(psqlConfString)
	if err != nil {
		panic(err)
	}

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	authService := serices.NewAuthService(cfg, stg) //! <----- services
	authorization.RegisterAuthServiceServer(srv, authService)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
