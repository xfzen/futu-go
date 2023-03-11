package utils

import (
	"fmt"

	"github.com/hokaccha/go-prettyjson"
)

func PrettyJson(v interface{}) string {
	s, _ := prettyjson.Marshal(v)
	return string(s)
}

func DumpSlice(title string, list []float64) {
	fmt.Printf("%v:", title)

	nlen := len(list)
	for i := (nlen - 10); i < len(list); i++ {
		fmt.Printf("%10.5f", list[i])
	}

	fmt.Printf("\n")
}

// 用于组按条件模糊查询的SQL语句中
func ListMap2Str(mapVue map[string]string) string {
	str := ""
	if len(mapVue) == 0 {
		str = "1=1" // 加此条件为了sql条件where不报错
		return str
	}

	for field := range mapVue {
		if field == "id" || field == "dept_id" || field == "roomId" || field == "station_no" || field == "category" {
			str = str + "`" + field + "` = "
			str = str + "'" + mapVue[field] + "' and "
		} else {
			str = str + "`" + field + "` like "
			str = str + "'%" + mapVue[field] + "%' and "
		}
	}

	str = string([]byte(str)[:len(str)-4])
	return str
}
