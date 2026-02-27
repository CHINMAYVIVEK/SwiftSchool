package server

// grpc.go — Sets up the gRPC server and gRPC-Gateway reverse proxy.

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"swiftschool/app/academics"
	"swiftschool/app/admissions"
	"swiftschool/app/auth"
	"swiftschool/app/common"
	"swiftschool/app/core"
	"swiftschool/internal/database"
	pb "swiftschool/proto"
)

// StartGRPCServer creates and starts the gRPC server on the given address.
// It registers all gRPC services and enables server reflection.
func StartGRPCServer(db *database.Database, grpcAddr string) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", grpcAddr, err)
	}

	grpcServer := grpc.NewServer()

	// ---- Register services ----

	// Core
	coreService := core.NewService(db)
	pb.RegisterCoreServiceServer(grpcServer, core.NewGRPCServer(coreService))

	// Auth
	authService := auth.NewService(db)
	pb.RegisterAuthServiceServer(grpcServer, auth.NewGRPCServer(authService))

	// Academics
	academicsService := academics.NewService(db)
	pb.RegisterAcademicsServiceServer(grpcServer, academics.NewGRPCServer(academicsService))

	// Admissions
	admissionsService := admissions.NewService(db)
	pb.RegisterAdmissionsServiceServer(grpcServer, admissions.NewGRPCServer(admissionsService))

	// Common (Documents & Notifications)
	commonService := common.NewService(db)
	pb.RegisterCommonModuleServiceServer(grpcServer, common.NewGRPCServer(commonService))

	// Enable gRPC server reflection (useful for grpcurl, grpcui)
	reflection.Register(grpcServer)

	log.Printf("gRPC server listening on %s", grpcAddr)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	return grpcServer, lis
}

// NewGatewayMux creates a gRPC-Gateway reverse proxy that forwards REST
// requests to the gRPC server at grpcAddr.
func NewGatewayMux(ctx context.Context, grpcAddr string) *runtime.ServeMux {
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register all gRPC-Gateway handlers
	if err := pb.RegisterCoreServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to register Core gateway: %v", err)
	}
	if err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to register Auth gateway: %v", err)
	}
	if err := pb.RegisterAcademicsServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to register Academics gateway: %v", err)
	}
	if err := pb.RegisterAdmissionsServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to register Admissions gateway: %v", err)
	}
	if err := pb.RegisterCommonModuleServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to register Common module gateway: %v", err)
	}

	return gwMux
}

// WrapGateway wraps the gRPC-Gateway ServeMux so it can be nested inside
// a standard http.ServeMux under the /api/ prefix.
func WrapGateway(gwMux *runtime.ServeMux) http.Handler {
	return gwMux
}
