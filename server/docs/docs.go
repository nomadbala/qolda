package docs

import "github.com/swaggo/swag"

var doc = `

`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Qolda Docs",
	Description:      "Qolda API Documentation",
	InfoInstanceName: "Docs",
	SwaggerTemplate:  doc,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
