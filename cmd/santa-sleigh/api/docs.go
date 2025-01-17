// SPDX-License-Identifier: ice License 1.0

// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "ice.io",
            "url": "https://ice.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1r/achievement-summaries/badges/users/{userId}": {
            "get": {
                "description": "Returns user's summary about badges.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Badges"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the id of the user you need summary for",
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
                                "$ref": "#/definitions/badges.BadgeSummary"
                            }
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1r/achievement-summaries/levels-and-roles/users/{userId}": {
            "get": {
                "description": "Returns user's summary about levels \u0026 roles.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Levels \u0026 Roles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the id of the user you need summary for",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/levelsandroles.Summary"
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1r/badges/{badgeType}/users/{userId}": {
            "get": {
                "description": "Returns all badges of the specific type for the user, with the progress for each of them.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Badges"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the id of the user you need progress for",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "level",
                            "coin",
                            "social"
                        ],
                        "type": "string",
                        "description": "the type of the badges",
                        "name": "badgeType",
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
                                "$ref": "#/definitions/badges.Badge"
                            }
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1r/tasks/x/users/{userId}": {
            "get": {
                "description": "Returns all the tasks and provided user's progress for each of them.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the id of the user you need progress for",
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
                                "$ref": "#/definitions/tasks.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1w/tasks/{taskType}/users/{userId}": {
            "put": {
                "description": "Completes the specific task (identified via task type) for the specified user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "enum": [
                            "claim_username",
                            "start_mining",
                            "upload_profile_picture",
                            "follow_us_on_twitter",
                            "join_telegram",
                            "invite_friends"
                        ],
                        "type": "string",
                        "description": "the type of the task",
                        "name": "taskType",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the id of the user that completed the task",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request params. Set it only if task completion requires additional data.",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/main.CompleteTaskRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "if user not found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "badges.AchievingRange": {
            "type": "object",
            "properties": {
                "fromInclusive": {
                    "type": "integer"
                },
                "toInclusive": {
                    "type": "integer"
                }
            }
        },
        "badges.Badge": {
            "type": "object",
            "properties": {
                "achieved": {
                    "type": "boolean"
                },
                "achievingRange": {
                    "$ref": "#/definitions/badges.AchievingRange"
                },
                "name": {
                    "type": "string"
                },
                "percentageOfUsersInProgress": {
                    "type": "number"
                },
                "type": {
                    "$ref": "#/definitions/badges.GroupType"
                }
            }
        },
        "badges.BadgeSummary": {
            "type": "object",
            "properties": {
                "index": {
                    "type": "integer"
                },
                "lastIndex": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/badges.GroupType"
                }
            }
        },
        "badges.GroupType": {
            "type": "string",
            "enum": [
                "level",
                "coin",
                "social"
            ],
            "x-enum-varnames": [
                "LevelGroupType",
                "CoinGroupType",
                "SocialGroupType"
            ]
        },
        "levelsandroles.Role": {
            "type": "object",
            "properties": {
                "enabled": {
                    "type": "boolean",
                    "example": true
                },
                "type": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/levelsandroles.RoleType"
                        }
                    ],
                    "example": "snowman"
                }
            }
        },
        "levelsandroles.RoleType": {
            "type": "string",
            "enum": [
                "ambassador"
            ],
            "x-enum-varnames": [
                "AmbassadorRoleType"
            ]
        },
        "levelsandroles.Summary": {
            "type": "object",
            "properties": {
                "level": {
                    "type": "integer",
                    "example": 11
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/levelsandroles.Role"
                    }
                }
            }
        },
        "main.CompleteTaskRequestBody": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/tasks.Data"
                }
            }
        },
        "server.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "SOMETHING_NOT_FOUND"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "error": {
                    "type": "string",
                    "example": "something is missing"
                }
            }
        },
        "tasks.Data": {
            "type": "object",
            "properties": {
                "requiredQuantity": {
                    "type": "integer",
                    "example": 3
                },
                "telegramUserHandle": {
                    "type": "string",
                    "example": "jdoe1"
                },
                "twitterUserHandle": {
                    "type": "string",
                    "example": "jdoe2"
                }
            }
        },
        "tasks.Task": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "data": {
                    "$ref": "#/definitions/tasks.Data"
                },
                "type": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/tasks.Type"
                        }
                    ],
                    "example": "claim_username"
                }
            }
        },
        "tasks.Type": {
            "type": "string",
            "enum": [
                "claim_username",
                "start_mining",
                "upload_profile_picture",
                "follow_us_on_twitter",
                "join_telegram",
                "invite_friends"
            ],
            "x-enum-varnames": [
                "ClaimUsernameType",
                "StartMiningType",
                "UploadProfilePictureType",
                "FollowUsOnTwitterType",
                "JoinTelegramType",
                "InviteFriendsType"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "latest",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https"},
	Title:            "Achievements API",
	Description:      "API that handles everything related to user's achievements and gamification progress.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
