package main

import (
	"fmt"
	"log"
	"os"

	"learning-grpc/rides/pb"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	req := pb.StartRequest{
		Id:       "1",
		DriverId: "007",
		Location: &pb.Location{
			Lat: 51.4567865,
			Lng: -0.1234556,
		},
		PassengerIds: []string{"M", "Q"},
		Time:         timestamppb.Now(),
		Type:         pb.RideType_POOL,
	}

	fmt.Println(&req)

	data, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var req2 pb.StartRequest
	if err := proto.Unmarshal(data, &req2); err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(&req2)

	fmt.Println("proto size:", len(data))

	jdata, err := protojson.Marshal(&req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println("json size:", len(jdata))

	os.Stdout.Write(jdata)
}
