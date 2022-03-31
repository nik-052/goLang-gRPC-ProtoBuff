package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "personBio/personBiopb"

	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedPersonServiceServer
}

func (*server) PerService(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	id := in.GetId()
	name := in.GetName()
	age := in.GetAge()
	gender := in.GetGender()
	fmt.Println("Recieved info")

	return &pb.PersonResponse{Id: id, Name: name, Age: age, Gender: gender}, nil

}

func main() {
	fmt.Println("HEllo from server")

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Error while listening , server %v", err)
	}

	sGRCP:=grpc.NewServer()
	pb.RegisterPersonServiceServer(sGRCP,&server{})

	if err:=sGRCP.Serve(lis);err!= nil{
		log.Fatalf("Error while runnig perService %v",err)
	}

}
