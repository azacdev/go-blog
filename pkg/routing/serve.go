package routing

import (
	"log"
)

func Serve() {
	r := GetRouter()

	err := r.Run()

	if err != nil {
		log.Fatal("Error in routing")
		return
	}

}
