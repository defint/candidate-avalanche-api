// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-04-24 14:31:12.130797091 +0300 EEST m=+0.058141446

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Avalanche Candidates API",
        "title": "Avalanche Candidates API",
        "contact": {
            "name": "Anton Matrenin",
            "email": "matrenin@ukr.net"
        },
        "license": {},
        "version": "1.0"
    },
    "paths": {
        "/candidate": {
            "get": {
                "description": "Show a list of candidates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a list of candidates",
                "operationId": "candidates-list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Candidate"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Create a candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a candidate",
                "operationId": "candidate-create",
                "parameters": [
                    {
                        "description": "Candidate",
                        "name": "candidate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Candidate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Candidate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete all candidates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete all candidates",
                "operationId": "candidates-delete",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/candidate/{id}": {
            "get": {
                "description": "Show a candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a candidate",
                "operationId": "candidate-item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Candidate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Update a candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a candidate",
                "operationId": "candidate-update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Candidate",
                        "name": "candidate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Candidate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Candidate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a candidate",
                "operationId": "candidate-delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/candidate/{id}/status": {
            "post": {
                "description": "Update a status of candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a status of candidate",
                "operationId": "candidate-status-update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Status",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Status"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Candidate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Candidate": {
            "type": "object",
            "properties": {
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.History"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.History": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "status_from": {
                    "type": "string"
                },
                "status_to": {
                    "type": "string"
                }
            }
        },
        "models.Status": {
            "type": "object",
            "properties": {
                "reason": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
