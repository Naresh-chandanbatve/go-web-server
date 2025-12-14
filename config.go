package main

import (
	"os"
	"strings"
)

type Config struct {
	ListenPort string
	Root       string
	Proxies    map[string]string
}

func LoadConfig(path string) Config {
	cfg := Config{
		ListenPort: "8080",
		Root:       "./static",
		Proxies:    make(map[string]string),
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)

		switch parts[0] {
		case "listen":
			cfg.ListenPort = parts[1]
		case "root":
			cfg.Root = parts[1]
		case "proxy":
			cfg.Proxies[parts[1]] = parts[2]
		}
	}

	return cfg
}
