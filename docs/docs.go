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
        "/admins/banners": {
            "get": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "admin get filtered banners",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admins-orders"
                ],
                "summary": "Admin Get Filtered Banners",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "feature_id",
                        "name": "feature_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "tag_id",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "admin create new banner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admins-banners"
                ],
                "summary": "Admin Create New Banner",
                "parameters": [
                    {
                        "description": "banner info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.createBannerInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    }
                }
            }
        },
        "/admins/banners/{id}": {
            "put": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "admin update banner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admins-banners"
                ],
                "summary": "Admin Update Banner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "banner id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "banner update info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.updateBannerInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "admin delete banner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admins-banners"
                ],
                "summary": "Admin Delete Banner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "banner id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    }
                }
            }
        },
        "/admins/sign-in": {
            "post": {
                "description": "admin sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admins-auth"
                ],
                "summary": "Admin SignIn",
                "parameters": [
                    {
                        "description": "sign up info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    }
                }
            }
        },
        "/users/banner": {
            "get": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "user get banner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users-banners"
                ],
                "summary": "User Get Banner",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "feature_id",
                        "name": "feature_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "tag_id",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "use_last_revision",
                        "name": "use_last_revision",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    }
                }
            }
        },
        "/users/sign-in": {
            "post": {
                "description": "user sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users-auth"
                ],
                "summary": "User SignIn",
                "parameters": [
                    {
                        "description": "sign up info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal_delivery_http_v1.createBannerInput": {
            "type": "object",
            "required": [
                "content",
                "feature_id",
                "tag_ids"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "feature_id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "tag_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "internal_delivery_http_v1.response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "internal_delivery_http_v1.signInInput": {
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
        "internal_delivery_http_v1.tokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        },
        "internal_delivery_http_v1.updateBannerInput": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "feature_id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "tag_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "description": "JWT token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Banner Management Service",
	Description:      "This service provides an interface for managing and retrieving banners.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
