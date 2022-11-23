package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/MSalehKoinWorks/grpc-one/student"
	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	log.Println("rpc FindStudentByEmail invoked")

	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("data/students.json")

	if err != nil {
		log.Fatalln("error when read file", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error when unmarshal data students", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1315")

	if err != nil {
		log.Fatalln("error when listen tcp", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when grpc serve", err.Error())
	}
}
