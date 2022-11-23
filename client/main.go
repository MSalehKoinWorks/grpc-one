package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MSalehKoinWorks/grpc-one/student"
	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)

	if err != nil {
		log.Fatalln("error when get student by email")
	}

	log.Println(student)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1315", opts...)

	if err != nil {
		log.Fatalln("error when dial", err.Error())
	}

	defer conn.Close()

	client := pb.NewDataStudentClient(conn)

	getDataStudentByEmail(client, "m.saleh.solahudin@gmail.com")
}
