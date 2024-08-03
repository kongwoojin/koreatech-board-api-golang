// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Developer",
            "email": "kongwoojin03@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/article": {
            "get": {
                "description": "Get article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notice"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ApiArticle"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/{department}/{board}": {
            "get": {
                "description": "Get notice list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notice"
                ],
                "summary": "Get notice list",
                "parameters": [
                    {
                        "enum": [
                            "arch",
                            "cse",
                            "dorm",
                            "mse",
                            "ace",
                            "ide",
                            "ite",
                            "mechanical",
                            "mechatronics",
                            "school",
                            "sim"
                        ],
                        "type": "string",
                        "description": "name of the department",
                        "name": "department",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "notice",
                            "free",
                            "job",
                            "pds",
                            "lecture",
                            "bachelor",
                            "scholar"
                        ],
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page of board",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "items per page",
                        "name": "num_of_items",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/{department}/{board}/search/title": {
            "get": {
                "description": "Search article from specific board by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notice"
                ],
                "summary": "Search article by title",
                "parameters": [
                    {
                        "enum": [
                            "arch",
                            "cse",
                            "dorm",
                            "mse",
                            "ace",
                            "ide",
                            "ite",
                            "mechanical",
                            "mechatronics",
                            "school",
                            "sim"
                        ],
                        "type": "string",
                        "description": "name of the department",
                        "name": "department",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "notice",
                            "free",
                            "job",
                            "pds",
                            "lecture",
                            "bachelor",
                            "scholar"
                        ],
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page of board",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "items per page",
                        "name": "num_of_items",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ApiArticle"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/{department}/{board}/widget": {
            "get": {
                "description": "Get minumum notice list for board widget, only 5 new notices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notice"
                ],
                "summary": "Get minumum notice list",
                "parameters": [
                    {
                        "enum": [
                            "arch",
                            "cse",
                            "dorm",
                            "mse",
                            "ace",
                            "ide",
                            "ite",
                            "mechanical",
                            "mechatronics",
                            "school",
                            "sim"
                        ],
                        "type": "string",
                        "description": "name of the department",
                        "name": "department",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "notice",
                            "free",
                            "job",
                            "pds",
                            "lecture",
                            "bachelor",
                            "scholar"
                        ],
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.APIData": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error message",
                    "type": "string"
                },
                "last_page": {
                    "description": "Last page of requested board",
                    "type": "integer"
                },
                "posts": {
                    "description": "Data of board",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Board"
                    }
                },
                "status_code": {
                    "description": "Status code of request",
                    "type": "integer"
                }
            }
        },
        "model.ApiArticle": {
            "type": "object",
            "properties": {
                "article_url": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "error": {
                    "description": "Error message",
                    "type": "string"
                },
                "files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Files"
                    }
                },
                "id": {
                    "type": "string"
                },
                "is_notice": {
                    "type": "boolean"
                },
                "num": {
                    "type": "integer"
                },
                "status_code": {
                    "description": "Status code of request",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "write_date": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "model.Board": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "is_new": {
                    "type": "boolean"
                },
                "is_notice": {
                    "type": "boolean"
                },
                "num": {
                    "type": "integer"
                },
                "read_count": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "write_date": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "model.Files": {
            "type": "object",
            "properties": {
                "file_name": {
                    "type": "string"
                },
                "file_url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v3",
	Schemes:          []string{},
	Title:            "KOREATECH board REST API",
	Description:      "This is unofficial version of KOREATECH board REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
