package main

import "testing"

func TestHello(t *testing.T) {
	helloTests := []struct {
		name     string
		language string
		greeting string
	}{
		{name: "Chris", language: "English", greeting: "Hello, Chris"},
		{name: "Elodie", language: "Spanish", greeting: "Hola, Elodie"},
		{name: "Francoise", language: "French", greeting: "Bonjour, Francoise"},
	}

	for _, tt := range helloTests {
		t.Run(tt.language, func(t *testing.T) {
			got := Hello(tt.name, tt.language)
			if got != tt.greeting {
				t.Errorf("%s got %s want %s", tt.language, got, tt.greeting)
			}
		})
	}
}
