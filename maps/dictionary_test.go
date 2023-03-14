package main

import "testing"

func Test_検索(t *testing.T) {
	t.Parallel()

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("存在する単語", func(t *testing.T) {
		t.Parallel()

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("存在しない単語", func(t *testing.T) {
		t.Parallel()

		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func Test_追加(t *testing.T) {
	t.Parallel()

	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	t.Run("新規追加", func(t *testing.T) {
		t.Parallel()

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run(`既存追加`, func(t *testing.T) {
		t.Parallel()

		dictionary.Add(word, definition)

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func Test_更新(t *testing.T) {
	t.Parallel()

	word := "test"
	definition := "this is just a test"

	t.Run(`更新`, func(t *testing.T) {
		t.Parallel()

		dictionary := Dictionary{word: definition}
		newDefinition := "new newDefinition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run(`新規更新`, func(t *testing.T) {
		t.Parallel()

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func Test_削除(t *testing.T) {
	t.Parallel()

	word := "test"
	definition := "test definition"

	t.Run(`存在するキーの削除`, func(t *testing.T) {
		t.Parallel()

		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
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
		if want == nil {
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
