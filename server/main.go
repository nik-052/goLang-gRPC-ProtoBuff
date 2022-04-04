package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	pb "personbio-API/personBiopb"

	"github.com/jackc/pgx/v4"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (*server) PerService(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	id := in.GetId()
	name := in.GetName()
	age := in.GetAge()
	gender := in.GetGender()
	fmt.Println("Recieved info")

	createDB(id, name, age, gender)

	return &pb.PersonResponse{Result: "Success"}, nil

}

func main() {
	fmt.Println("HEllo from server")

	lis, err := net.Listen("tcp", "0.0.0.0:0001")
	if err != nil {
		log.Fatalf("Error while listening , server %v", err)
	}

	sGRCP := grpc.NewServer()
	pb.RegisterPersonServiceServer(sGRCP, &server{})

	if err := sGRCP.Serve(lis); err != nil {
		log.Fatalf("Error while runnig perService %v", err)
	}

}

func createDB(id int64, name string, age int64, gender string) {

	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	createSql := ` 
	create table if not exists personBio(
		person_id int,
		name VARCHAR,
		age int,
		gender VARCHAR
	);
`
	_, error := conn.Exec(context.Background(), createSql)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Table creation: %v\n", error)
		os.Exit(1)
	}
	_, newerr := conn.Exec(context.Background(), "insert into personBio(person_id,name,age,gender) values ($1,$2,$3,$4)", id, name, age, gender)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", newerr)
		os.Exit(1)
	}

	_ , as := conn.Query(context.Background(), "insert into personBio(person_id,name,age,gender) values ($1,$2,$3,$4)", id, name, age, gender)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", as)
		os.Exit(1)
	}



	//idnew := 12
	//results, err := conn.Exec(context.Background(), "select * from personBio")
	//log.Println(results)

	/* AllRows, errorRet := conn.Query(context.Background(), "select * from personBio")
    if errorRet != nil {
        fmt.Fprintf(os.Stderr, "Retrival Failed : %v\n", errorRet)
        os.Exit(1)
    }
    defer AllRows.Close()
    for AllRows.Next() {
        values, err := AllRows.Values()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Retrival Failed : %v\n", err)
            os.Exit(1)
        }
        log.Fatal(values[0], values[1], values[2], values[3])
    } */
	




}
