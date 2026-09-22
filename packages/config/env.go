package config

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func LoadEnv(paths ...string) (map[string]string, error) {
	var err error
	vars := make(map[string]string)
	path_length := len(paths)

	if path_length == 0 {
		return vars, err
	}

	for _, path := range paths {
		temp_vars, err := LoadConfigurationFile(path)

		if err != nil {
			return vars, err
		}

		for key, value := range temp_vars {
			vars[key] = value
		}
	}

	return vars, err
}

func LoadConfigurationFile(filePath string) (map[string]string, error) {
	var err error
	env_vars := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		pattern := regexp.MustCompile(`^\s*(("[^"]*"|'[^']*'|[^"'=:])*)\s*[:=]\s*(.*)$`)
		matches := pattern.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		key, value := matches[1], matches[3]

		env_vars[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return env_vars, err
}
