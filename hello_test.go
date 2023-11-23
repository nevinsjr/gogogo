package main

import "testing"

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		// Arrange
		want := englishHelloPrefix + "John!"

		// Act
		got := Hello("John", "")

		// Assert
		assertMessage(t, got, want)
	})
	t.Run("says 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		// Arrange
		want := "Hello, World!"

		// Act
		got := Hello("", "")

		// Assert
		assertMessage(t, got, want)
	})
	t.Run("greets in Spanish", func(t *testing.T) {
		// Arrange
		want := "Hola, Juan!"

		// Act
		got := Hello("Juan", "Spanish")

		// Assert
		assertMessage(t, got, want)
	})
	t.Run("greets in French", func(t *testing.T) {
		// Arrange
		want := "Bonjour, Jean!"

		// Act
		got := Hello("Jean", "French")

		// Assert
		assertMessage(t, got, want)
	})
}
