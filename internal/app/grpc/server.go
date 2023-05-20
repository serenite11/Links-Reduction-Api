package grpc

import (
	"context"
	desc "github.com/serenite11/Links-Reduction-Api/internal/app/grpc/api"
	"github.com/serenite11/Links-Reduction-Api/internal/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	desc.UnimplementedAdderServer
	services   *service.Service
	grpcServer *grpc.Server
}

func NewGrpcServer(service *service.Service) *Server {
	s := &Server{
		services:   service,
		grpcServer: grpc.NewServer(),
	}
	desc.RegisterAdderServer(s.grpcServer, &Server{services: s.services})
	return s
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return s.grpcServer.Serve(lis)
}

func (s *Server) Add(ctx context.Context, req *desc.AddRequest) (*desc.AddResponse, error) {
	shortUrl, err := s.services.LinksShortener.CreateShortUrl(req.Url)
	if err != nil {
		log.Error("%v", err)
		return &desc.AddResponse{Url: "{}"}, err
	}
	return &desc.AddResponse{Url: shortUrl}, err
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	longUrl, err := s.services.LinksShortener.GetLongUrl(req.Url)
	if err != nil {
		log.Error("%v", err)
		return &desc.GetResponse{Url: "{}"}, err
	}
	return &desc.GetResponse{Url: longUrl}, err
}

func (s *Server) ShuttingDown() {
	s.grpcServer.GracefulStop()
}
