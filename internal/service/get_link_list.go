package service

import (
	"context"

	gatewayapi "github.com/artemiyKew/grpc-link-shortener/pkg/proto/link"
)

func (s *Service) GetLinksList(ctx context.Context, req *gatewayapi.GetLinksListRequest) (*gatewayapi.GetLinksListResponse, error) {
	memLinks := s.linkMem.GetList()
	apiLinks := make([]*gatewayapi.Link, 0)

	for _, link := range memLinks {
		apiLinks = append(apiLinks, &gatewayapi.Link{
			FullUrl:     link.FullURL,
			ShortUrl:    link.ShortURL,
			VisitsCount: int32(link.VisitsCount),
		})

	}
	return &gatewayapi.GetLinksListResponse{
		Result: apiLinks,
	}, nil
}
