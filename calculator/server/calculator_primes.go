package main

import (
	"log"

	pb "github.com/ddiogoo/grpc/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream grpc.ServerStreamingServer[pb.CalculatorResponse]) error {
	log.Printf("Primes function was invoked with %v\n", in)
	primes := factorization(in)
	for _, prime := range primes {
		stream.Send(&pb.CalculatorResponse{Result: prime})
	}
	return nil
}

func factorization(in *pb.PrimesRequest) []int32 {
	var k int32 = 2
	var n int32 = in.Number
	var primes []int32
	for n > 1 {
		for n%k == 0 {
			primes = append(primes, k)
			n /= k
		}
		k++
	}
	return primes
}
