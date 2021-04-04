package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/708u/useless-auth-server/internal/client/di"
)

func main() {
	s := di.NewServer()

	if err := http.ListenAndServe(":"+strconv.Itoa(int(s.Config.HTTP.Port)), s.Router); err != nil {
		fmt.Fprintf(os.Stdout, "failed to serve http server: %s", err.Error())
	}
}
