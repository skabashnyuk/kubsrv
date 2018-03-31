package render


import (
	"net/http"

	"github.com/ghodss/yaml"
)

type GYAML struct {
	Data interface{}
}

var yamlContentType = []string{"application/x-yaml; charset=utf-8"}

func (r GYAML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := yaml.Marshal(r.Data)
	if err != nil {
		return err
	}

	w.Write(bytes)
	return nil
}

func (r GYAML) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = yamlContentType
	}
}
