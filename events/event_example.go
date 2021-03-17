package events

import (
	"fmt"
	"time"
)

const HELLO_WORLD = "helloWorld"

func main() {
	dispatcher := NewEventDispatcher()
	listener := NewEventListener(myEventListener)
	dispatcher.AddEventListener(HELLO_WORLD, listener)
	time.Sleep(time.Second * 2)
	//dispatcher.RemoveEventListener(HELLO_WORLD, listener)

	dispatcher.DispatchEvent(NewEvent(HELLO_WORLD, nil))

}

func myEventListener(event Event) {
	fmt.Println(event.Type, event.Object, event.Target)
}
