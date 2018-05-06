package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	pb "github.com/sak0620/chaos/test/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CustomerService struct {
	customers []*pb.Person
	m         sync.Mutex
}

func (cs *CustomerService) ListPerson(p *pb.RequestType, stream pb.CustomerService_ListPersonServer) error {
	cs.m.Lock()
	defer cs.m.Unlock()
	for _, p := range cs.customers {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil

}

func (cs *CustomerService) AddPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	cs.m.Lock()
	defer cs.m.Unlock()
	cs.customers = append(cs.customers, p)
	return new(pb.ResponseType), nil
}

func main() {
	//port := os.Getenv("SERVER_PORT")
	port := "8080"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCustomerServiceServer(server, new(CustomerService))

	go func() {
		log.Printf("start grpc server port: %s", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	server.GracefulStop()
}
