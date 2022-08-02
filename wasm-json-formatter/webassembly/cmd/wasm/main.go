package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Welcome to webAssembly")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments!"
		}

		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return "Unable to get JS document object!"
		}
		outputTextbox := jsDoc.Call("getElementById", "jsonOutput")
		if !outputTextbox.Truthy() {
			return "Unable to get output textbox"
		}

		input := args[0].String()
		fmt.Printf("input: %s\n", input)
		pretty, err := prettyJson(input)
		if err != nil {
			errValue := fmt.Sprintf("Unable to convert to JSON: %s", err)
			errResult := map[string]interface{}{
				"wasmCustomError": errValue,
			}
			return errResult
		}
		outputTextbox.Set("value", pretty)
		return nil
	})
	return jsonFunc
}
