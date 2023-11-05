package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/acifani/vita/lib/game"
	"github.com/fermyon/spin/sdk/go/v2/redis"
	"github.com/fermyon/spin/sdk/go/v2/variables"
)

const (
	width  = 32
	height = 32
)

func init() {
	redis.Handle(func(payload []byte) error {
		incomingMsg := string(payload)
		fmt.Println("incoming message:", incomingMsg)
		localState := make([]byte, width*height)
		remoteState := make([]byte, width*height)

		stringID, err := variables.Get("local_id")
		if err != nil {
			return errorLogger(err)
		}
		id, err := strconv.Atoi(stringID)
		if err != nil {
			return errorLogger(err)
		}
		key := "universe-" + stringID
		fmt.Println("I am", key)

		senderID, err := strconv.Atoi(incomingMsg[:1])
		if err != nil {
			return errorLogger(err)
		}

		if senderID != id-1 {
			// Skipping message, it's not for us
			fmt.Println("skipping message")
			return nil
		}

		redisHost, err := variables.Get("redis_host")
		if err != nil {
			return errorLogger(err)
		}

		db := redis.NewClient(redisHost)

		universe := game.NewMultiverse(height, width)
		res, err := db.Execute("EXISTS", key)
		if err != nil {
			return errorLogger(err)
		}
		if len(res) != 1 {
			return errorLogger(fmt.Errorf("expected 1 result but got %d", len(res)))
		}
		found, ok := res[0].Val.(int64)
		if !ok {
			return errorLogger(fmt.Errorf("expected int result but got %v of type %v", res[0].Val, reflect.TypeOf(res[0].Val)))
		}

		if found != 1 {
			fmt.Println("a new universe is born")
			universe.Randomize(50)
			universe.Read(localState)
			err := db.Set(key, localState)
			if err != nil {
				return errorLogger(err)
			}
		}

		fmt.Println("fetching existing universe")
		remoteState, err = db.Get(key)
		if err != nil {
			return errorLogger(err)
		}

		universe.Write(remoteState)

		fmt.Println("making contact")
		incomingColumn := make([]uint8, height)
		for idx, cell := range incomingMsg[1:] {
			switch cell {
			case '1':
				incomingColumn[idx] = game.Alive
			case '0':
				incomingColumn[idx] = game.Dead
			}
		}
		universe.MakeContact(incomingColumn)

		fmt.Println("Before")
		fmt.Println(universe)

		universe.Tick()

		fmt.Println("After")
		fmt.Println(universe)

		universe.Read(localState)

		err = db.Set(key, localState)
		if err != nil {
			return errorLogger(err)
		}

		msg := stringID
		for i := uint32(0); i < height; i++ {
			idx := universe.GetIndex(i, width-1)
			msg = msg + strconv.Itoa(int(universe.Cell(idx)))
		}

		fmt.Println("emitting new event")
		fmt.Println(msg)

		db.Publish("universe", []byte(msg))

		return nil
	})
}

// main functiion must be included for the compiler but is not executed.
func main() {}

func errorLogger(err error) error {
	fmt.Fprintf(os.Stderr, "Error: %v", err.Error())
	return err
}
