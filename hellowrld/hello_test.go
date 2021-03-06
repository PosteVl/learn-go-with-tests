package main

import "testing"

func TestHello(t *testing.T){
    assertCorrectMessage := func(t testing.TB, got, want string){
        // tells test to report failing line from subtest not helper function
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }
    t.Run("saying hello to people", func(t *testing.T){
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })
    t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T){
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
    t.Run("in Spanish", func(t *testing.T){
        got := Hello("Elodie", "spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
    t.Run("in French", func(t *testing.T){
        got := Hello("Pierre", "french")
        want := "Bonjour, Pierre"
        assertCorrectMessage(t, got, want)
    })
}
