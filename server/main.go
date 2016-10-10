package main

import (
	"log"
	"net"

	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	ca "google.golang.org/grpc/examples/helloworld/ca"
	"google.golang.org/grpc/examples/helloworld/crypto"
	pb "google.golang.org/grpc/examples/helloworld/proto"
)

const (
	port = ":50051"
)

var slogger = logging.MustGetLogger("server")
var cap *ca.CA


type whitelistServer struct{}

func (s *whitelistServer) GetWhitelist(ctx context.Context, in *pb.NoParam) (*pb.IPList, error) {
	res := &pb.IPList{}
	res.Ip = make([]string, 2)
	res.Ip[0] = "127.0.0.1"
	res.Ip[1] = "192.168.0.1"

	return res, nil
}

type CAServer struct{}

func (s *CAServer)IssueCertificate(ctx context.Context, cr *pb.CertificateRequest) (*pb.CertificateReply, error) {
	if cap == nil {
		return nil, nil
	}

	reply := pb.CertificateReply{}
	if cert, err := cap.IssueCertificate(cr.In, cr.Name); err != nil {
		slogger.Panicf("Failed IssueCertificate [%s]", err)
		return nil, err
	}else {
		reply.In = cert
	}

	return &reply, nil
}

func (s *CAServer)GetCACertificate(ctx context.Context, np *pb.NoParam) (*pb.CertificateReply, error) {
	if cap == nil {
		return nil, nil
	}

	reply := pb.CertificateReply{}
	reply.In = cap.GetCACertificate()

	return &reply, nil
}

func (s *CAServer)VerifySignature(ctx context.Context, certData *pb.CertificateData) (*pb.SignatureValid, error) {
	valid := pb.SignatureValid{}

	err := ca.VerifySignature(certData.Cert, certData.Root)
	valid.Valid = err == nil

	return &valid, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterWhitelistServer(s, &whitelistServer{})
	pb.RegisterCAServer(s, &CAServer{})

	//////////////////////////////////////////////////

	// Init the crypto layer
	if err := crypto.Init(); err != nil {
		slogger.Panicf("Failed initializing the crypto layer [%s]", err)
	}

	ca.CacheConfiguration()
	cap = ca.NewCA("Silei", ca.InitializeCommonTables)

//	const Pub = `-----BEGIN PUBLIC KEY-----
//MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEs0Hsfojry7g3TLBzID4JjjIhGJF2GMJ5
//acT38++yWsju1UKRWUxFrfqJXjRYz4yf5dduk6pbPWGOUdfdAOAPJQ==
//-----END PUBLIC KEY-----`


//	const Pub = `-----BEGIN ECDSA PUBLIC KEY-----
//MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzqR158ptAz23PsGiKeAAQfdgaUP3
//1j7hyO4lqc+b1rUwsCW9ED5P94ysslg6e75MT6UCKYLqRYlIr3bOqfT51w==
//-----END ECDSA PUBLIC KEY-----`
//
//	cap.IssueCertificate([]byte(Pub), "test")

	s.Serve(lis)
}
