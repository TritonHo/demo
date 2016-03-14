//FIXME: it doesn't allow the config change in runtime, need rethinking

package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

//It will panic if such config value don't exists
//or the value is not an bool
func GetBool(key string) bool {
	if value, ok := getEnv(key); !ok {
		log.Panic(`Environmental variable [` + key + `] don't exists. Please add this environment variable to your environment.`)
	} else {
		if output, err := strconv.ParseBool(value); err != nil {
			log.Panic(`Environmental variable [` + key + `] is not an integer`)
		} else {
			return output
		}
	}
	//should not reach this line
	return false
}

//It will panic if such config value don't exists
//or the value is not an integer
func GetInt(key string) int {
	if value, ok := getEnv(key); !ok {
		log.Panic(`Environmental variable [` + key + `] don't exists. Please add this environment variable to your environment.`)
	} else {
		if output, err := strconv.Atoi(value); err != nil {
			log.Panic(`Environmental variable [` + key + `] is not an integer`)
		} else {
			return output
		}
	}
	//should not reach this line
	return 0
}

//It will panic if such config value don't exists
func GetStr(key string) string {
	value, ok := getEnv(key)
	if !ok {
		log.Panic(`Environmental variable [` + key + `] don't exists`)
	}
	return value
}

//if the config value don't exists, or it is valid, return the defaultValue instead
func GetIntConfigWithDefault(key string, defaultValue int) int {
	if value, ok := getEnv(key); ok {
		if output, err := strconv.Atoi(value); err == nil {
			return output
		}
	}
	return defaultValue
}

//if the config value don't exists, or it is valid, return the defaultValue instead
func GetStrWithDefault(key string, defaultValue string) string {
	value, ok := getEnv(key)
	if ok {
		return value
	}
	return defaultValue
}

var envKeyValues []string = os.Environ()

func getEnv(key string) (string, bool) {
	for _, env := range envKeyValues {
		ss := strings.SplitN(env, "=", 2)
		if len(ss) < 2 {
			//act as sentinel
			continue
		}
		if ss[0] == key {
			return ss[1], true
		}
	}
	return ``, false
}
