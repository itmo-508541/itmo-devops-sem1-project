package settings

import "fmt"

type WebSettings struct {
	Host string
	Port string
}

func (c WebSettings) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
