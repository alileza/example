package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	pb "github.com/alileza/example/autogen/pb"
	"github.com/alileza/example/services"
)

type Server struct {
	ListenAddress       string
	SwaggerDocsFilePath string

	httpSrv *http.Server
	grpcSrv *grpc.Server
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		return fmt.Errorf("Cannot listen to port: %w", err)
	}

	m := cmux.New(listener)
	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.Match(cmux.HTTP2())

	runtimeMux := createRuntimeMux()

	mux := http.NewServeMux()
	mux.Handle("/api/", runtimeMux)
	mux.Handle("/api/docs", docsHandler(s.SwaggerDocsFilePath))
	mux.Handle("/metrics", promhttp.Handler())
	s.grpcSrv = grpc.NewServer()
	s.httpSrv = &http.Server{Handler: mux}

	grpcDialOpts := []grpc.DialOption{grpc.WithInsecure()}

	// Registering service api
	pb.RegisterExampleServiceV1Server(s.grpcSrv, &services.ExampleService{})
	pb.RegisterExampleServiceV1HandlerFromEndpoint(ctx, runtimeMux, s.ListenAddress, grpcDialOpts)
	//

	var g run.Group
	g.Add(func() error { return s.grpcSrv.Serve(grpcL) }, printErr)
	g.Add(func() error { return s.httpSrv.Serve(httpL) }, printErr)
	g.Add(func() error { return m.Serve() }, printErr)
	return g.Run()
}

func printErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func createRuntimeMux() *runtime.ServeMux {
	return runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			OrigName:     true,
			EmitDefaults: true,
		},
	))
}
