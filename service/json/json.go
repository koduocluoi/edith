package json

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

func Run(args []string) {
	var result map[string]interface{}
	json.Unmarshal([]byte(args[0]), &result)

	data, _ := json.Marshal(result)
	fmt.Println(string(pretty.Pretty(data)))
}
