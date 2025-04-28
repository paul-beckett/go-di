package example3

import "strings"

type DBClient struct {
	data []string
}

func (d *DBClient) List() (string, error) {
	return strings.Join(d.data, ","), nil
}

type Service struct {
	client *DBClient
}

func NewService(client *DBClient) *Service {
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
