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
	desc.UnimplementedLinksServiceServer
	services   *service.Service
	grpcServer *grpc.Server
}

func NewGrpcServer(service *service.Service) *Server {
	return &Server{
		services:   service,
		grpcServer: grpc.NewServer(),
	}
}

func (s *Server) Run(port string, services *service.Service) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	s = &Server{
		services:   services,
		grpcServer: grpc.NewServer(),
	}
	desc.RegisterLinksServiceServer(s.grpcServer, s)
	return s.grpcServer.Serve(lis)
}

func (s *Server) CreateShortUrl(ctx context.Context, req *desc.CreateShortUrlRequest) (*desc.CreateShortUrlResponse, error) {
	shortUrl, err := s.services.LinksShortener.CreateShortUrl(req.Url)
	if err != nil {
		log.Error("%v", err)
		return &desc.CreateShortUrlResponse{Url: "{}"}, err
	}
	return &desc.CreateShortUrlResponse{Url: shortUrl}, err
}

func (s *Server) GetLongUrl(ctx context.Context, req *desc.GetLongUrlRequest) (*desc.GetLongUrlResponse, error) {
	longUrl, err := s.services.LinksShortener.GetLongUrl(req.Url)
	if err != nil {
		log.Error("%v", err)
		return &desc.GetLongUrlResponse{Url: "{}"}, err
	}
	return &desc.GetLongUrlResponse{Url: longUrl}, err
}

func (s *Server) ShuttingDown() {
	s.grpcServer.GracefulStop()
}
