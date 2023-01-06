package scafolds

import (
	"bytes"
	"strings"

	"github.com/mohsenMj/go-starter-kit/utils"
)

func Apply(scafoldName string, model string) string {
	output := utils.ReadFileByte("app/scafolds/" + scafoldName)
	lower := strings.ToLower(model)
	capital := utils.ToCapital(model)
	output = bytes.Replace(output, []byte("$model"), []byte(lower), -1)
	output = bytes.Replace(output, []byte("$Model"), []byte(capital), -1)
	return string(output)
}
