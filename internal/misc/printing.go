package misc

import (
	"encoding/json"
	"fmt"
)

func PrettyPrintInterface(i interface{}) {
	data, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic("failed to marshal interface")
	}

	fmt.Println(string(data))
}
