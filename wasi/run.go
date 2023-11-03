package main

import (
	"fmt"
	"net/http"

	"github.com/acifani/vita/lib/game"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	kv "github.com/fermyon/spin/sdk/go/key_value"
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

		store, err := kv.Open("default")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer kv.Close(store)

		exists, err := kv.Exists(store, key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !exists {
			universe.Randomize(livePopulation)
			universe.Read(localCopy)
			kv.Set(store, key, localCopy)
		}

		remoteCopy, err := kv.Get(store, key)
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
		kv.Set(store, key, localCopy)
	})
}

func main() {}
