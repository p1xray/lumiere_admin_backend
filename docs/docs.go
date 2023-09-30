// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/cinema": {
            "get": {
                "description": "Получить список кинотеатров",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cinema"
                ],
                "summary": "Получить список кинотеатров",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DataResponse-CinemaList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Создать новый кинотеатр",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cinema"
                ],
                "summary": "Создать новый кинотеатр",
                "parameters": [
                    {
                        "description": "Входные параметры для создания кинотеатра",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CinemaInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DataResponse-bool"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/api/v1/cinema/{id}": {
            "get": {
                "description": "Получить подробности кинотеатра по идентификатору",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cinema"
                ],
                "summary": "Получить подробности кинотеатра",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор кинотеатра",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DataResponse-Cinema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Редактировать кинотеатр по идентификатору",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cinema"
                ],
                "summary": "Редактировать кинотеатр",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор кинотеатра",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Входные параметры для редактирования кинотеатра",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CinemaInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DataResponse-bool"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удалить кинотеатр по идентификатору",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cinema"
                ],
                "summary": "Удалить кинотеатр",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор кинотеатра",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DataResponse-bool"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Cinema": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "CinemaInput": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "CinemaList": {
            "type": "object",
            "properties": {
                "cinemas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Cinema"
                    }
                }
            }
        },
        "DataResponse-Cinema": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/Cinema"
                },
                "is_success": {
                    "type": "boolean"
                }
            }
        },
        "DataResponse-CinemaList": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/CinemaList"
                },
                "is_success": {
                    "type": "boolean"
                }
            }
        },
        "DataResponse-bool": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "boolean"
                },
                "is_success": {
                    "type": "boolean"
                }
            }
        },
        "Response": {
            "type": "object",
            "properties": {
                "is_success": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Lumiere admin API",
	Description:      "API sever for Lumiere admin",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}