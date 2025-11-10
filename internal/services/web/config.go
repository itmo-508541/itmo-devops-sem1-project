package web

import "fmt"

type Config struct {
	Host string
	Port string
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
