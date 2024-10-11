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
        "/healthcheck": {
            "get": {
                "description": "Check the health status of the service",
                "produces": [
                    "text/plain"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "STATUS: OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate a user and return a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Log out the user by clearing the authentication token",
                "produces": [
                    "application/json"
                ],
                "summary": "Logout a user",
                "responses": {
                    "200": {
                        "description": "Logged out successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/places": {
            "get": {
                "description": "Retrieve a list of places from the database",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of places",
                "responses": {
                    "200": {
                        "description": "List of places",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Place"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/places/{placeID}/reviews": {
            "get": {
                "description": "Get all reviews for a specific place",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieve reviews by place ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Place ID",
                        "name": "placeID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of reviews",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Review"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid place ID",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "No reviews found for the place",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve reviews",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reviews": {
            "post": {
                "description": "Create a new review for a place",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new review",
                "parameters": [
                    {
                        "description": "Review details",
                        "name": "review",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Review"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Review created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Review"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create review",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reviews/{id}": {
            "get": {
                "description": "Get review details by review ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieve a review by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Review ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Review details",
                        "schema": {
                            "$ref": "#/definitions/models.Review"
                        }
                    },
                    "400": {
                        "description": "Invalid review ID",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Review not found",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve review",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update review details by review ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an existing review",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Review ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated review details",
                        "name": "review",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Review"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Review updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Review"
                        }
                    },
                    "400": {
                        "description": "Invalid review ID",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Review not found",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update review",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a review by review ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a review",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Review ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Review deleted successfully"
                    },
                    "400": {
                        "description": "Invalid review ID",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Review not found",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to delete review",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Create a new user with login and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Sign up a new user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.Credentials"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "description": "Retrieve the current authenticated user information",
                "produces": [
                    "application/json"
                ],
                "summary": "Get the current user",
                "responses": {
                    "200": {
                        "description": "Current user",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpresponses.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.Credentials": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "httpresponses.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Place": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Review": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "place_id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "review_text": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "login": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
