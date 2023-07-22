package output

import (
	"encoding/json"
	"os"
)

func LangsMapToJSONStr(langs map[string]int) []byte {
	jsonStr, err := json.Marshal(langs)
	if err != nil {
		panic(err)
	}
	return jsonStr
}

func OutputJSON(jsonStr []byte) {
	langsJSONFile, err := os.Create("langs.json")
	if err != nil {
		panic(err)
	}
	_,err = langsJSONFile.Write(jsonStr)
	if err != nil {
		panic(err)
	}
	langsJSONFile.Close()
}

func ConstructJSONFile(langs map[string]int) {
	langsJSONStr := LangsMapToJSONStr(langs)
	OutputJSON(langsJSONStr)
}