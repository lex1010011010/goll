package goll

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

var projectBasePath string

type CustomJSONFormatter struct {
	OrderedFields []string
}

func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	getPath()
	pc, file, line, ok := runtime.Caller(8)
	if !ok {
		return nil, fmt.Errorf("error during runtime.Caller")
	}
	funcName := runtime.FuncForPC(pc).Name()
	relativePath := strings.TrimPrefix(file, projectBasePath)

	entry.Data["file"] = relativePath
	entry.Data["func"] = fmt.Sprintf("%s:%d", funcName, line)
	entry.Data["msg"] = entry.Message
	entry.Data["level"] = entry.Level.String()
	entry.Data["time"] = entry.Time.Format(time.RFC3339)

	var serialized []byte
	serialized = append(serialized, '{')

	first := true
	for _, key := range f.OrderedFields {
		if value, ok := entry.Data[key]; ok {
			if !first {
				serialized = append(serialized, ',')
			} else {
				first = false
			}

			keySerialized, err := json.Marshal(key)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal key to JSON, %v", err)
			}

			valueSerialized, err := json.Marshal(value)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal value to JSON, %v", err)
			}

			serialized = append(serialized, keySerialized...)
			serialized = append(serialized, ':')
			serialized = append(serialized, valueSerialized...)
		}
	}

	serialized = append(serialized, '}', '\n')
	return serialized, nil
}

func getPath() {
	var err error
	projectBasePath, err = os.Getwd()
	if err != nil {
		fmt.Printf("failed to get project path, %v", err)
	}
}
