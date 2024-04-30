package client

import (
	"context"
	http "github.com/sabrina-djebbar/spelling-app-backend/lib/shttp/client"
)

const (
	CreateSpellingWordPath         = "/create_spelling_word"
	CreateSpellingSetPath          = "/create_spelling_set"
	ListSpellingWordsPath          = "/list_spelling_words"
	ListSpellingSetsPath           = "/list_spelling_sets"
	EditSpellingSetPath            = "/edit_spelling_set"
	EditSpellingSetWordsPath       = "/edit_spelling_set_words"
	CreateSpellingAttemptPath      = "/create_spelling_attempt"
	ListSpellingExerciseByUserPath = "/list_spelling_exercise_by_user"
)

type Client interface {
	CreateSpellingWord(ctx context.Context, req CreateSpellingWordRequest) (CreateSpellingWordResponse, error)
	CreateSpellingSet(ctx context.Context, req CreateSpellingSetRequest) (CreateSpellingSetResponse, error)
	ListSpellingSets(ctx context.Context, req ListSpellingSetsRequest) (ListSpellingSetResponse, error)
}

type client struct {
	internal *http.InternalClient
}

func New() *client {
	cfg := http.InternalClientOptions{

		Host:    "http://user",
		Timeout: 5,
	}

	/* s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	} */

	return &client{internal: http.NewInternalClient(cfg)}
}
