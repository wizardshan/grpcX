package main

import (
	"context"
	"github.com/wizardshan/grpcX/domain"
	"github.com/wizardshan/grpcX/request"
	"github.com/wizardshan/grpcX/server/serverconnect"
	"log"
	"net/http"

	"connectrpc.com/connect"
)

func main() {
	client := serverconnect.NewGreeterClient(
		http.DefaultClient,
		"http://localhost:8080/",
	)
	req := connect.NewRequest(&request.UserGetRequest{
		ID: &domain.ID{
			Title: "11111",
		},
	})
	req.Header().Set("Some-Header", "hello from connect")
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
	log.Println(res.Header().Get("Some-Other-Header"))
}
