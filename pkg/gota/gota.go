package gota

import (
	"strings"

	"github.com/go-gota/gota/dataframe"
)

func DataframeFromJSON(json string) dataframe.DataFrame {
	return dataframe.ReadJSON(strings.NewReader(json))
}
