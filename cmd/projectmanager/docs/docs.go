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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/account": {
            "post": {
                "description": "add new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add new account",
                "parameters": [
                    {
                        "description": "새로운 Account 등록",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/gin.AddAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.AddAccountResponse"
                        }
                    }
                }
            }
        },
        "/api/account/delete/{accountId}": {
            "delete": {
                "description": "delete account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Account 삭제",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/gin.DeleteAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.DeleteAccountResponse"
                        }
                    }
                }
            }
        },
        "/api/account/detail/{accountId}": {
            "get": {
                "description": "get account info by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get account info by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.GetAccountResponse"
                        }
                    }
                }
            }
        },
        "/api/account/list": {
            "get": {
                "description": "list current accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list current accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.ListAccountResponse"
                        }
                    }
                }
            }
        },
        "/api/account/nickname/{accountId}": {
            "put": {
                "description": "update nickname of existing account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update nickname of existing account",
                "parameters": [
                    {
                        "description": "기존 Account Nickname 수정",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/gin.UpdateAccountNicknameRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.UpdateAccountNicknameResponse"
                        }
                    }
                }
            }
        },
        "/api/account/password/{accountId}": {
            "put": {
                "description": "update password of existing account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update password of existing account",
                "parameters": [
                    {
                        "description": "기존 Account Password 수정",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/gin.UpdateAccountPasswordRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.UpdateAccountPasswordResponse"
                        }
                    }
                }
            }
        },
        "/api/registry/list": {
            "get": {
                "description": "list registry type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list registry type",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.ListRegistryTypeResponse"
                        }
                    }
                }
            }
        },
        "/api/repository": {
            "get": {
                "description": "list current repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list current repository",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.ListRepositoryResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "add new repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add new repository",
                "parameters": [
                    {
                        "description": "Repository 등록",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/gin.AddRepositoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.AddRepositoryResponse"
                        }
                    }
                }
            }
        },
        "/api/repository/delete/{repositoryId}": {
            "delete": {
                "description": "delete repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete repository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository ID",
                        "name": "repositoryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Repository 삭제",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/gin.DeleteRepositoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.DeleteRepositoryResponse"
                        }
                    }
                }
            }
        },
        "/api/repository/detail/{repositoryId}": {
            "get": {
                "description": "get repository info by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get repository info by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository ID",
                        "name": "repositoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.GetRepositoryResponse"
                        }
                    }
                }
            }
        },
        "/api/repository/image/{repositoryId}": {
            "put": {
                "description": "Create or Update Image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create or Update Image",
                "parameters": [
                    {
                        "description": "Image 정보",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/gin.CreateOrUpdateImageRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Repository ID",
                        "name": "repositoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.CreateOrUpdateImageResponse"
                        }
                    }
                }
            }
        },
        "/api/repository/images/{repositoryId}": {
            "get": {
                "description": "List Image By Repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List Image By Repository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository ID",
                        "name": "repositoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.ListImageByRepository"
                        }
                    }
                }
            }
        },
        "/api/repository/latest/{repositoryId}": {
            "get": {
                "description": "Get Latest Image by Repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Latest Image by Repository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository ID",
                        "name": "repositoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.GetLatestImageByRepository"
                        }
                    }
                }
            }
        },
        "/internal/account/{accountId}/private": {
            "get": {
                "description": "get account info by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get account info by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.GetAccountPrivateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.Account": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "registry_type": {
                    "type": "string"
                },
                "registry_url": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "gin.AccountPrivate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "gin.AddAccountRequest": {
            "type": "object",
            "required": [
                "account_password",
                "account_username",
                "registry_type"
            ],
            "properties": {
                "account_nickname": {
                    "type": "string"
                },
                "account_password": {
                    "type": "string"
                },
                "account_username": {
                    "type": "string"
                },
                "registry_type": {
                    "type": "string"
                },
                "registry_url": {
                    "type": "string"
                }
            }
        },
        "gin.AddAccountResponse": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                }
            }
        },
        "gin.AddRepositoryRequest": {
            "type": "object",
            "required": [
                "account_id",
                "repository_name"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "repository_name": {
                    "type": "string"
                }
            }
        },
        "gin.AddRepositoryResponse": {
            "type": "object",
            "properties": {
                "repository_id": {
                    "type": "string"
                }
            }
        },
        "gin.CreateOrUpdateImageRequest": {
            "type": "object",
            "required": [
                "tag"
            ],
            "properties": {
                "tag": {
                    "type": "string"
                }
            }
        },
        "gin.CreateOrUpdateImageResponse": {
            "type": "object",
            "properties": {
                "image_id": {
                    "type": "string"
                }
            }
        },
        "gin.DeleteAccountRequest": {
            "type": "object"
        },
        "gin.DeleteAccountResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "gin.DeleteRepositoryRequest": {
            "type": "object"
        },
        "gin.DeleteRepositoryResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "gin.GetAccountPrivateResponse": {
            "type": "object",
            "properties": {
                "account_private": {
                    "$ref": "#/definitions/gin.AccountPrivate"
                }
            }
        },
        "gin.GetAccountResponse": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/gin.Account"
                }
            }
        },
        "gin.GetLatestImageByRepository": {
            "type": "object",
            "properties": {
                "image": {
                    "$ref": "#/definitions/gin.Image"
                }
            }
        },
        "gin.GetRepositoryResponse": {
            "type": "object",
            "properties": {
                "repository": {
                    "$ref": "#/definitions/gin.Repository"
                }
            }
        },
        "gin.Image": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "repository_id": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "gin.ListAccountResponse": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/gin.Account"
                    }
                }
            }
        },
        "gin.ListImageByRepository": {
            "type": "object",
            "properties": {
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/gin.Image"
                    }
                }
            }
        },
        "gin.ListRegistryTypeResponse": {
            "type": "object",
            "properties": {
                "repository_types": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "gin.ListRepositoryResponse": {
            "type": "object",
            "properties": {
                "repositories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/gin.Repository"
                    }
                }
            }
        },
        "gin.Repository": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "gin.UpdateAccountNicknameRequest": {
            "type": "object",
            "required": [
                "nickname"
            ],
            "properties": {
                "nickname": {
                    "type": "string"
                }
            }
        },
        "gin.UpdateAccountNicknameResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "gin.UpdateAccountPasswordRequest": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "gin.UpdateAccountPasswordResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
