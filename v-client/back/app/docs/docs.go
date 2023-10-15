// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://example.com/terms",
        "contact": {
            "name": "Your Name",
            "email": "youremail@example.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/atm": {
            "get": {
                "description": "Retrieve a list of ATMs",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of ATMs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.ATM"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new ATM entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new ATM",
                "parameters": [
                    {
                        "description": "New ATM data",
                        "name": "newATM",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ATM"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.ATM"
                        }
                    }
                }
            }
        },
        "/atm_filters": {
            "get": {
                "description": "Retrieve a list of ATM filters",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of ATM filters",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.ATMFilter"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new ATM filter entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new ATM filter",
                "parameters": [
                    {
                        "description": "New ATM filter data",
                        "name": "newATMFilter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ATMFilter"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.ATMFilter"
                        }
                    }
                }
            }
        },
        "/salepoint": {
            "get": {
                "description": "Retrieve a list of SalePoints",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of SalePoints",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.SalePoint"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new SalePoint entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new SalePoint",
                "parameters": [
                    {
                        "description": "New SalePoint data",
                        "name": "newSalePoint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SalePoint"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.SalePoint"
                        }
                    }
                }
            }
        },
        "/salepoint_filters": {
            "get": {
                "description": "Retrieve a list of SalePoint filters",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of SalePoint filters",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.SalePointFilter"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new SalePoint filter entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new SalePoint filter",
                "parameters": [
                    {
                        "description": "New SalePoint filter data",
                        "name": "newSalePointFilter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SalePointFilter"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.SalePointFilter"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ATM": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "ATM address\nrequired: true",
                    "type": "string"
                },
                "allDay": {
                    "description": "Indicates if the ATM operates 24/7\nrequired: true",
                    "type": "string"
                },
                "id_atms": {
                    "description": "ID of the ATM\nrequired: true",
                    "type": "integer"
                },
                "latitude": {
                    "description": "Latitude of the ATM\nrequired: true",
                    "type": "number"
                },
                "longitude": {
                    "description": "Longitude of the ATM\nrequired: true",
                    "type": "number"
                },
                "services": {
                    "description": "Services provided by the ATM\nrequired: true",
                    "type": "string"
                }
            }
        },
        "main.ATMFilter": {
            "type": "object",
            "properties": {
                "cash": {
                    "description": "Cash filter for the ATM\nrequired: true",
                    "type": "integer"
                },
                "id_atms": {
                    "description": "ID of the ATM Filter\nrequired: true",
                    "type": "integer"
                }
            }
        },
        "main.SalePoint": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "SalePoint Address\nrequired: true",
                    "type": "string"
                },
                "distance": {
                    "description": "Distance\nrequired: true",
                    "type": "integer"
                },
                "hasRamp": {
                    "description": "Has Ramp\nrequired: true",
                    "type": "string"
                },
                "kep": {
                    "description": "Kep\nrequired: true",
                    "type": "string"
                },
                "latitude": {
                    "description": "Latitude of the SalePoint\nrequired: true",
                    "type": "number"
                },
                "longitude": {
                    "description": "Longitude of the SalePoint\nrequired: true",
                    "type": "number"
                },
                "metroStation": {
                    "description": "Metro Station\nrequired: true",
                    "type": "string"
                },
                "myBranch": {
                    "description": "My Branch\nrequired: true",
                    "type": "string"
                },
                "network": {
                    "description": "Network\nrequired: true",
                    "type": "string"
                },
                "officeType": {
                    "description": "Office Type\nrequired: true",
                    "type": "string"
                },
                "offices_id": {
                    "description": "ID of the SalePoint\nrequired: true",
                    "type": "integer"
                },
                "openHours": {
                    "description": "SalePoint Open Hours\nrequired: true",
                    "type": "string"
                },
                "openHoursIndividual": {
                    "description": "Open Hours Individual\nrequired: true",
                    "type": "string"
                },
                "rko": {
                    "description": "RKO\nrequired: true",
                    "type": "string"
                },
                "salePointCode": {
                    "description": "SalePoint Code\nrequired: true",
                    "type": "string"
                },
                "salePointFormat": {
                    "description": "SalePoint Format\nrequired: true",
                    "type": "string"
                },
                "salePointName": {
                    "description": "SalePoint Name\nrequired: true",
                    "type": "string"
                },
                "status": {
                    "description": "SalePoint Status\nrequired: true",
                    "type": "string"
                },
                "suoAvailability": {
                    "description": "SUO Availability\nrequired: true",
                    "type": "string"
                }
            }
        },
        "main.SalePointFilter": {
            "type": "object",
            "properties": {
                "current_workload": {
                    "description": "Current Workload filter for SalePoint\nrequired: true",
                    "type": "integer"
                },
                "offices_id": {
                    "description": "ID of the SalePoint Filter\nrequired: true",
                    "type": "integer"
                },
                "rating": {
                    "description": "Rating filter for SalePoint\nrequired: true",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Your API Title",
	Description:      "Your API description",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
