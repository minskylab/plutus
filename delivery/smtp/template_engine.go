package smtp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

type TemplateEngine interface {
	GetTemplate(pathname ...string) (string, error)
}

type FileSystemEngine struct {
	absolutePath string
}

type NetworkEngine struct {
	baseAPI string
}

func (engine *FileSystemEngine) GetTemplate(pathname ...string) (string, error) {
	finalPath := path.Join(engine.absolutePath)
	for _, p := range pathname {
		finalPath = path.Join(finalPath, p)
	}

	data, err := ioutil.ReadFile(finalPath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (engine *NetworkEngine) GetTemplate(pathname ...string) (string, error) {
	finalURI := path.Join(engine.baseAPI)
	for _, p := range pathname {
		if strings.HasPrefix(p, "?") {
			finalURI = finalURI + p
		} else {
			finalURI = path.Join(finalURI, p)
		}
	}
	response, err := http.Get(finalURI)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	result := map[string]interface{}{}

	result["code"] = response.StatusCode
	result["body"] = string(data)

	finalResponse, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(finalResponse), nil
}
