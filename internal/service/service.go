package service

import (
	"github.com/artemiyKew/grpc-link-shortener/internal/db"
	gatewayapi "github.com/artemiyKew/grpc-link-shortener/pkg/proto/link"
)

type Service struct {
	gatewayapi.UnimplementedLinkShortenerServer

	linkMem *db.LinkMem
}

func New() *Service {
	return &Service{
		linkMem: db.NewMemDB(),
	}
}
