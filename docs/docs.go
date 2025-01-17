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
        "/songs": {
            "get": {
                "description": "Получение списка из 20 песен упорядоченного по названию с пагинацией.\nЕсть возможность филтрации по дате релиза, названию группы, названию песни, а также возможность изменения количества песен на странице.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Manage songs"
                ],
                "summary": "Получение списка песен",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Release Date",
                        "name": "releaseDate",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 20,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of songs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request\" SchemaExample({\"code\":400, \"message\":\"Bad request\"})",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error\" SchemaExample({\"code\":500, \"message\":\"Internal server error\"})",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/newsong": {
            "post": {
                "description": "Добавление песни в библиотеку. При добавлении записываются только название песни и группы. Далее происходит запрос в сторонний сервис и после песня добавляется в библиотеку.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Manage songs"
                ],
                "summary": "Добавление песни",
                "parameters": [
                    {
                        "description": "Song details to create",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success message\"                  \tSchemaExample({\"code\":200, \"message\":\"The song has been added to the library.\"})",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request or missing required fields\" SchemaExample({\"code\":400, \"message\":\"Bad request\"})",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error\"            SchemaExample({\"code\":500, \"message\":\"Internal server error.\"})",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/{songId}/editsong": {
            "put": {
                "description": "Редактирование параметров песни по Id. Все поля песни являются редактируемыми.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Manage songs"
                ],
                "summary": "Редактирование песни",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song ID",
                        "name": "songId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Song details to edit",
                        "name": "song",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success message\"               SchemaExample({\"code\":200, \"message\":\"The song has been successfully edited.\"})",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request\"                   SchemaExample({\"code\":400, \"message\":\"Bad request\"})",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Song not found\"                SchemaExample({\"code\":404, \"message\":\"Song not found\"})",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error\"         SchemaExample({\"code\":500, \"message\":\"Internal server error.\"})",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/{songId}/textsong": {
            "get": {
                "description": "Получение текста песни по ID с пагинацией по куплетам.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Manage songs"
                ],
                "summary": "Получение текста песни",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song ID",
                        "name": "songId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song lyrics text\"         SchemaExample({\"code\":200, \"message\":\"Textttt of song\"})",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Song not found\"           SchemaExample({\"code\":404, \"message\":\"Song not found\"})",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error\"    SchemaExample({\"code\":500, \"message\":\"Internal server error.\"})",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "Muse"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=N-_mHedypEU"
                },
                "releaseDate": {
                    "type": "string",
                    "example": "Oh baby don't you know I suffer?"
                },
                "song": {
                    "type": "string",
                    "example": "Supermassive Black Hole"
                },
                "text": {
                    "type": "string",
                    "example": "19.06.2006"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/songs",
	Schemes:          []string{},
	Title:            "Music info API",
	Description:      "This is an example of a server for managing songs as a test task.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
