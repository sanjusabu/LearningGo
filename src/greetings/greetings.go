package greetings

import (
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no one to greet, please enter a name")
	}
	message := randomGreeting(name)
	return message, nil
}
func randomGreeting(name string) string {
	formats := []string{
		"Hi, " + name + " Welcome!",
		"Great to see you, " + name,
		"Hail," + name + "! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func MultiHello(names []string) (map[string]string, error) {
	maps := make(map[string]string)
	if len(names) == 0 {
		return maps, errors.New("empty names")
	}
	for _, name := range names {
		if name == "" {
			continue
		}
		maps[name], _ = Hello(name)
	}

	return maps, nil
}
