package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/proto"
	"context"
	//"io/ioutil"
	"io"
	"bytes"
	//"strings"
	"bufio"
)

const (
	address     = "localhost:50051"

	Pub = `-----BEGIN ECDSA PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzqR158ptAz23PsGiKeAAQfdgaUP3
1j7hyO4lqc+b1rUwsCW9ED5P94ysslg6e75MT6UCKYLqRYlIr3bOqfT51w==
-----END ECDSA PUBLIC KEY-----`
)

func readFully(conn net.Conn) ([]byte, error) {

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])

		fmt.Println(n)

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}

func getCACertificate() []byte{
	// Set up a connection to the server.
	conn1, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn1.Close()

	c3 := pb.NewCAClient(conn1)

	r3, err3 := c3.GetCACertificate(context.Background(), &pb.NoParam{})
	if err3 != nil {
		fmt.Println("could not GetCACertificate: %v", err3)
	}
	fmt.Println("GetCACertificate: %s", r3.In)

	return r3.In
}

func verifySignature(cert, root []byte) bool {
	conn1, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn1.Close()

	c3 := pb.NewCAClient(conn1)

	r4, err4 := c3.VerifySignature(context.Background(), &pb.CertificateData{Cert:cert, Root:root})
	if err4 != nil {
		fmt.Println("could not VerifySignature: %v", err4)
	}

	fmt.Println("VerifySignature: %s", r4.Valid)

	return r4.Valid
}

func main() {

	fmt.Println("Launching server...")


	ln, _ := net.Listen("tcp", "127.0.0.1:8081")
	conn, _ := ln.Accept()

	{
		fmt.Print("begin read message")


		//message, _ := bufio.NewReader(conn).ReadString('\n')
		message := make([]byte, 1000)
		n, err := bufio.NewReader(conn).Read([]byte(message))
		if err != nil {
			fmt.Println("could not readfull: %v", err)
		}
		fmt.Println(n)

		msg := message[0:n]
		fmt.Println("Message Received:", string(msg))

		root := getCACertificate()
		valid := verifySignature(msg, root)


		if valid {
			conn.Write([]byte("certificate pass"))
		}else {
			conn.Write([]byte("certificate failed"))
		}

	}
}
