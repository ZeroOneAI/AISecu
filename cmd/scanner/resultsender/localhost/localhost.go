package localhost

import (
	"net/http"
	"os"
)

type Localhost struct {
	port string
}

func NewLocalhost(port string) *Localhost { return &Localhost{port: port} }

func (l *Localhost) Send(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	client := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:" + l.port + "/scan/result", file)
	if err != nil {
		return err
	}
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
