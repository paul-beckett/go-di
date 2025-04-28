package example4

import "strings"

type DBClient interface {
	List() (string, error)
}

type BasicDBClient struct {
	data []string
}

func (d *BasicDBClient) List() (string, error) {
	return strings.Join(d.data, ","), nil
}

type Service struct {
	client DBClient
}

func NewService(client DBClient) *Service {
	return &Service{client: client}
}

func (s *Service) Function1() (string, error) {
	return s.Function2()
}

func (s *Service) Function2() (string, error) {
	return s.Function3()
}

func (s *Service) Function3() (string, error) {
	return s.client.List()
}
