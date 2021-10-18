// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users/{userId}/posts": {
            "get": {
                "description": "get posts by userId",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "getUserPosts",
                "operationId": "getUserPosts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Post"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    }
                }
            },
            "post": {
                "description": "createPost",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "createPost",
                "operationId": "createPost",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "post",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/api/v1/users/{userId}/posts/{postId}"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userId}/posts/{postId}": {
            "get": {
                "description": "getUserPostById",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "getUserPostById",
                "operationId": "getUserPostById",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "postId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    }
                }
            },
            "put": {
                "description": "updatePost",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "updatePost",
                "operationId": "updatePost",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "postId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "post",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    }
                }
            },
            "delete": {
                "description": "deletePost",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "deletePost",
                "operationId": "deletePost",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "postId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "description": "signIn and get access and refresh tokens",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "signIn",
                "operationId": "signIn",
                "parameters": [
                    {
                        "description": "body",
                        "name": "signInUserInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.signInUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.signInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "signUp new user",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "signUp",
                "operationId": "signUp",
                "parameters": [
                    {
                        "description": "a body",
                        "name": "signUpUserInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.signUpUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Post": {
            "type": "object",
            "required": [
                "body",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "model.UpdatePost": {
            "type": "object",
            "required": [
                "body",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "server.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "server.signInResponse": {
            "type": "object",
            "properties": {
                "access-token": {
                    "type": "string"
                },
                "refresh-token": {
                    "type": "string"
                }
            }
        },
        "server.signInUserInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "server.signUpUserInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "NIX_EDU_APP",
	Description: "This is a backend for nix tasks https://education.nixsolutions.com/mod/page/view.php?id=79",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
