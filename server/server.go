package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	pb "github.com/alileza/example/autogen/pb"
	"github.com/alileza/example/services"
	"github.com/alileza/example/version"
)

type Server struct {
	ListenAddress       string
	Logger              *logrus.Logger
	UIDirectoryPath     string
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
	mux.Handle("/", uiMux(s.UIDirectoryPath))
	mux.Handle("/api/", runtimeMux)
	mux.Handle("/api/docs", docsHandler(s.SwaggerDocsFilePath))
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/version", version.Handler())

	unaryMiddlewares, streamMiddlewares := initiateGRPCMiddlewares(s.Logger)
	s.grpcSrv = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(unaryMiddlewares...),
		grpc_middleware.WithStreamServerChain(streamMiddlewares...),
	)
	grpc_prometheus.Register(s.grpcSrv)

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

	s.Logger.Infof("Start serving on address %s", s.ListenAddress)
	s.Logger.Infof("UI directory path %s", s.UIDirectoryPath)
	s.Logger.Infof("Docs file path %s", s.SwaggerDocsFilePath)
	return g.Run()
}

func initiateGRPCMiddlewares(logger *logrus.Logger) ([]grpc.UnaryServerInterceptor, []grpc.StreamServerInterceptor) {
	logrusEntry := logrus.NewEntry(logger)
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	grpc_prometheus.EnableHandlingTimeHistogram()

	limiter := &alwaysPassLimiter{}
	return []grpc.UnaryServerInterceptor{
			grpc_prometheus.UnaryServerInterceptor,
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(exampleAuthFunc),
			ratelimit.UnaryServerInterceptor(limiter),
		}, []grpc.StreamServerInterceptor{
			grpc_prometheus.StreamServerInterceptor,
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
			grpc_recovery.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(exampleAuthFunc),
			ratelimit.StreamServerInterceptor(limiter),
		}
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
