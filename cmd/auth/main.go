package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/708u/useless-auth-server/internal/auth"
)

func main() {
	s, err := auth.NewServer()
	if err != nil {
		fmt.Fprintf(
			os.Stdout,
			"creating new server failed: %s", err.Error(),
		)
	}

	if http.ListenAndServe(":8080", s.Router); err != nil {
		fmt.Fprintf(os.Stdout, "failed to serve http server: %s", err.Error())
	}
}
