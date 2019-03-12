package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ensureDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func findMygrationsInDir(path string) []string {
	res, _ := filepath.Glob(path + string(os.PathSeparator) + "*_*.go")
	return res
}

func parseIDAndName(path string) (int, string, error) {
	_, filename := filepath.Split(path)
	tmp := strings.TrimSuffix(filename, filepath.Ext(filename))

	index := strings.Index(tmp, "_")
	if index == -1 {
		return 0, "", fmt.Errorf("file is not in a valid format '%s'", filename)
	}

	id, err := strconv.Atoi(tmp[:index])
	if err != nil {
		return 0, "", fmt.Errorf("file is not in a valid format '%s'", filename)
	}

	return id, tmp[index+1:], nil
}
