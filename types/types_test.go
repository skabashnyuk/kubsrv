package types

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/ghodss/yaml"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-cheservice.yaml")
	check(err)
	var cheService CheService
	err = yaml.Unmarshal(dat, &cheService)
	check(err)
}

func TestCheServiceYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-cheservice.yaml")
	check(err)
	expected := CheService{
		TypeMeta: TypeMeta{
			APIVersion: "che.eclipse.org/v1",
			Kind:       "CheService"},
		ObjectMeta: ObjectMeta{
			Name:   "io.typefox.theia-ide.che-service",
			Labels: map[string]string(nil)},
		Spec: CheServiceSpec{
			Services: []Service{
				{
					TypeMeta: TypeMeta{
						APIVersion: "v1",
						Kind:       "Service"},
					ObjectMeta: ObjectMeta{
						Name:   "io.typefox.theia-ide.che-service.mainservice",
						Labels: map[string]string(nil)},
					Spec: ServiceSpec{
						Ports: []ServicePort{
							{
								Name:       "",
								Protocol:   "TCP",
								Port:       80,
								TargetPort: 9376,
							},
						},
						Selector: map[string]string{"app": "theia"},
					},
				},
			},
			Pods: []Pod{
				{
					TypeMeta: TypeMeta{
						APIVersion: "v1",
						Kind:       "Pod"},
					ObjectMeta: ObjectMeta{
						Name:   "io.typefox.theia-ide.che-service.mainpod",
						Labels: map[string]string(nil),
					},
					Spec: PodSpec{
						Volumes: []Volume(nil),
						Containers: []Container{
							{
								Name:  "theia",
								Image: "eclipse/che-theia:nightly",
								Env: []EnvVar{
									{
										Name:  "THEIA_PLUGINS",
										Value: "${THEIA_PLUGINS}",
									},
								},
								VolumeMounts: []VolumeMount{
									{
										Name:      "projects",
										ReadOnly:  false,
										MountPath: "/projects",
										SubPath:   "",
									},
								},
							},
							{
								Name:         "theia-sidecar",
								Image:        "maven:latest",
								Env:          []EnvVar(nil),
								VolumeMounts: []VolumeMount(nil),
							},
						},
					},
				},
			},
			Commands: []CheCommand{
				{
					TypeMeta: TypeMeta{
						APIVersion: "che.eclipse.org/v1",
						Kind:       "CheCommand",
					},
					ObjectMeta: ObjectMeta{
						Name:   "mvn-clean",
						Labels: map[string]string(nil),
					},
					Spec: CheCommandSpec{
						TargetLabelSelector: "che.eclipse.org/container-name= theia-maven-sidecar",
						WorkingDirectory:    "$(project)",
						Commands: []string{"mvn", "clean", "install",
						},
					},
				},
			},
			Version: Version{
				Version: "1.2.2",
			},
		},
	}
	var object CheService
	err = yaml.Unmarshal(dat, &object)
	check(err)
	print(&object)
	// now we know that object isn't nil, we are safe to make
	// further assertions without causing any errors
	assert.Equal(t, expected, object, "Checking CheServiceParsing")
}

func TestServiceArrayYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-arrayofservice.yaml")
	check(err)
	var object TestStruct
	err = yaml.Unmarshal(dat, &object)
	check(err)
}

func TestServiceArrayJson(t *testing.T) {
	dat, err := ioutil.ReadFile("test-arrayofservices.json")
	check(err)
	var object TestStruct
	err = json.Unmarshal(dat, &object)
	check(err)
	print(&object)
}

func TestJson(t *testing.T) {
	dat, err := ioutil.ReadFile("test-cheservice.json")
	check(err)
	var object CheService
	err = json.Unmarshal(dat, &object)
	check(err)
	print(&object)
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

type TestStruct struct {
	Services []Service `json:"services,omitempty"`
}
