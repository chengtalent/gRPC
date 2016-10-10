/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

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
