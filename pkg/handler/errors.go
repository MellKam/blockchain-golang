package handler

import "log"

func HandlePossibleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
