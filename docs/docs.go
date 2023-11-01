// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "耗子科技",
            "email": "i@haozi.net"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/panel/cert/algorithms": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "获取面板证书管理支持的算法列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "获取算法列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/caProviders": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "获取面板证书管理支持的 CA 提供商",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "获取 CA 提供商",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/certs": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "获取面板证书管理的证书列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "获取证书列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllers.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.CertList"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "添加证书到面板证书管理",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "添加证书",
                "parameters": [
                    {
                        "description": "证书信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CertAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/certs/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "删除面板证书管理的证书",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "删除证书",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "证书 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/dns": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "获取面板证书管理的 DNS 接口列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "获取 DNS 接口列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllers.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.DNSList"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "添加 DNS 接口到面板证书管理",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "添加 DNS 接口",
                "parameters": [
                    {
                        "description": "DNS 接口信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.DNSAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/dns/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "删除面板证书管理的 DNS 接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "删除 DNS 接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "DNS 接口 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/dnsProviders": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "获取面板证书管理支持的 DNS 提供商",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "获取 DNS 提供商",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/manualDNS": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "获取签发证书所需的 DNS 记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "获取手动 DNS 记录",
                "parameters": [
                    {
                        "description": "证书信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.Obtain"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllers.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "$ref": "#/definitions/acme.Resolve"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/obtain": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "签发面板证书管理的证书",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "签发证书",
                "parameters": [
                    {
                        "description": "证书信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.Obtain"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/renew": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "续签面板证书管理的证书",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "续签证书",
                "parameters": [
                    {
                        "description": "证书信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.Renew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/users": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "获取面板证书管理的 ACME 用户列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "获取用户列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllers.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.CertList"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "添加 ACME 用户到面板证书管理",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "添加 ACME 用户",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/cert/users/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "删除面板证书管理的 ACME 用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "证书"
                ],
                "summary": "删除 ACME 用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "登录已过期",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/panel/user/login": {
            "post": {
                "description": "通过用户名和密码获取访问令牌",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "403": {
                        "description": "用户名或密码错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "系统内部错误",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "acme.DNSParam": {
            "type": "object",
            "properties": {
                "access_key": {
                    "type": "string"
                },
                "api_key": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "secret_key": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "acme.Resolve": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controllers.SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Cert": {
            "type": "object",
            "properties": {
                "auto_renew": {
                    "description": "自动续签",
                    "type": "boolean"
                },
                "cert": {
                    "description": "证书内容",
                    "type": "string"
                },
                "cert_url": {
                    "description": "证书 URL (续签时使用)",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "dns": {
                    "$ref": "#/definitions/models.CertDNS"
                },
                "dns_id": {
                    "description": "关联的 DNS ID",
                    "type": "integer"
                },
                "domains": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "description": "私钥内容",
                    "type": "string"
                },
                "type": {
                    "description": "证书类型 (P256, P384, 2048, 4096)",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.CertUser"
                },
                "user_id": {
                    "description": "关联的 ACME 用户 ID",
                    "type": "integer"
                },
                "website": {
                    "$ref": "#/definitions/models.Website"
                },
                "website_id": {
                    "description": "关联的网站 ID",
                    "type": "integer"
                }
            }
        },
        "models.CertDNS": {
            "type": "object",
            "properties": {
                "certs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Cert"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "dns_param": {
                    "$ref": "#/definitions/acme.DNSParam"
                },
                "id": {
                    "type": "integer"
                },
                "type": {
                    "description": "DNS 提供商 (dnspod, aliyun, cloudflare)",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CertUser": {
            "type": "object",
            "properties": {
                "ca": {
                    "description": "CA 提供商 (letsencrypt, zerossl, sslcom, google, buypass)",
                    "type": "string"
                },
                "certs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Cert"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "hmac_encoded": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key_type": {
                    "type": "string"
                },
                "kid": {
                    "type": "string"
                },
                "private_key": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Website": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "php": {
                    "type": "integer"
                },
                "remark": {
                    "type": "string"
                },
                "ssl": {
                    "type": "boolean"
                },
                "status": {
                    "type": "boolean"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "requests.CertAdd": {
            "type": "object",
            "properties": {
                "auto_renew": {
                    "type": "boolean"
                },
                "dns_id": {
                    "type": "integer"
                },
                "domains": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "requests.DNSAdd": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/acme.DNSParam"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "requests.Login": {
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
        "requests.Obtain": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "requests.Renew": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "requests.UserAdd": {
            "type": "object",
            "properties": {
                "ca": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "hmac_encoded": {
                    "type": "string"
                },
                "key_type": {
                    "type": "string"
                },
                "kid": {
                    "type": "string"
                }
            }
        },
        "responses.CertList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Cert"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "responses.DNSList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CertDNS"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "耗子 Linux 面板 API",
	Description:      "耗子 Linux 面板的 API 信息",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
