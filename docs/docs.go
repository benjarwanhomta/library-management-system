package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add_book": {
            "post": {
                "description": "เพิ่มข้อมูลหนังสือ",
                "summary": "เพิ่มข้อมูลหนังสือ",
                "parameters": [
                    {
                        "in": "body",
                        "name": "book",
                        "description": "ข้อมูลหนังสือที่ต้องการเพิ่ม",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful add a new book",
                        "schema": {
                            "$ref": "#/definitions/globalResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Book": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                },
                "author_id": {
                    "type": "integer"
                },
                "category_id": {
                    "type": "integer"
                },
                "publish_year": {
                    "type": "integer"
                },
                "isbn": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "available_qty": {
                    "type": "integer"
                }
            },
            "required": [
                "title",
                "author_id",
                "category_id",
                "publish_year",
                "isbn",
                "description",
                "available_qty"
            ]
        },
        "globalResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "value": {
                    "type": "object"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{"https", "http"},
	Title:       "Library Management System API",
	Description: "การเริ่มต้นใช้งาน API ระบบจัดการห้องสมุด",
}

type s struct{}

// ReadDoc generates the Swagger documentation.
func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
