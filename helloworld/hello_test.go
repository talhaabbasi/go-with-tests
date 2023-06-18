package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloPerson(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloPerson("Talha")
		want := "Hello, Talha"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
