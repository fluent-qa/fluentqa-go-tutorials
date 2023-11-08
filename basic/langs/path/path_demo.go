package path

import (
	"fmt"
	"path/filepath"
)

func GetFileExtension() {
	filePath := "a.json"
	filePostFix := filepath.Ext(filePath)
	fmt.Println(filePostFix)
}
