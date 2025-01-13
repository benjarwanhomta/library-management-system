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
                            "$ref": "#/definitions/addBookReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/definitions/globalResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/detail_book/{book_id}": {
            "get": {
                "description": "ดึงรายละเอียดข้อมูลหนังสือ",
                "summary": "ดึงรายละเอียดข้อมูลหนังสือตาม book_id",
                "parameters": [
                    {
                        "name": "book_id",
                        "in": "path",
                        "description": "ID ของหนังสือ",
                        "required": true,
                        "type": "integer"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ข้อมูลหนังสือ",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/definitions/bookDetailResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/delete_book/{book_id}": {
            "delete": {
                "description": "ลบข้อมูลหนังสือ",
                "summary": "ลบข้อมูลหนังสือตาม book_id",
                "parameters": [
                    {
                        "name": "book_id",
                        "in": "path",
                        "description": "ID ของหนังสือ",
                        "required": true,
                        "type": "integer"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ลบข้อมูลหนังสือสำเร็จ",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/definitions/globalResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/edit_book": {
            "put": {
                "description": "แก้ไขข้อมูลหนังสือ",
                "summary": "แก้ไขข้อมูลหนังสือ",
                "parameters": [
                    {
                        "in": "body",
                        "name": "book",
                        "description": "แก้ไขข้อมูลหนังสือที่ต้องการเพิ่ม",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/editBookReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/definitions/globalResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/search_book_all": {
            "get": {
                "description": "แสดงข้อมูลหนังสือที่ต้องการ",
                "summary": "แสดงข้อมูลหนังสือที่ต้องการ",
                "parameters": [
                    {
                        "in": "query",
                        "name": "title",
                        "description": "ชื่อของหนังสือที่ต้องการค้นหา",
                        "required": false,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "in": "query",
                        "name": "author",
                        "description": "ชื่อผู้แต่งของหนังสือ",
                        "required": false,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "in": "query",
                        "name": "category",
                        "description": "หมวดหมู่ของหนังสือ",
                        "required": false,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/definitions/globalResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/top_borrowed_book": {
            "get": {
                "description": "ดึงข้อมูลหนังสือที่ถูกยืมมากที่สุด",
                "summary": "ดึงข้อมูลหนังสือที่ถูกยืมมากที่สุด",
                "responses": {
                    "200": {
                        "description": "รายการหนังสือที่ถูกยืมมากที่สุด",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/definitions/topBorrowedBookResponse"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
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
        },
        "bookDetailResponse": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer",
                    "example": 2
                },
                "title": {
                    "type": "string",
                    "example": "หนังสือ A"
                },
                "author_id": {
                    "type": "integer",
                    "example": 1
                },
                "category_id": {
                    "type": "integer",
                    "example": 2
                },
                "publish_year": {
                    "type": "integer",
                    "example": 2020
                },
                "isbn": {
                    "type": "string",
                    "example": "9783-16-148410-0"
                },
                "description": {
                    "type": "string",
                    "example": "คำอธิบายของหนังสือ A"
                },
                "available_qty": {
                    "type": "integer",
                    "example": 5
                },
                "Author": {
                    "type": "object",
                    "properties": {
                        "author_id": {
                            "type": "integer",
                            "example": 1
                        },
                        "name": {
                            "type": "string",
                            "example": "จุฑามาศ พันธ์เจริญ"
                        },
                        "biography": {
                            "type": "string",
                            "example": "ผู้แต่งนิยายและบทความด้านสังคมศาสตร์ มีผลงานมากมายในด้านการศึกษาประวัติศาสตร์และวัฒนธรรมไทย"
                        }
                    }
                },
                "Category": {
                    "type": "object",
                    "properties": {
                        "category_id": {
                            "type": "integer",
                            "example": 2
                        },
                        "category_name": {
                            "type": "string",
                            "example": "วิทยาศาสตร์"
                        }
                    }
                }
            }
        },
        "addBookReq": {
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
        "editBookReq": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
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
                "book_id",
                "title",
                "author_id",
                "category_id",
                "publish_year",
                "isbn",
                "description",
                "available_qty"
            ]
        },
        "topBorrowedBookResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully get top borrowed book"
                },
                "status": {
                    "type": "string",
                    "example": "Success"
                },
                "value": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "title": {
                                "type": "string",
                                "example": "Book A"
                            },
                            "borrow_count": {
                                "type": "integer",
                                "example": 1
                            }
                        }
                    }
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
