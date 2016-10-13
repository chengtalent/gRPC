package main

import (
	"net"
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/proto"
	"context"
)

const (
	address     = "localhost:50051"

	pub = `-----BEGIN CERTIFICATE-----
MIIB3DCCAYGgAwIBAgIBATAKBggqhkjOPQQDAzBXMR8wHQYDVQQGExZwa2kuY2Eu
c3ViamVjdC5jb3VudHJ5MSQwIgYDVQQKExtwa2kuY2Euc3ViamVjdC5vcmdhbml6
YXRpb24xDjAMBgNVBAMTBVNpbGVpMB4XDTE2MTAxMjA5Mzc1OVoXDTE3MDExMDA5
Mzc1OVowVTEfMB0GA1UEBhMWcGtpLmNhLnN1YmplY3QuY291bnRyeTEkMCIGA1UE
ChMbcGtpLmNhLnN1YmplY3Qub3JnYW5pemF0aW9uMQwwCgYDVQQDEwNwdWIwWTAT
BgcqhkjOPQIBBggqhkjOPQMBBwNCAATOpHXnym0DPbc+waIp4ABB92BpQ/fWPuHI
7iWpz5vWtTCwJb0QPk/3jKyyWDp7vkxPpQIpgupFiUivds6p9PnXo0AwPjAOBgNV
HQ8BAf8EBAMCAoQwDAYDVR0TAQH/BAIwADANBgNVHQ4EBgQEAQIDBDAPBgNVHSME
CDAGgAQBAgMEMAoGCCqGSM49BAMDA0kAMEYCIQCBKU6vWfMdzydughFM/FuAvqG6
CXqK8o6NaKZrfWxhAgIhAKXG2v0BgXTqdoJo/RS1gUDWnUFYDibMcNLne8O4f7Ry
-----END CERTIFICATE-----`
)

func issueCertificate() []byte{
	// Set up a connection to the server.
	conn1, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn1.Close()

	c3 := pb.NewCAClient(conn1)
	r3, err3 := c3.IssueCertificate(context.Background(), &pb.CertificateRequest{In:[]byte(pub), Name:"pub"})
	if err3 != nil {
		fmt.Println("could not IssueCertificate: %v", err3)
	}
	fmt.Println("IssueCertificate: %s", r3.In)

	return r3.In
}

func main() {


	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	{
		// read in input from stdin
		//reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Text to send: ")
		//text, _ := reader.ReadString('\n')

		// send to socket
		//fmt.Fprintf(conn, text + "\n")

		fmt.Println("begin Write message")

		//cert := issueCertificate()

		if _, err := conn.Write([]byte(pub)); err != nil {
			fmt.Println(err)
		}

		fmt.Println("end Write message")

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from server: "+message)
	}
}