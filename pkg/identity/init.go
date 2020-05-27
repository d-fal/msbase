package identity

import (
	"fmt"
	"log"
	"regexp"

	"msbase/pkg/conf"
)

func init() {
	fmt.Println("identity init")
	var err error
	Init()

	r, err = regexp.Compile(conf.GetConfigObject().GetEventReceiver().SearchedPhrase)

	if err != nil {
		log.Println("Cannot run regex")
	}
	trim, err = regexp.Compile("[[:space:]]")

	if err != nil {
		log.Println("Cannot setup trimmer")
	}

}

// Init initiates config readers
func Init() {

}
