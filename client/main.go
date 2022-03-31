package main

import (
	"context"
	"fmt"
	"log"
	pb "personbio/personBiopb"

	"google.golang.org/grpc"
)

func main(){

	fmt.Println("hello from client")

	c,err:=grpc.Dial("localhost:8080",grpc.WithInsecure())
	if err!= nil{
		log.Fatalf("Error while connecting , client %v",err)
	}

	defer c.Close()

	cc:=pb.NewPersonServiceClient(c)

	doPrintperson(cc)

}

func doPrintperson(cc pb.PersonServiceClient){

	req := &pb.PersonRequest{
			Id:1,
			Name:"Nikhil Saji",
			Age:22,
			Gender: "M",
	}

	res, err := cc.PerService(context.Background(),req)

	if err != nil{
		log.Fatalf("Big error while request client %v",err)

	}

	log.Printf("Response from Person service %v",res)

}