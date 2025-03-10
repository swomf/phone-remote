package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bendahl/uinput"
)

type keyboardInput struct {
	Key int
}

var epickeyboard uinput.Keyboard

var allowedKeys []int

func press(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var t keyboardInput
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	switch t.Key {
	case uinput.KeyPagedown:
		fallthrough
	case uinput.KeyPageup:
		epickeyboard.KeyPress(t.Key)
	case uinput.KeyUp:
		fallthrough
	case uinput.KeyDown:
		for i := 0; i < 3; i++ {
			epickeyboard.KeyPress(t.Key)
		}
	default:
		fmt.Fprintf(w, "Evil man.")
	}

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}

// alternatively (to use specific version), use this:
// import "gopkg.in/bendahl/uinput.v1"
func main() {
	allowedKeys = []int{uinput.KeyDown, uinput.KeyUp, uinput.KeyPagedown, uinput.KeyPageup}

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("testkeyboard"))
	if err != nil {
		return
	}
	epickeyboard = keyboard
	defer keyboard.Close()
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/press", press)

	http.ListenAndServe(":8000", nil)
}
