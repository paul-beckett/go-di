package example2

import "strings"

type DBClient struct {
	data []string
}

func (d *DBClient) List() (string, error) {
	return strings.Join(d.data, ","), nil
}

func Function1(client *DBClient) (string, error) {
	return Function2(client)
}

func Function2(client *DBClient) (string, error) {
	return Function3(client)
}

func Function3(client *DBClient) (string, error) {
	return client.List()
}
