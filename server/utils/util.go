package utils

import (
	"fmt"
	"github.com/yunussandikci/shareterm/common"
)

func BuildPasteReadURL(fileName string, key string) string {
	return fmt.Sprintf("%s/%s?key=%s", common.GetHost(), fileName, key)
}
