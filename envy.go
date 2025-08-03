package envy

import (
	"bufio"
	"os"
	"strings"
)

// wrapper around the environment variables map
// to obtain an instance use the envy.Instance method
// to access a variable use the Get(string) method
type Envy struct {
	variables map[string]string
}

func (e *Envy) Get(name string) string {
	return e.variables[name]
}

var instances map[string]*Envy = make(map[string]*Envy)

// returns the singleton instance of the Environment struct
// if not previously initialised, the Environment struct is constructed by reading from the provided envFileLocation file
// if envFileLocation is an empty string a .env file is searched the the '.' dir the program is ran from
// the parsing is performed using a scanner so that each line is only processed once
func GetInstance(envFileLocation string) *Envy {
	if instances[envFileLocation] == nil {
		instances[envFileLocation] = initialise(envFileLocation)
	}
	return instances[envFileLocation]
}

func initialise(envFileLocation string) *Envy {
	if envFileLocation == "" {
		envFileLocation = ".env"
	}

	envy := &Envy{
		variables: make(map[string]string),
	}

	file, err := os.Open(envFileLocation)
	if err != nil {
		if os.IsNotExist(err) {
			return envy
		}
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, "\"'")

		envy.variables[key] = value
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return envy
}
