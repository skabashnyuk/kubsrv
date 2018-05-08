package storage

import (
	"reflect"
	"testing"
	"github.com/skabashnyuk/kubsrv/types"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestStorage_GetPlugins(t *testing.T) {
	type fields struct {
		CheRegistryRepository string
		CheRegistryGithubUrl  string
	}
	type args struct {
		Limit  int
		Offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]types.ChePlugin
		wantErr bool
	}{
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				CheRegistryRepository: tt.fields.CheRegistryRepository,
				CheRegistryGithubUrl:  tt.fields.CheRegistryGithubUrl,
			}
			got, err := storage.GetPlugins(tt.args.Limit, tt.args.Offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetPlugins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetPlugins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetPlugin(t *testing.T) {

	type fields struct {
		CheRegistryRepository string
		CheRegistryGithubUrl  string
	}
	type args struct {
		Id *ItemId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.ChePlugin
		wantErr bool
	}{
		{
			name: "Get PluginByID",
			fields: fields{
				CheRegistryGithubUrl:  "",
				CheRegistryRepository: "testdata",
			},
			args: args{
				Id: &ItemId{
					Name:    "org.eclipse.che.theia-ide",
					Version: "0.0.1",
				},
			},
			want: &types.ChePlugin{
				Name:              "che-theia-ide",
				Version:           "0.0.1",
				Title:             "Theia IDE service",
				CreatedAt:         "2015-11-22T05:40:57Z",
				UpdatedAt:         "2018-04-22T00:27:28Z",
				InstallationCount: 88,
				Description:       "Eclipse Che Theia IDE",
				Licese: &types.PluginLicense{
					Key:  "mit",
					Name: "MIT License",
					Url:  "https://api.github.com/licenses/mit",
				},
				Owner: &types.PluginOwner{
					Name:       "redhat",
					AvatarUrl:  "",
					GravatarId: "",
					Url:        "https://redhat.com",
				}},
			wantErr: false,
		},
		{
			name: "Get Github",
			fields: fields{
				CheRegistryGithubUrl:  "",
				CheRegistryRepository: "testdata",
			},
			args: args{
				Id: &ItemId{
					Name:    "org.eclipse.che.che-theia-github",
					Version: "0.0.1",
				},
			},
			want: &types.ChePlugin{
				Name:              "che-theia-github",
				Version:           "0.0.1",
				Title:             "Github Client",
				CreatedAt:         "2015-11-22T05:40:57Z",
				UpdatedAt:         "2018-04-22T00:27:28Z",
				InstallationCount: 88,
				Description:       "Eclipse Che Plugin for Github",
				Licese: &types.PluginLicense{
					Key:  "mit",
					Name: "MIT License",
					Url:  "https://api.github.com/licenses/mit",
				},
				Owner: &types.PluginOwner{
					Name:       "redhat",
					AvatarUrl:  "",
					GravatarId: "",
					Url:        "https://redhat.com",
				}},
			wantErr: false,
		}, {
			name: "Get SSH plugin",
			fields: fields{
				CheRegistryGithubUrl:  "",
				CheRegistryRepository: "testdata",
			},
			args: args{
				Id: &ItemId{
					Name:    "org.eclipse.che.che-theia-ssh",
					Version: "0.0.1",
				},
			},
			want: &types.ChePlugin{
				Name:              "che-theia-ssh",
				Version:           "0.0.1",
				Title:             "SSH Client",
				CreatedAt:         "2015-11-22T05:40:57Z",
				UpdatedAt:         "2018-04-22T00:27:28Z",
				InstallationCount: 88,
				Description:       "Eclipse Che Plugin for SSH",
				Licese: &types.PluginLicense{
					Key:  "mit",
					Name: "MIT License",
					Url:  "https://api.github.com/licenses/mit",
				},
				Owner: &types.PluginOwner{
					Name:       "redhat",
					AvatarUrl:  "",
					GravatarId: "",
					Url:        "https://redhat.com",
				}},
			wantErr: false,
		}, {
			name: "Should FAIL PluginByID with wrong name",
			fields: fields{
				CheRegistryGithubUrl:  "",
				CheRegistryRepository: "testdata",
			},
			args: args{
				Id: &ItemId{
					Name:    "org.eclipse.che.che-theia-some",
					Version: "0.0.1",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Should FAIL PluginByID with wrong version",
			fields: fields{
				CheRegistryGithubUrl:  "",
				CheRegistryRepository: "testdata",
			},
			args: args{
				Id: &ItemId{
					Name:    "org.eclipse.che.che-theia-ssh",
					Version: "0.0.2",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Should FAIL PluginByID with parse error",
			fields: fields{
				CheRegistryGithubUrl:  "",
				CheRegistryRepository: "testdata",
			},
			args: args{
				Id: &ItemId{
					Name:    "org.eclipse.che.invalid",
					Version: "0.0.1",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				CheRegistryRepository: tt.fields.CheRegistryRepository,
				CheRegistryGithubUrl:  tt.fields.CheRegistryGithubUrl,
			}
			got, err := storage.GetPlugin(tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetPlugin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStorage_GetCheService(t *testing.T) {
	type fields struct {
		CheRegistryRepository string
		CheRegistryGithubUrl  string
	}
	type args struct {
		Id *ItemId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.CheService
		wantErr bool
	}{
		{
			name: "Should GetService By ID",
			fields: fields{
				CheRegistryGithubUrl:  "",
				CheRegistryRepository: "testdata",
			},
			args: args{
				Id: &ItemId{
					Name:    "org.eclipse.che.theia-ide",
					Version: "0.0.1",
				},
			},
			want: &types.CheService{
				TypeMeta: types.TypeMeta{APIVersion: "v1", Kind: "CheService"},
				ObjectMeta: types.ObjectMeta{
					Name: "io.typefox.theia-ide.che-service",
				},
				Spec: types.CheServiceSpec{
					Version: "0.0.1",
					Containers: []types.Container{
						{
							Image: "eclipse/che-theia:nightly",
							Env: []types.EnvVar{
								{Name: "THEIA_PLUGINS", Value: "${THEIA_PLUGINS}"},
							},
							Resources: types.ResourceRequirements{
								Requests: types.ResourceList{"memory": "200Mi"},
							},
							Commands: []types.Command{
								{
									Name:       "build",
									WorkingDir: "$(project)",
									Command:    []string{"mvn", "clean", "install"},
								},
							},
							Servers: []types.Server{
								{
									Name:       "theia",
									Port:       3000,
									Protocol:   "http",
									Attributes: map[string]string{"internal": "true", "type": "ide"},
								},
							},
							Volumes: []types.Volume{
								{Name: "projects", MountPath: "/projects"},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				CheRegistryRepository: tt.fields.CheRegistryRepository,
				CheRegistryGithubUrl:  tt.fields.CheRegistryGithubUrl,
			}
			got, err := storage.GetCheService(tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetCheService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStorage_GetCheFeature(t *testing.T) {
	type fields struct {
		CheRegistryRepository string
		CheRegistryGithubUrl  string
	}
	type args struct {
		Id *ItemId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.CheFeature
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				CheRegistryRepository: tt.fields.CheRegistryRepository,
				CheRegistryGithubUrl:  tt.fields.CheRegistryGithubUrl,
			}
			got, err := storage.GetCheFeature(tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetCheFeature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetCheFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdateStorage(t *testing.T) {
	type fields struct {
		CheRegistryRepository string
		CheRegistryGithubUrl  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				CheRegistryRepository: tt.fields.CheRegistryRepository,
				CheRegistryGithubUrl:  tt.fields.CheRegistryGithubUrl,
			}
			storage.UpdateStorage()
		})
	}
}

func TestStorage_EnsureExists(t *testing.T) {
	type fields struct {
		CheRegistryRepository string
		CheRegistryGithubUrl  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				CheRegistryRepository: tt.fields.CheRegistryRepository,
				CheRegistryGithubUrl:  tt.fields.CheRegistryGithubUrl,
			}
			storage.EnsureExists()
		})
	}
}

func mustReadFile(t *testing.T, filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
