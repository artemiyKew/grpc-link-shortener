package db

import (
	"errors"
	"sync"

	"github.com/artemiyKew/grpc-link-shortener/internal/entity"
)

type LinkMem struct {
	mu    *sync.Mutex
	links []*entity.LinkModel
}

func NewMemDB() *LinkMem {
	return &LinkMem{
		mu:    &sync.Mutex{},
		links: make([]*entity.LinkModel, 0, 32),
	}
}

func (m *LinkMem) Add(link *entity.LinkModel) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.links = append(m.links, link)
}

func (m *LinkMem) GetShortLink(shortURL string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.links {
		if l.ShortURL == shortURL {
			l.VisitsCount += 1
			return l.FullURL, nil
		}
	}
	return "", errors.New("cannot found short link")
}

func (m *LinkMem) GetList() []*entity.LinkModel {
	m.mu.Lock()
	defer m.mu.Unlock()
	newLinks := make([]*entity.LinkModel, len(m.links))
	copy(newLinks, m.links)
	return newLinks
}

func (m *LinkMem) CheckShortLinkExist(fullURL string) (*entity.LinkModel, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.links {
		if l.FullURL == fullURL {
			return l, true
		}
	}
	return &entity.LinkModel{}, false
}
