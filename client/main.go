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
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	/////////////////////////////////////////////////////

	c2 := pb.NewWhitelistClient(conn)
	r2, err2 := c2.GetWhitelist(context.Background(), &pb.NoParam{})
	if err2 != nil {
		log.Fatalf("could not GetWhitelist: %v", err2)
	}
	log.Printf("GetWhitelist: %s", r2.Ip)

	/////////////////////////////////////////////////////

	const Pub = `-----BEGIN ECDSA PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzqR158ptAz23PsGiKeAAQfdgaUP3
1j7hyO4lqc+b1rUwsCW9ED5P94ysslg6e75MT6UCKYLqRYlIr3bOqfT51w==
-----END ECDSA PUBLIC KEY-----`

	c3 := pb.NewCAClient(conn)
	r3, err3 := c3.IssueCertificate(context.Background(), &pb.CertificateRequest{In:[]byte(Pub), Name:"clientCertReq"})
	if err3 != nil {
		log.Fatalf("could not IssueCertificate: %v", err3)
	}
	log.Printf("IssueCertificate: %s", r3.In)

	r3, err3 = c3.GetCACertificate(context.Background(), &pb.NoParam{})
	if err3 != nil {
		log.Fatalf("could not GetCACertificate: %v", err3)
	}
	log.Printf("GetCACertificate: %s", r3.In)


	const clientCert = `-----BEGIN CERTIFICATE-----
MIIB5jCCAYugAwIBAgIBATAKBggqhkjOPQQDAzBXMR8wHQYDVQQGExZwa2kuY2Eu
c3ViamVjdC5jb3VudHJ5MSQwIgYDVQQKExtwa2kuY2Euc3ViamVjdC5vcmdhbml6
YXRpb24xDjAMBgNVBAMTBVNpbGVpMB4XDTE2MTAxMDA2MzEyNloXDTE3MDEwODA2
MzEyNlowXzEfMB0GA1UEBhMWcGtpLmNhLnN1YmplY3QuY291bnRyeTEkMCIGA1UE
ChMbcGtpLmNhLnN1YmplY3Qub3JnYW5pemF0aW9uMRYwFAYDVQQDEw1jbGllbnRD
ZXJ0UmVxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzqR158ptAz23PsGiKeAA
QfdgaUP31j7hyO4lqc+b1rUwsCW9ED5P94ysslg6e75MT6UCKYLqRYlIr3bOqfT5
16NAMD4wDgYDVR0PAQH/BAQDAgKEMAwGA1UdEwEB/wQCMAAwDQYDVR0OBAYEBAEC
AwQwDwYDVR0jBAgwBoAEAQIDBDAKBggqhkjOPQQDAwNJADBGAiEAhY5JgVq2M4Jo
mpaqCUzucRc2dddyU+zuBK0x5C+pLJQCIQDK+fpm8XQtmBdtVH9jrnIaryn5VuIW
OMaWpUGvf57onQ==
-----END CERTIFICATE-----`

	const client2Cert  = `-----BEGIN CERTIFICATE-----
MIIBwTCCAUegAwIBAgIBATAKBggqhkjOPQQDAzApMQswCQYDVQQGEwJVUzEMMAoG
A1UEChMDSUJNMQwwCgYDVQQDEwNPQkMwHhcNMTYwMTIxMjI0OTUxWhcNMTYwNDIw
MjI0OTUxWjApMQswCQYDVQQGEwJVUzEMMAoGA1UEChMDSUJNMQwwCgYDVQQDEwNP
QkMwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAAR6YAoPOwMzIVi+P83V79I6BeIyJeaM
meqWbmwQsTRlKD6g0L0YvczQO2vp+DbxRN11okGq3O/ctcPzvPXvm7Mcbb3whgXW
RjbsX6wn25tF2/hU6fQsyQLPiJuNj/yxknSjQzBBMA4GA1UdDwEB/wQEAwIChDAP
BgNVHRMBAf8EBTADAQH/MA0GA1UdDgQGBAQBAgMEMA8GA1UdIwQIMAaABAECAwQw
CgYIKoZIzj0EAwMDaAAwZQIxAITGmq+x5N7Q1jrLt3QFRtTKsuNIosnlV4LR54l3
yyDo17Ts0YLyC0pZQFd+GURSOQIwP/XAwoMcbJJtOVeW/UL2EOqmKA2ygmWX5kte
9Lngf550S6gPEWuDQOcY95B+x3eH
-----END CERTIFICATE-----`

	const rootCert = `-----BEGIN CERTIFICATE-----
MIIBzzCCAXWgAwIBAgIBATAKBggqhkjOPQQDAzBXMR8wHQYDVQQGExZwa2kuY2Eu
c3ViamVjdC5jb3VudHJ5MSQwIgYDVQQKExtwa2kuY2Euc3ViamVjdC5vcmdhbml6
YXRpb24xDjAMBgNVBAMTBVNpbGVpMB4XDTE2MTAwOTA3NDcwMVoXDTE3MDEwNzA3
NDcwMVowVzEfMB0GA1UEBhMWcGtpLmNhLnN1YmplY3QuY291bnRyeTEkMCIGA1UE
ChMbcGtpLmNhLnN1YmplY3Qub3JnYW5pemF0aW9uMQ4wDAYDVQQDEwVTaWxlaTBZ
MBMGByqGSM49AgEGCCqGSM49AwEHA0IABM6kdefKbQM9tz7BoingAEH3YGlD99Y+
4cjuJanPm9a1MLAlvRA+T/eMrLJYOnu+TE+lAimC6kWJSK92zqn0+dejMjAwMA4G
A1UdDwEB/wQEAwIChDAPBgNVHRMBAf8EBTADAQH/MA0GA1UdDgQGBAQBAgMEMAoG
CCqGSM49BAMDA0gAMEUCIDDNYY7dRsHjCixU4R4Rwc/v2XjGQ8oIsu5BHsI4KBmn
AiEAyEji95uqNpSQ+6Al5cFAe87lQvfT7lZMOIN2CwpXgxA=
-----END CERTIFICATE-----`

	r4, err4 := c3.VerifySignature(context.Background(), &pb.CertificateData{Cert:[]byte(clientCert), Root:[]byte(rootCert)})
	if err4 != nil {
		log.Fatalf("could not VerifySignature: %v", err4)
	}
	log.Printf("VerifySignature: %s", r4.Valid)
}
