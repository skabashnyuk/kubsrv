## What is it
Prototype of Che Registry implementation. It is used to serve CheFeatures, CheServices and ChePlugins.
Based on this api https://app..swaggerhub.com/apis/skabashniuk/Che/1.0.0 For now only ```GET``` method has been implemented.
```
[GIN-debug] GET    /                         --> github.com/skabashnyuk/kubsrv/controller.APIEndpoints (3 handlers)
[GIN-debug] GET    /plugin/:name/:version    --> github.com/skabashnyuk/kubsrv/controller.(*Plugin).GetPlugin-fm (3 handlers)
[GIN-debug] GET    /plugin/                  --> github.com/skabashnyuk/kubsrv/controller.(*Plugin).GetLatestPluginsList-fm (3 handlers)
[GIN-debug] GET    /service/:name/:version   --> github.com/skabashnyuk/kubsrv/controller.(*Service).GetService-fm (3 handlers)
[GIN-debug] GET    /service                  --> github.com/skabashnyuk/kubsrv/controller.(*Service).GetServiceByIdList-fm (3 handlers)
[GIN-debug] GET    /feature/:name/:version   --> github.com/skabashnyuk/kubsrv/controller.(*Feature).GetFeature-fm (3 handlers)
[GIN-debug] GET    /feature                  --> github.com/skabashnyuk/kubsrv/controller.(*Feature).GetFeatureByIdList-fm (3 handlers)

```


## How to run

1. Run ```docker pull ksmster/kubsrv``` Pull latest docker image. ```ksmster/kubsrv:latest``` are build on each commit.
   ```ksmster/kubsrv:tagname```  are build on each git tag.
2. Run ```docker run -it  -p 3000:3000 ksmster/kubsrv``` server with default configuration. 


## How to configure executable
| Parameter name  | Description   | Default value | 
|---|---|---|
| github   | Git url of repository to clone  | n/a  |
| registry  | Location of repository on filesystem   | n/a  |
| update  | Storage update interval in seconds. Set -1 to disable  |  n/a  |

Example ```kubsrv_Linux_x86_64 -github https://github.com/skabashnyuk/che-registry.git -registry /kubsrv/repo   -update 60```



## How to configure ksmster/kubsrv docker image
| Parameter name  | Description   | Default value | 
|---|---|---|
| CHE_REGISTRY_GITHUB_URL   | Git url of repository to clone  | https://github.com/skabashnyuk/che-registry.git  |
| CHE_REGISTRY_REPOSITORY  | Location of repository on filesystem   | /kubsrv/repo  |
| CHE_REGISTRY_UPDATE_INTERVAL  | Storage update interval in seconds. Set -1 to disable  | 60  |

Example ```docker run -it -v /home/user/mylocalrepo:/kubsrv/repo -e CHE_REGISTRY_UPDATE_INTERVAL=0 -p 3000:3000 ksmster/kubsrv ```


## How to use
1. Run ```docker run -it  -p 3000:3000 ksmster/kubsrv``` to start Che registry. Registry content will be taken from here https://github.com/skabashnyuk/che-registry.git
2. Run  ``` curl  http://localhost:3000/plugin/``` To get latest version of all plugins.
```
[
  {
    "name": "che-theia-github",
    "version": "0.0.1",
    "title": "SSH Client",
    "created_at": "2015-11-22T05:40:57Z",
    "updated_at": "2018-04-22T00:27:28Z",
    "installation_count": 88,
    "description": "Eclipse Che Plugin for Github",
    "owner": {
      "name": "redhat",
      "url": "https://redhat.com"
    }
  },
  {
    "name": "che-theia-ssh",
    "version": "0.0.1",
    "title": "SSH Client",
    "created_at": "2015-11-22T05:40:57Z",
    "updated_at": "2018-04-22T00:27:28Z",
    "installation_count": 88,
    "description": "Eclipse Che Plugin for SSH",
    "owner": {
      "name": "redhat",
      "url": "https://redhat.com"
    }
  },
  {
    "name": "che-theia-ide",
    "version": "0.0.1",
    "title": "Theia IDE service",
    "created_at": "2015-11-22T05:40:57Z",
    "updated_at": "2018-04-22T00:27:28Z",
    "installation_count": 88,
    "description": "Eclipse Che Theia IDE",
    "owner": {
      "name": "redhat",
      "url": "https://redhat.com"
    }
  }
]

```
3. Run  ```curl  http://localhost:3000/plugin/org.eclipse.che.che-theia-ssh/0.0.1``` To get concrete plugin.
```
{
  "name": "che-theia-ssh",
  "version": "0.0.1",
  "title": "SSH Client",
  "created_at": "2015-11-22T05:40:57Z",
  "updated_at": "2018-04-22T00:27:28Z",
  "installation_count": 88,
  "description": "Eclipse Che Plugin for SSH",
  "owner": {
    "name": "redhat",
    "url": "https://redhat.com"
  }
}

```
4. Run  ```curl  http://localhost:3000/feature/org.eclipse.che.che-theia-ssh/0.0.1``` To get concrete feature.
```
{
  "apiVersion": "v1",
  "kind": "CheFeature",
  "metadata": {
    "name": "che-theia-ssh"
  },
  "spec": {
    "version": "0.0.1",
    "services": [
      {
        "name": "org.eclipse.che.theia-ide",
        "version": "0.0.1",
        "parameters": [
          {
            "name": "THEIA_PLUGINS",
            "value": "eclipse-che-ssh-extension.tar.gz"
          }
        ]
      }
    ]
  }
}

```
5. Run ```curl  "http://localhost:3000/feature?id=org.eclipse.che.che-theia-ssh:0.0.1&id=org.eclipse.che.che-theia-github:0.0.1"``` 
   To get features by ids
