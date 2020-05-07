package sessions

import (
	"github.com/gorilla/sessions"
)

// Store - var for creating cookies
var Store = sessions.NewCookieStore([]byte("my-secret"))
