// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/quotes": {
            "get": {
                "description": "Возвращает наиболее свежую котировку",
                "produces": [
                    "application/json"
                ],
                "summary": "Получение котировки по коду валют",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Код валюты, например EUR/MXN",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное получение котировки",
                        "schema": {
                            "$ref": "#/definitions/httpserver.quoteResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный формат входных данных",
                        "schema": {
                            "$ref": "#/definitions/httpserver.quoteResponse"
                        }
                    },
                    "500": {
                        "description": "Проблемы на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/httpserver.quoteResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Обновляет одну котировку и сохраняет её в БД. Возвращает id, по которому можно получить котировку.",
                "produces": [
                    "application/json"
                ],
                "summary": "Обновление котировки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Код валюты, например EUR/MXN",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "id котировки",
                        "schema": {
                            "$ref": "#/definitions/httpserver.idResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный формат входных данных",
                        "schema": {
                            "$ref": "#/definitions/httpserver.idResponse"
                        }
                    },
                    "500": {
                        "description": "Проблемы на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/httpserver.idResponse"
                        }
                    }
                }
            }
        },
        "/quotes/{id}": {
            "get": {
                "description": "Возвращает обновлённую котировку",
                "produces": [
                    "application/json"
                ],
                "summary": "Получение котировки по id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id котировки",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное получение котировки",
                        "schema": {
                            "$ref": "#/definitions/httpserver.quoteResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный формат входных данных",
                        "schema": {
                            "$ref": "#/definitions/httpserver.quoteResponse"
                        }
                    },
                    "404": {
                        "description": "Котировки с указанным id не существует",
                        "schema": {
                            "$ref": "#/definitions/httpserver.quoteResponse"
                        }
                    },
                    "500": {
                        "description": "Проблемы на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/httpserver.quoteResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpserver.idData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "httpserver.idResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/httpserver.idData"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "httpserver.quoteData": {
            "type": "object",
            "properties": {
                "refresh_time": {
                    "type": "integer"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "httpserver.quoteResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/httpserver.quoteData"
                },
                "error": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "curr-quote",
	Description:      "Swagger документация к API сервиса котировок валютных курсов",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}