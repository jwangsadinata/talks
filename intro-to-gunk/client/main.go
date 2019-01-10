package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/jwangsadinata/talks/intro-to-gunk/proto/v1/pokemon"
	"google.golang.org/grpc"
)

const (
	address          = "localhost:9090"
	defaultPokemonID = "25"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPokemonServiceClient(conn)

	input := defaultPokemonID
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	id, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("invalid id: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.GetReq{ID: int32(id)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Pokemon #%d is %s (%s)", r.ID, r.Name, r.Rarity.String())
}
