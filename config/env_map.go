package config

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// NewEnvMap to read the env variables
func NewEnvMap(filename string) (EnvMap, error) {
	var (
		envMap EnvMap
		err    error
	)

	if filename != "" {
		if envMap.fileEnvMap, err = envMap.ParseFileEnv(filename); err != nil {
			return envMap, err
		}
	} else {
		envMap.fileEnvMap = map[string]string{}
	}

	envMap.osEnvMap = envMap.ParseOSEnv()

	return envMap, nil
}

type EnvMap struct {
	osEnvMap   map[string]string
	fileEnvMap map[string]string
}

func (p EnvMap) String(key string) string {
	val, ok := p.osEnvMap[key]
	if ok {
		return val
	}

	return p.fileEnvMap[key]
}

func (p EnvMap) Int(key string) (int, error) {
	val, ok := p.osEnvMap[key]
	if ok {
		return strconv.Atoi(val)
	}

	val = p.fileEnvMap[key]

	return strconv.Atoi(val)
}

func (p EnvMap) MustInt(key string) int {
	val, _ := p.Int(key)

	return val
}

func (p EnvMap) Strings(key string, sep ...string) []string {
	eSep := ","
	if len(sep) > 0 {
		eSep = sep[0]
	}

	val, ok := p.osEnvMap[key]
	if ok {
		return strings.Split(val, eSep)
	}

	val = p.fileEnvMap[key]

	return strings.Split(val, eSep)
}

// ParseOSEnv for product environment
func (p EnvMap) ParseOSEnv() map[string]string {
	environ := os.Environ()

	m := make(map[string]string, len(environ))

	for _, envValue := range environ {
		envPair := strings.SplitN(envValue, "=", 2)
		m[envPair[0]] = strings.TrimSpace(envPair[1])
	}

	return m
}

// ParseFileEnv parse local file for dev and debug
func (p EnvMap) ParseFileEnv(filename string) (map[string]string, error) {
	sf, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.New("open env file failed, err:" + err.Error())
	}
	defer sf.Close()

	m := make(map[string]string, 16)

	scanner := bufio.NewScanner(sf)
	for scanner.Scan() {
		envValue := strings.TrimSpace(scanner.Text())

		if envValue == "" || envValue[0] == '#' {
			continue
		}

		envPair := strings.SplitN(envValue, "=", 2)

		m[envPair[0]] = strings.TrimSpace(envPair[1])

	}

	if err = scanner.Err(); err != nil {
		return nil, errors.New("scan env file failed, err: " + err.Error())
	}

	return m, nil
}
