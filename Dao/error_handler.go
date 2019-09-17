package Dao

import (
	"log"
)

func errorHandler(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
