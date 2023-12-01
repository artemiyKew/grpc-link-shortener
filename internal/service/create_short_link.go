package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"
	"sync"

	"github.com/artemiyKew/grpc-link-shortener/internal/entity"
	gatewayapi "github.com/artemiyKew/grpc-link-shortener/pkg/proto/link"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO service-delivery
func (s *Service) CreateShortLink(ctx context.Context, req *gatewayapi.CreateShortLinkRequest) (*gatewayapi.CreateShortLinkResponse, error) {
	link, ok := s.linkMem.CheckShortLinkExist(req.FullUrl)
	if ok {
		return &gatewayapi.CreateShortLinkResponse{
			FullUrl:     link.FullURL,
			ShortUrl:    link.ShortURL,
			VisitsCount: int32(link.VisitsCount),
		}, nil
	}

	link = &entity.LinkModel{
		FullURL:     s.validateAndFixURL(req.FullUrl),
		ShortURL:    generateShortLink(req.GetFullUrl(), 10),
		VisitsCount: 0,
	}
	_, err := url.ParseRequestURI(link.FullURL)
	if err != nil {
		return &gatewayapi.CreateShortLinkResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	s.linkMem.Add(link)

	return &gatewayapi.CreateShortLinkResponse{
		FullUrl:     link.FullURL,
		ShortUrl:    link.ShortURL,
		VisitsCount: int32(link.VisitsCount),
	}, nil
}

func generateShortLink(inputURL string, tokenLength int) string {
	var tokenMutex sync.Mutex
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	hash := sha256.Sum256([]byte(inputURL))

	shortenedURL := hex.EncodeToString(hash[:5])

	return shortenedURL
}

func (s *Service) validateAndFixURL(url string) string {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}
	return url
}
