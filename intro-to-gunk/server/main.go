package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/jwangsadinata/talks/intro-to-gunk/proto/v1/pokemon"
	"github.com/jwangsadinata/talks/intro-to-gunk/server/pokedex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct{}

func (s *Server) Get(ctxt context.Context, req *pb.GetReq) (*pb.Pokemon, error) {
	id := req.GetID()
	if id <= 0 || id > 151 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pokemon id")
	}

	pokemon := &pb.Pokemon{
		Name: pokedex.Pokedex[id],
		ID:   id,
	}

	switch id {
	case 151:
		pokemon.Rarity = pb.Rarity_MYTHIC
	case 144, 145, 146, 150:
		pokemon.Rarity = pb.Rarity_LEGENDARY
	default:
		pokemon.Rarity = pb.Rarity_NORMAL
	}
	return pokemon, nil
}

func main() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("starting the server, listening at :9090")
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterPokemonServiceServer(s, &Server{})
	log.Fatal(s.Serve(l))
}
