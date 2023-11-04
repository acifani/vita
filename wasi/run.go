package main

import (
	"fmt"
	"net/http"

	"github.com/acifani/vita/lib/game"
	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
	"github.com/fermyon/spin/sdk/go/v2/kv"
)

const (
	width, height  uint32 = 24, 24
	livePopulation        = 75
	key                   = "universe"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		var localCopy = make([]byte, width*height)
		universe := game.NewUniverse(width, height)

		store, err := kv.OpenStore("default")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer store.Close()

		exists, err := store.Exists(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !exists {
			universe.Randomize(livePopulation)
			universe.Read(localCopy)
			store.Set(key, localCopy)
		}

		remoteCopy, err := store.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		drawing := ""
		for idx, cell := range remoteCopy {
			if idx > 0 && idx%int(height) == 0 {
				drawing += "\n"
			}

			if cell == game.Alive {
				drawing += "O "
			} else {
				drawing += ". "
			}
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, drawing)

		universe.Write(remoteCopy)
		universe.Tick()
		universe.Read(localCopy)
		store.Set(key, localCopy)
	})
}

func main() {}
