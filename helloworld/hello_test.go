package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloPerson(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloPerson("Talha", "")
		want := "Hello, Talha"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello with empty string", func(t *testing.T) {
		got := HelloPerson("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := HelloPerson("Talha", "Spanish")
		want := "Hola, Talha"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
