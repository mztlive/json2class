package language

import (
	"os"
	"strings"
)

//下划线转小驼峰
func toCamelCase(s string, ucFirst bool) string {
	names := strings.Split(s, "_")
	for i := 0; i < len(names); i++ {
		if i == 0 && ucFirst == false {
			continue
		}

		names[i] = strings.Title(names[i])
	}

	return strings.Join(names, "")
}

func createFloder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}

	return nil
}

func openFile(path string) (*os.File, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Create(path)
	}

	return os.OpenFile(path, os.O_RDWR|os.O_TRUNC, os.ModePerm)
}
