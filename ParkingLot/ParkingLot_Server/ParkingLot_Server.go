package main

import (
	"context"

	"net"
	"log"

	pb "example.com/go-Parkinglot-grpc/Parkinglot"
	"google.golang.org/grpc"
)

const (
	port= ":50051"
)
func NewParkingLotServer() *ParkinigLotServer{
	return &ParkinigLotServer{
		Cars_list: &pb.CarDetails{},
	}
}
type ParkinigLotServer struct {
	pb.UnimplementedParkinigLotServer
	Cars_list *pb.CarDetails
}



func (server *ParkinigLotServer) Run() error {
	lis,err :=net.Listen("tcp", port)
	if err!=nil{
		log.Fatalf(" failed to connect to the por: %v", err)
	}

	s:= grpc.NewServer()
	pb.RegisterParkinigLotServer(s,server)
	log.Printf(" server listening at port %v", lis.Addr())
	return s.Serve(lis)
}

func (server *ParkinigLotServer) ParkNewCar(ctx context.Context, in *pb.CarDetailsInfo) (*pb.SlotInfo, error){
	log.Printf("IN parking: %v", in.GetCarId())
	var slot int32=0
	Park_list :=&pb.SlotInfo{CarName : in.GetCarName(), CarId: in.GetCarId(), CarColor: in.GetCarColor(), Slot: slot }
	if len(server.Cars_list.Details)==0{
		Park_list.Slot=1
		server.Cars_list.Details=append(server.Cars_list.Details,Park_list)
		return Park_list,nil;
	}
	var tmpslot int=(len(server.Cars_list.Details));
	for i:=0 ;i<tmpslot;i++{
		if (server.Cars_list.Details[i].CarId==""){
			tmpslot=i+1
			Park_list.Slot=int32(tmpslot)
			server.Cars_list.Details[i]=Park_list
			return Park_list,nil
		}
	}
	tmpslot+=1
	Park_list.Slot=int32(tmpslot)
	server.Cars_list.Details=append(server.Cars_list.Details,Park_list)
	return Park_list,nil;
}

func (server *ParkinigLotServer) ParkingDetails(ctx context.Context , in *pb.DetailsInfoParams) (*pb.CarDetails, error){
	return server.Cars_list, nil;
}

func (server *ParkinigLotServer) UnparkCar(ctc context.Context, in *pb.GetId) (*pb.CarDetails, error){
	log.Printf("Out parking: %v", in.GetCarId())
	var tmpslot int=(len(server.Cars_list.Details));
	Park_list :=&pb.SlotInfo{};
	for i:=0 ;i<tmpslot;i++{
		if (server.Cars_list.Details[i].CarId==in.GetCarId()){
			
			Park_list.Slot=0
			Park_list.CarName=""
			Park_list.CarId=""
			Park_list.CarColor=""
			server.Cars_list.Details[i]=Park_list
		}
	}
	return server.Cars_list,nil

}

func main(){
	var new_parkinglot *ParkinigLotServer= NewParkingLotServer()
	if err:= new_parkinglot.Run();err!=nil{
		log.Fatalf("failed to connect or serv: %v",err)
	}
}