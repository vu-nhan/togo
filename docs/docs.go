// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/configurations": {
            "post": {
                "description": "Create New Configuration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Configuration"
                ],
                "summary": "Create New Configuration",
                "parameters": [
                    {
                        "description": "Information of CreateConfigurationRequest",
                        "name": "CreateConfigurationRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateConfigurationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateConfigurationResponse"
                        }
                    }
                }
            }
        },
        "/configurations/date": {
            "get": {
                "description": "Get Configuration By Date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Configuration"
                ],
                "summary": "Get Configuration By Date",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Current Date",
                        "name": "current_date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.GetListTaskResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "get": {
                "description": "Get Authentication Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Authentication Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.TokenResponse"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Get List Tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get List Tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Created Date",
                        "name": "created_date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.GetListTaskResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Get New Task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Create New Task",
                "parameters": [
                    {
                        "description": "Information of CreateTaskRequest",
                        "name": "CreateTaskRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateTaskResponse"
                        }
                    }
                }
            }
        },
        "/users/": {
            "get": {
                "description": "Get All Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get All Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ApiResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get By Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                }
            }
        },
        "dtos.ConfigurationDto": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateConfigurationRequest": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateConfigurationResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/dtos.ConfigurationDto"
                }
            }
        },
        "dtos.CreateTaskRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateTaskResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/dtos.TaskDto"
                }
            }
        },
        "dtos.GetListTaskResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.TaskDto"
                    }
                }
            }
        },
        "dtos.TaskDto": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dtos.TokenResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
