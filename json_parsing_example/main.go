// Generic example of parsing JSON data using a Go map that supports multiple data types.

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hi! i'm a string.",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Could not marshal JSON data: %s\n", err)
		return
	}

	fmt.Printf("Voila! The JSON data is: %s\n", jsonData)
}
