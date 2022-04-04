package main

import (
	"fmt"
	"log"
	"net/http"
	pb "personbio-API/personBiopb"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type user_data struct {
	Id     int64  `json:"person_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Age    int64  `json:"age" binding:"required"`
	Gender string `json:"gender" binding:"required"`
}

func main() {

	fmt.Println("hello from client")

	c, err := grpc.Dial("localhost:0001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while connecting , client %v", err)
	}

	defer c.Close()

	cc := pb.NewPersonServiceClient(c)

	doPrintperson(cc)

}

func doPrintperson(cc pb.PersonServiceClient) {

	g := gin.Default()
	var input user_data

	g.GET("/PerService/:id/:name/:age/:gender", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			log.Fatalf("Id error %v", err)
		}
		name := ctx.Param("name")
		age, err := strconv.Atoi(ctx.Param("age"))
		gender := ctx.Param("gender")

		req := &pb.PersonRequest{Id: id, Name: name, Age: int64(age), Gender: gender}
		if response, err := cc.PerService(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Result ": response,
			})

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	g.POST("/createPerson", func(ctx *gin.Context) {

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req := &pb.PersonRequest{Id: input.Id, Name: input.Name, Age: int64(input.Age), Gender: input.Gender}
		if response, err := cc.PerService(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Result ": response,
			})

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		log.Println(input.Age)

	})

	g.Run(":0002")
}
