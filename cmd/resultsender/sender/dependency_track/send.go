package dependency_track

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type DependencyTrackBomSendBody struct {
	Project string `json:"project"`
	Bom     string `json:"bom"`
}

func (c *Controller) Send(sendData []byte, projectName string, projectVersion string) error {
	fmt.Println(string(sendData))
	endpoint := c.hostname + "/api/v1/bom"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("bom", "result.json")
	if err != nil {
		return err
	}
	writeCnt, err := part.Write(sendData)
	if err != nil {
		return err
	}
	if len(sendData) != writeCnt {
		return errors.New(fmt.Sprintf("copy not valid src len : [%d], dst len : [%d]\n", len(sendData), writeCnt))
	}
	writer.WriteField("autoCreate", "true")
	writer.WriteField("projectName", projectName)
	writer.WriteField("projectVersion", projectVersion)
	writer.Close()

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return err
	}
	req.Header.Add("X-API-Key", c.apiKey)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	fmt.Println(string(b), err)
	return nil
}
