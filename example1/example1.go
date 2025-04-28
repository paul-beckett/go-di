package example1

import "strings"

type DBClient struct {
	data []string
}

func (d *DBClient) List() (string, error) {
	return strings.Join(d.data, ","), nil
}

func Function1() (string, error) {
	return Function2()
}

func Function2() (string, error) {
	return Function3()
}

func Function3() (string, error) {
	data := []string{"a", "b", "c"}
	client := DBClient{data: data}
	return client.List()
}
