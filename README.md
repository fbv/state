# state

State management library

## Usage

```go

import "github.com/fbv/state"

// define secret key
const secret = "top_secret"

// generate state and send it to thrid party
func sendRedirect(w http.ResponseWriter, r *http.Request) {
    stateString := state.Generate(secret)
    url := config.AuthCodeURL(stateString)
    log.Println("Redirect to", url)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// check responce with a state
func handleResponse(w http.ResponseWriter, r *http.Request) {
    stateString := r.FormValue("state")
    if !state.Valid(stateString, secret) {
        log.Printf("Invalid state\n")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
    // state is valid ...
}

```

See [example](example/example.go)
