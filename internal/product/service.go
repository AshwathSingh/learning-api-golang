package product

import (
	"context"
	"log"
)

type Service interface {
	ListProducts(ctx context.Context) (error)
}


type svc struct {
	// repository of data
}


func NewService() Service {
	return &svc{

	}
}


func (s *svc) ListProducts(ctx context.Context) (error) {

	err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
	}
	return nil
}
