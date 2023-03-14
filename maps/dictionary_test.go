package main

import "testing"

func Test_検索(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("存在する単語", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("存在しない単語", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func Test_追加(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	t.Run("新規追加", func(t *testing.T) {
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run(`既存追加`, func(t *testing.T) {
		dictionary.Add(word, definition)

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}

	if got == nil {
		if want != nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
