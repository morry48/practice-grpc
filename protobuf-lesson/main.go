package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"protobuf-lesson/pb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test.@exapmle.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"projenctX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Cant serialize", err)
	}
	if err := ioutil.WriteFile("text.bin", binData, 0666); err != nil {
		log.Fatalln("cant write", err)
	}

	in, err := ioutil.ReadFile("text.bin")

	if err != nil {
		log.Fatal("cant read file", err)
	}

	readEmployee := &pb.Employee{}

	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("can't deserialize")
	}

	fmt.Println(readEmployee)
}
