package main

import (
	"context"
	"log"
	"time"
	"fmt"
	pb "example.com/go-Parkinglot-grpc/Parkinglot"
	"google.golang.org/grpc"
)

const (
	address= ":50051"
)
func input()(string , string ,string){
	var Name string 
	var  Id string 
	var Color string 
	fmt.Println("Enter Name of Car :")
	fmt.Scanln(&Name)
	fmt.Println("Enter Id of your Car")

	fmt.Scanln(&Id)
	fmt.Println("Enter Color of your Car")
	fmt.Scanln(&Color)
	return Name,Id,Color

}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewParkinigLotClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10000)
	defer cancel()
	for {
		fmt.Println(" Choose one of the following to proceed further")
		fmt.Println("ParkNewCar GetDetails UnparkCar")
		var name string
		fmt.Scanln(&name)
		if (name=="ParkNewCar"){
			Name,Id,Color :=input()
			r1,err :=c.ParkNewCar(ctx, &pb.CarDetailsInfo{CarName: Name, CarId: Id, CarColor: Color })
			if err != nil {
				log.Fatalf("could not create user: %v", err)
			}
			fmt.Println("Slot the car to be parked is:")
			log.Println(r1.GetSlot())
		}else if (name=="GetDetails"){
			r1,err :=c.ParkingDetails(ctx, &pb.DetailsInfoParams{})
			if err != nil {
				log.Fatalf("could not create user: %v", err)
			}
			fmt.Println(r1)
		}else if(name=="UnparkCar"){
			var Id string
			fmt.Scanln(&Id)
			r2,err := c.UnparkCar(ctx, &pb.GetId{CarId : Id})
			if err != nil {
				log.Fatalf("could not create user: %v", err)
			}
			log.Println(r2)
		}else{
			fmt.Println("Try spell out correct one")
		}
	}
}
