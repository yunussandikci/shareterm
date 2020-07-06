package utils

import (
	"fmt"
	"github.com/yunussandikci/shareterm/common"
	"log"
)

func BuildPasteWebReadURL(fileName string, key string) string {
	return fmt.Sprintf("%s/web/%s?key=%s", common.GetHost(), fileName, key)
}

func PrintStartInfo() {
	log.Printf("\n   ______              ______            \n  / __/ /  ___ _______/_  __/__ ______ _ \n _\\ \\/ _ \\/ _ `/ __/ -_) / / -_) __/  ' \\\n/___/_//_/\\_,_/_/  \\__/_/  \\__/_/ /_/_/_/")
	fmt.Printf("Host: %s\n", common.GetHost())
	fmt.Printf("Shareterm running...")
}