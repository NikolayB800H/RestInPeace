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
            "name": "Рабраб",
            "url": "https://github.com/NikolayB800H",
            "email": "gorkunovnm@gmail.com"
        },
        "license": {
            "name": "AS IS (NO WARRANTY)"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/data_types": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Список видов данных включает только те, что со статусом \"доступен\"",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Виды данных"
                ],
                "summary": "Запросить все виды данных прогнозов и черновик заявки на прогноз",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.GetAllDataTypesResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Добавляет один вид данных с заданными полями",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Виды данных"
                ],
                "summary": "Запросить добавление вида данных прогнозов",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Изображение вида данных",
                        "name": "image_path",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Название вида данных",
                        "name": "data_type_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Погрешность предсказания вида данных",
                        "name": "precision",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Описание вида данных",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Единица измерения вида данных",
                        "name": "unit",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Статус вида данных",
                        "name": "data_type_status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/data_types/{data_type_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает более подробную информацию об одном виде данных",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Виды данных"
                ],
                "summary": "Запросить один вид данных прогнозов",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор вида данных",
                        "name": "data_type_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.DataTypes"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Изменяет один вид данных",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Виды данных"
                ],
                "summary": "Запросить изменение вида данных прогнозов",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор вида данных",
                        "name": "data_type_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Изображение вида данных",
                        "name": "image_path",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Название вида данных",
                        "name": "data_type_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Погрешность предсказания вида данных",
                        "name": "precision",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Описание вида данных",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Единица измерения вида данных",
                        "name": "unit",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Статус вида данных",
                        "name": "data_type_status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удаляет один вид данных по его data_type_id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Виды данных"
                ],
                "summary": "Запросить удаление вида данных прогнозов",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор вида данных",
                        "name": "data_type_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/data_types/{data_type_id}/add_to_forecast_application": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Добавляет данный вид данных в черновик заявки",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Виды данных"
                ],
                "summary": "Запросить добавление вида данных в заявку на прогноз",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор вида данных",
                        "name": "data_type_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllDataTypesResponse"
                        }
                    }
                }
            }
        },
        "/api/forecast_applications": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает все заявки с фильтрацией по статусу и дате формирования",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Запросить все заявки на прогнозы",
                "parameters": [
                    {
                        "type": "string",
                        "description": "статус заявки",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "начальная дата формирования",
                        "name": "formation_date_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "конечная дата формирвания",
                        "name": "formation_date_end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllForecastApplicationssResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удаляет черновую заявку пользователя",
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Удалить черновую заявку",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/forecast_applications/delete_data_type/{data_type_id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удаляет один вид данных по его data_type_id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Запросить удаление вида данных из черновика заявки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор вида данных",
                        "name": "data_type_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllDataTypesResponse"
                        }
                    }
                }
            }
        },
        "/api/forecast_applications/set_input/{data_type_id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Изменяет входные данные в связи ММ",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Запросить изменение входных данных вида данных черновика",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор вида данных",
                        "name": "data_type_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Входное значение за первый день",
                        "name": "input_first",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Входное значение за второй день",
                        "name": "input_second",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Входное значение за третий день",
                        "name": "input_third",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/forecast_applications/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Изменяет дату начала входных измерений черновика и возвращает его",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Запросить изменение черновика",
                "parameters": [
                    {
                        "type": "string",
                        "description": "дата начала входных измерений",
                        "name": "input_start_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.UpdateForecastApplicationsResponse"
                        }
                    }
                }
            }
        },
        "/api/forecast_applications/user_confirm": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Сформировать заявку пользователем",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Запросить формирование заявки",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.UpdateForecastApplicationsResponse"
                        }
                    }
                }
            }
        },
        "/api/forecast_applications/{application_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает более подробную информацию о заявке",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Запросить одну заявку на прогноз",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор заявки",
                        "name": "application_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.ForecastApplicationsResponse"
                        }
                    }
                }
            }
        },
        "/api/forecast_applications/{application_id}/moderator_confirm": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Подтвердить или отменить заявку модератором",
                "tags": [
                    "Заявки на прогнозы"
                ],
                "summary": "Подтвердить заявку",
                "parameters": [
                    {
                        "type": "string",
                        "description": "уникальный идентификатор заявки",
                        "name": "application_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "статус заявки",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.UpdateForecastApplicationsResponse"
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "description": "Авторизует пользователя по логиню, паролю и отдаёт jwt токен для дальнейших запросов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Авторизация",
                "parameters": [
                    {
                        "description": "login and password",
                        "name": "user_credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemes.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.SwaggerLoginResp"
                        }
                    }
                }
            }
        },
        "/api/user/logout": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Выход из аккаунта",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Выйти из аккаунта",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/user/sign_up": {
            "post": {
                "description": "Регистрация нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Регистрация",
                "parameters": [
                    {
                        "description": "login and password",
                        "name": "user_credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemes.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.SwaggerLoginResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ds.ConnectorAppsTypesDataTypes": {
            "type": "object",
            "required": [
                "data_type_name",
                "data_type_status",
                "description",
                "precision",
                "unit"
            ],
            "properties": {
                "data_type_id": {
                    "type": "string"
                },
                "data_type_name": {
                    "type": "string",
                    "maxLength": 128
                },
                "data_type_status": {
                    "description": "Replace with Enum",
                    "type": "string",
                    "maxLength": 50
                },
                "description": {
                    "type": "string",
                    "maxLength": 1024
                },
                "image_path": {
                    "type": "string"
                },
                "inputFirst": {
                    "type": "number"
                },
                "inputSecond": {
                    "type": "number"
                },
                "inputThird": {
                    "type": "number"
                },
                "output": {
                    "type": "number"
                },
                "precision": {
                    "type": "number"
                },
                "unit": {
                    "type": "string",
                    "maxLength": 32
                }
            }
        },
        "ds.DataTypes": {
            "type": "object",
            "required": [
                "data_type_name",
                "data_type_status",
                "description",
                "precision",
                "unit"
            ],
            "properties": {
                "data_type_id": {
                    "type": "string"
                },
                "data_type_name": {
                    "type": "string",
                    "maxLength": 128
                },
                "data_type_status": {
                    "description": "Replace with Enum",
                    "type": "string",
                    "maxLength": 50
                },
                "description": {
                    "type": "string",
                    "maxLength": 1024
                },
                "image_path": {
                    "type": "string"
                },
                "precision": {
                    "type": "number"
                },
                "unit": {
                    "type": "string",
                    "maxLength": 32
                }
            }
        },
        "schemes.AllDataTypesResponse": {
            "type": "object",
            "properties": {
                "data_types": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.DataTypes"
                    }
                }
            }
        },
        "schemes.AllForecastApplicationssResponse": {
            "type": "object",
            "properties": {
                "applications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemes.ForecastApplicationsOutput"
                    }
                }
            }
        },
        "schemes.ForecastApplicationsOutput": {
            "type": "object",
            "properties": {
                "application_completion_date": {
                    "type": "string"
                },
                "application_creation_date": {
                    "type": "string"
                },
                "application_formation_date": {
                    "type": "string"
                },
                "application_id": {
                    "type": "string"
                },
                "application_status": {
                    "description": "Replace with Enum",
                    "type": "string"
                },
                "calculate_status": {
                    "description": "Replace with Enum",
                    "type": "string"
                },
                "creator": {
                    "type": "string"
                },
                "input_start_date": {
                    "type": "string"
                },
                "moderator": {
                    "type": "string"
                }
            }
        },
        "schemes.ForecastApplicationsResponse": {
            "type": "object",
            "properties": {
                "application": {
                    "$ref": "#/definitions/schemes.ForecastApplicationsOutput"
                },
                "data_types": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.ConnectorAppsTypesDataTypes"
                    }
                }
            }
        },
        "schemes.ForecastApplicationsShort": {
            "type": "object",
            "properties": {
                "application_id": {
                    "type": "string"
                },
                "data_type_count": {
                    "type": "integer"
                }
            }
        },
        "schemes.GetAllDataTypesResponse": {
            "description": "Ответ с черновикомм заявки на прогноз и со всеми типами данных",
            "type": "object",
            "properties": {
                "data_types": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.DataTypes"
                    }
                },
                "draft_application": {
                    "$ref": "#/definitions/schemes.ForecastApplicationsShort"
                }
            }
        },
        "schemes.LoginReq": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 256
                },
                "password": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "schemes.RegisterReq": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 256
                },
                "password": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "schemes.SwaggerLoginResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "role": {
                    "type": "integer"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "schemes.UpdateForecastApplicationsResponse": {
            "type": "object",
            "properties": {
                "application": {
                    "$ref": "#/definitions/schemes.ForecastApplicationsOutput"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "0.0.0.0:8084",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "Прогнозы",
	Description:      "Сервис прогнозирования погодных параметров (условий)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
