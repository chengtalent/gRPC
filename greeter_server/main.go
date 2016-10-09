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
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

var slogger = logging.MustGetLogger("server")

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

type whitelistServer struct{}

func (s *whitelistServer) GetWhitelist(ctx context.Context, in *pb.NoParam) (*pb.IPList, error) {
	res := &pb.IPList{}
	res.Ip = make([]string, 2)
	res.Ip[0] = "127.0.0.1"
	res.Ip[1] = "192.168.0.1"

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterWhitelistServer(s, &whitelistServer{})

	//////////////////////////////////////////////////

	// Init the crypto layer
	if err := crypto.Init(); err != nil {
		slogger.Panicf("Failed initializing the crypto layer [%s]", err)
	}

	ca.CacheConfiguration()
	ca := ca.NewCA("Silei", ca.InitializeCommonTables)

	const Pub = `-----BEGIN PUBLIC KEY-----
MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEs0Hsfojry7g3TLBzID4JjjIhGJF2GMJ5
acT38++yWsju1UKRWUxFrfqJXjRYz4yf5dduk6pbPWGOUdfdAOAPJQ==
-----END PUBLIC KEY-----`

	ca.IssueCertificate([]byte(Pub), "test")

	s.Serve(lis)
}
