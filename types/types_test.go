package types

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/ghodss/yaml"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestCheServiceYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-cheservice.yaml")
	check(err)

	expected := CheService{
		TypeMeta: TypeMeta{APIVersion: "v1", Kind: "CheService"},
		ObjectMeta: ObjectMeta{
			Name: "io.typefox.theia-ide.che-service",
		},
		Spec: CheServiceSpec{
			Version: "1.2.0",
			Containers: []Container{
				{
					Image: "eclipse/che-theia:nightly",
					Env: []EnvVar{
						{Name: "THEIA_PLUGINS", Value: "${THEIA_PLUGINS}"},
					},
					Resources: ResourceRequirements{
						Requests: ResourceList{"memory": "200Mi"},
					},
					Commands: []Command{
						{
							Name:       "build",
							WorkingDir: "$(project)",
							Command:    []string{"mvn", "clean", "install"},
						},
					},
					Servers: []Server{
						{
							Name:       "theia",
							Port:       3000,
							Protocol:   "http",
							Attributes: map[string]string{"internal": "true", "type": "ide"},
						},
					},
					Volumes: []Volume{
						{Name: "projects", MountPath: "/projects"},
					},
				},
			},
		},
	}

	var object CheService
	err = yaml.Unmarshal(dat, &object)
	check(err)
	// now we know that object isn't nil, we are safe to make
	// further assertions without causing any errors
	assert.Equal(t, expected, object, "Checking CheServiceParsing")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func print(data interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		panic(err)
	}
}

//