```
[
  {
    "apiVersion": "v1",
    "kind": "CheFeature",
    "metadata": {
      "name": "che-theia-ssh"
    },
    "spec": {
      "version": "0.0.1",
      "services": [
        {
          "name": "org.eclipse.che.theia-ide",
          "version": "0.0.1",
          "parameters": [
            {
              "name": "THEIA_PLUGINS",
              "value": "eclipse-che-ssh-extension.tar.gz"
            }
          ]
        }
      ]
    }
  },
  {
    "apiVersion": "v1",
    "kind": "CheFeature",
    "metadata": {
      "name": "che-theia-github"
    },
    "spec": {
      "version": "0.0.1",
      "services": [
        {
          "name": "org.eclipse.che.theia-ide",
          "version": "0.0.1",
          "parameters": [
            {
              "name": "THEIA_PLUGINS",
              "value": "che-theia-github.tar.gz"
            }
          ]
        }
      ]
    }
  }
]

```
6.  Run ```curl  http://localhost:3000/service/org.eclipse.che.theia-ide/0.0.1``` To get concrete CheService
```
{
  "apiVersion": "v1",
  "kind": "CheService",
  "metadata": {
    "name": "io.typefox.theia-ide.che-service"
  },
  "spec": {
    "version": "0.0.1",
    "containers": [
      {
        "image": "eclipse/che-theia:nightly",
        "env": [
          {
            "name": "THEIA_PLUGINS",
            "value": "${THEIA_PLUGINS}"
          }
        ],
        "resources": {
          "requests": {
            "memory": "200Mi"
          }
        },
        "servers": [
          {
            "name": "theia",
            "port": 3000,
            "protocol": "http",
            "attributes": {
              "internal": "true",
              "type": "ide"
            }
          }
        ],
        "volumes": [
          {
            "name": "projects",
            "mountPath": "/projects"
          }
        ]
      }
    ]
  }
}
```
7. Run ```curl  http://localhost:3000/service?id=org.eclipse.che.theia-ide:0.0.1``` To get services by ids
```
[
  {
    "apiVersion": "v1",
    "kind": "CheService",
    "metadata": {
      "name": "io.typefox.theia-ide.che-service"
    },
    "spec": {
      "version": "0.0.1",
      "containers": [
        {
          "image": "eclipse/che-theia:nightly",
          "env": [
            {
              "name": "THEIA_PLUGINS",
              "value": "${THEIA_PLUGINS}"
            }
          ],
          "resources": {
            "requests": {
              "memory": "200Mi"
            }
          },
          "servers": [
            {
              "name": "theia",
              "port": 3000,
              "protocol": "http",
              "attributes": {
                "internal": "true",
                "type": "ide"
              }
            }
          ],
          "volumes": [
            {
              "name": "projects",
              "mountPath": "/projects"
            }
          ]
        }
      ]
    }
  }
]
```


## How to add new ChePlugin
1. Add CheMata.yaml to folder like this ```https://github.com/skabashnyuk/che-registry/tree/master/org/eclipse/che/che-theia-github/0.0.1```
   Where path consructed with name and version and dots in name replaces with slashes.
2. CheFeature.yaml in same folder   
3. Add Theia plugin in same folder
4. Add CheService in same folder if needed. Like this ```https://github.com/skabashnyuk/che-registry/blob/master/org/eclipse/che/theia-ide/0.0.1/CheService.yaml```    


## TODO
 - [ ] Return latest version of all plugins. Now all plugins will be returned
 - [ ] Plugins pagination.