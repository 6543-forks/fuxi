// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-10 22:59:39.913315 +0800 CST m=+0.051515520

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "laik author",
            "url": "github.com/yametech",
            "email": "laik.lj@me.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/base/apis/fuxi.nip.io/v1/basedepartments": {
            "get": {
                "description": "base service for list base department",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseDepartment"
                ],
                "summary": "base department list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/baseroles": {
            "get": {
                "description": "base service for list base role",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseRole"
                ],
                "summary": "base role list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/baseroleusers": {
            "get": {
                "description": "base service for list base role user",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseRoleUser"
                ],
                "summary": "base role user list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/baseusers": {
            "get": {
                "description": "base service for list base role",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseUser"
                ],
                "summary": "base role list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/:namespace/baseroles/": {
            "post": {
                "description": "base service for base role",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseRole"
                ],
                "summary": "base role list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/:namespace/baseroleusers/": {
            "post": {
                "description": "base service for base role user",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseRoleUser"
                ],
                "summary": "base role user list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/{namespace}/basedepartments/": {
            "post": {
                "description": "base service for base department",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseDepartment"
                ],
                "summary": "base department list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/{namespace}/basedepartments/{name}": {
            "get": {
                "description": "base service for get a base department detail",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseDepartment"
                ],
                "summary": "base department get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseroles/{name}": {
            "get": {
                "description": "base service for get a base role detail",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseRole"
                ],
                "summary": "base role get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseroleusers/{name}": {
            "get": {
                "description": "base service for get a base roleuser detail",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseRoleUser"
                ],
                "summary": "base role user get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseusers/": {
            "post": {
                "description": "base service for base role",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseUser"
                ],
                "summary": "base role list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseusers/{name}": {
            "get": {
                "description": "base service for get a base role detail",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BaseUser"
                ],
                "summary": "base role get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/permission_auth_value": {
            "post": {
                "description": "base service for list base permission",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "base permission",
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/permission_list": {
            "get": {
                "description": "base service for list base permission",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "base permission",
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/permission_transfer/{value}": {
            "get": {
                "description": "base service for list base permission",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "base permission",
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"Failed\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Gin swagger",
	Description: "Gin swagger base",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
