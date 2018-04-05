package server

import (
	"testing"
	"github.com/appleboy/gofight"
	"net/http"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"github.com/ghodss/yaml"
	"log"
	"fmt"


	chev1 "github.com/skabashnyuk/kubsrv/pkg/api/che.eclipse.org/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

func TestGinHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		SetDebug(true).
		Run(Setup(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		assert.Equal(t, http.StatusOK, r.Code)

		personMap := make(map[string]string)

		err := json.Unmarshal([]byte(r.Body.String()), &personMap)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, "http://localhost:8080/feature/{name}/{version}", personMap["feature_url"])

	})
}

func TestYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-service.yaml")
	check(err)
	var cheService chev1.CheService;
	err = yaml.Unmarshal(dat, &cheService)
	check(err)
}

func TestJson(t *testing.T) {
	dat, err := ioutil.ReadFile("test-service.json")
	check(err)
	var cheService chev1.CheService
	err = json.Unmarshal(dat, &cheService)
	check(err)
}

func TestJson2(t *testing.T) {
	//dat, err := os.Open("test-service.yaml")
	//check(err)

	dat, err := ioutil.ReadFile("test-service.json")
	check(err)

	decode := codecs.UniversalDeserializer().Decode
	obj, groupVersionKind, err := decode([]byte(dat), nil, nil)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
	}



	fmt.Sprintf("%s", obj)
	fmt.Sprintf("%s", groupVersionKind)


//	var cheService v1.CheService

	//json := k8syaml.NewYAMLOrJSONDecoder(dat,4096)
	//
	//err = json.Decode(&cheService)
	//check(err)


	//
	//d := scheme.Codecs.UniversalDecoder()
	//obj, _, err := d.Decode([]byte(dat), nil, nil)
	//if err != nil {
	//	log.Fatalf("could not decode yaml: %s\n%s", dat, err)
	//}
	//
	//fmt.Println(obj)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}



var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	AddToScheme(scheme)
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
func AddToScheme(scheme *runtime.Scheme) {
	chev1.AddToScheme(scheme)
}
