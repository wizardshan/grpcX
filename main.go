package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"context"
	"github.com/wizardshan/grpcX/domain"
	"github.com/wizardshan/grpcX/request"
	"github.com/wizardshan/grpcX/server/serverconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

// server is used to implement helloworld.GreeterServer.
type srv struct {
	serverconnect.UnimplementedGreeterHandler
}

func (s *srv) SayHello(ctx context.Context, req *connect.Request[request.UserGetRequest]) (*connect.Response[request.UserGetResponse], error) {
	log.Println(req.Header().Get("Some-Header"))
	res := connect.NewResponse(&request.UserGetResponse{
		// req.Msg is a strongly-typed *pingv1.PingRequest, so we can access its
		// fields without type assertions.
		ID: &domain.ID{Title: "fffffffff"},
	})
	res.Header().Set("Some-Other-Header", "hello!")
	return res, nil
}

func main() {

	//msg := &domain.ID{
	//	Title: "",
	//}
	//
	//v, err := protovalidate.New()
	//if err != nil {
	//	fmt.Println("failed to initialize validator:", err)
	//}
	//
	//if err = v.Validate(msg); err != nil {
	//	fmt.Println("validation failed:", err)
	//} else {
	//	fmt.Println("validation succeeded")
	//}

	interceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal(err)
	}

	interceptors := connect.WithInterceptors(interceptor)

	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	mux.Handle(serverconnect.NewGreeterHandler(&srv{}, interceptors))
	err = http.ListenAndServe(
		"localhost:8080",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)

	//var interceptor grpc.UnaryServerInterceptor
	//interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//	// 继续处理请求
	//	if r, ok := req.(Validator); ok {
	//		if err := r.Validate(); err != nil {
	//			return nil, status.Error(codes.InvalidArgument, err.Error())
	//		}
	//	}
	//	return handler(ctx, req)
	//}
	//var opts []grpc.ServerOption
	//opts = append(opts, grpc.UnaryInterceptor(interceptor))
	//
	//flag.Parse()
	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer(opts...)
	//server.RegisterGreeterServer(s, &srv{})
	//log.Printf("server listening at %v", lis.Addr())
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
