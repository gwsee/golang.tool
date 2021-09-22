package easy

import (
	"encoding/json"
	"encoding/xml"
	"mahonia"
	"strings"
)

func ReadText(by []byte, language string) (res []map[string]string) {
	var title []string
	arr := strings.Split(string(by), "\n")
	enc := mahonia.NewDecoder(language)
	for k, v := range arr {
		if v == "" {
			continue
		}
		if k == 0 {
			title = strings.Split(v, "\t")
			continue
		}
		item := strings.Split(v, "\t")
		m := map[string]string{}
		for k1, v1 := range title {
			if enc == nil {
				m[v1] = item[k1]
			} else {
				m[enc.ConvertString(v1)] = enc.ConvertString(item[k1])
			}

		}
		res = append(res, m)
	}
	return
}

// TransMapToStruct map 转结构体
func TransMapToStruct(mp []map[string]string, res interface{}) (err error) {
	arr, err := json.Marshal(mp)
	if err != nil {
		return
	}
	err = json.Unmarshal(arr, res)
	return
}

// TransXmlToStruct xml 转结构体
func TransXmlToStruct(by []byte, res interface{}) (err error) {
	err = xml.Unmarshal(by, res)
	return
}
