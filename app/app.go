package app

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	IgnoreCase bool
	FilePath   string
	Query      string
}

func (config *Config) Run() {
	file, err := os.ReadFile(config.FilePath)

	if err != nil {
		fmt.Println("Application error: {}", err.Error())
		return
	}

	for _, line := range strings.Split(string(file), "\n") {
		if strings.Contains(line, config.Query) || (config.IgnoreCase && strings.Contains(strings.ToLower(line), strings.ToLower(config.Query))) {
			fmt.Println(line)
		}
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) BuildConfig() {
	ignoreCase := flag.Bool("ignore-case", false, "ignore case sensitivity")

	flag.Parse()

	var positionalArgs []string
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			flag.CommandLine.Parse([]string{arg})
		} else {
			positionalArgs = append(positionalArgs, arg)
		}
	}
	isFlagPassed := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "ignore-case" {
			isFlagPassed = true
		}
	})

	config.IgnoreCase = *ignoreCase
	if !isFlagPassed {
		config.IgnoreCase = getEnvOrDefault("IGNORE_CASE", false)
	}

	if len(positionalArgs) < 2 {
		fmt.Println("Parsing error: not enough arguments provided")
		return
	}

	config.Query = positionalArgs[0]
	config.FilePath = positionalArgs[1]
}

func getEnvOrDefault(envVar string, defaultValue bool) bool {
	envVal, exists := os.LookupEnv(envVar)
	if !exists {
		return defaultValue
	}
	return envVal == "true"
}
