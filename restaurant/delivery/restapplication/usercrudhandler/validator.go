package usercrudhandler

import (
	"bytes"
	"fmt"
	logger "log"

	customerrors "github.com/gohttpexamples/restaurant/delivery/restapplication/packages/errors"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateUserCreateUpdateRequest(rStr string) (bool, error) {
	//logger := loggerutils.GetLogger()
	schemaStr := `
	{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "object",
		"properties": {
		  "Name ": {
            "type": "string",
			"minLength": 1,
			"maxLength": 20
		  },
		  "Address": {
            "type": "string",
			"minLength": 1,
			"maxLength": 20
		  },
		  "AddressLine2": {
            "type": "string",
			"minLength": 1,
			"maxLength": 20
		  },
		  "URL": {
            "type": "string",
			"minLength": 1,
			"maxLength": 20
		  },
		  "Outcode": {
            "type": "string",
			"minLength": 1,
			"maxLength": 20
		  },
		  "Postcode": {
            "type": "string",
			"minLength": 1,
			"maxLength": 20
		  },
		  "Rating": {
            "type": "number",
			"minLength": 1,
			"maxLength": 20
		  },
		  "Type_of_food": {
            "type": "string",
			"minimum": 1,
			"maximum": 120
		  }
		},
		"required": [
		  "Name",
		  "Address",
		  "AddressLine2"
		]
	  }`

	schema := gojsonschema.NewStringLoader(schemaStr)
	content := gojsonschema.NewStringLoader(rStr)
	result, err := gojsonschema.Validate(schema, content)

	if err != nil {
		logger.Fatalf("Invalid Json Schema Error: %v", err)
		return false, customerrors.InternalError(fmt.Sprintf("Invalid Json Schema Error: %v", err))
		//panic(err)
	}
	if result.Valid() {
		return true, nil
	} else {
		var buffer bytes.Buffer
		for _, resulterr := range result.Errors() {
			logger.Printf("- %s\n", resulterr)
			errString := fmt.Sprintf("Field: %s - %s, ", resulterr.Field(), resulterr.Description())
			buffer.Write([]byte(errString))

		}
		errorDesc := buffer.String()
		return false, customerrors.BadRequest(errorDesc)
	}

}
