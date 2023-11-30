package service

import (
	"context"

	gatewayapi "github.com/artemiyKew/grpc-link-shortener/pkg/proto/link"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetShortLink(ctx context.Context, req *gatewayapi.GetShortLinkRequest) (*gatewayapi.GetShortLinkResponse, error) {
	fullURL, err := s.linkMem.GetShortLink(req.GetShortUrl())
	if err != nil {
		return &gatewayapi.GetShortLinkResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &gatewayapi.GetShortLinkResponse{
		FullUrl: fullURL,
	}, nil
}
