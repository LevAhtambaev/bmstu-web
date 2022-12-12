// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://vk.com/id250446192",
            "email": "fotchin@mail.ru"
        },
        "license": {
            "name": "AS IS (NO WARRANTY)"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comics/": {
            "get": {
                "description": "Get a list of all comics",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "Get all records",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.Comics"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ModelError"
                        }
                    }
                }
            },
            "post": {
                "description": "Adding a new comics to database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Add"
                ],
                "summary": "Add a new comics",
                "parameters": [
                    {
                        "description": "Название",
                        "name": "Name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Рейтинг",
                        "name": "Rate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Год производства",
                        "name": "Year",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Жанр",
                        "name": "Genre",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Цена",
                        "name": "Price",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Количество серий",
                        "name": "Episodes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Описание",
                        "name": "Description",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ModelComicsCreated"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ModelError"
                        }
                    }
                }
            }
        },
        "/comics/{uuid}": {
            "get": {
                "description": "Get a comics via uuid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "Get comics with corresponding name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID комикса",
                        "name": "UUID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ModelComicsDesc"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ModelError"
                        }
                    }
                }
            },
            "put": {
                "description": "Change a description of comics via its uuid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Change"
                ],
                "summary": "Change comics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID комикса",
                        "name": "UUID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Новое описание",
                        "name": "Description",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ModelDescChanged"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ModelError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a comics via its uuid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Change"
                ],
                "summary": "Delete a comics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID комикса",
                        "name": "UUID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ModelComicsDeleted"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ModelError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ds.Comics": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "rate": {
                    "type": "number"
                },
                "uuid": {
                    "type": "string"
                },
                "volumes": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.ModelComicsCreated": {
            "type": "object",
            "properties": {
                "success": {
                    "description": "success",
                    "type": "boolean"
                }
            }
        },
        "models.ModelComicsDeleted": {
            "type": "object",
            "properties": {
                "delete": {
                    "type": "boolean"
                }
            }
        },
        "models.ModelComicsDesc": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                }
            }
        },
        "models.ModelDescChanged": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "boolean"
                }
            }
        },
        "models.ModelError": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "description",
                    "type": "string"
                },
                "error": {
                    "description": "error",
                    "type": "string"
                },
                "type": {
                    "description": "type",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Comics Store",
	Description:      "Comics store",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
