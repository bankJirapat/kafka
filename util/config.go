package util

import (
	"bufio"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

func ReadConfig(path string) (kafka.ConfigMap, error) {
	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(path)
	if err != nil {
		logrus.Warnf("Failed to open config file: %s", err)
		os.Exit(1)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			before, after, found := strings.Cut(line, "=")
			if found {
				parameter := strings.TrimSpace(before)
				value := strings.TrimSpace(after)
				m[parameter] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		logrus.Warnf("Failed to read config file: %s", err)
		os.Exit(1)
		return nil, err
	}

	return m, nil
}
