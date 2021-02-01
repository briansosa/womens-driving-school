package utils

import (
	"encoding/json"
	"fmt"
)

func MapperStructs(structFrom, structTo interface{}) error {
	structBytes, err := json.Marshal(structFrom)
	if err != nil {
		return fmt.Errorf("Error al realizar el marshal: %s", err)
	}
	err = json.Unmarshal(structBytes, structTo)
	if err != nil {
		return fmt.Errorf("Error al realizar el unmarshal: %s", err)
	}
	return nil
}
