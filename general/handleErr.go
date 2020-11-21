package general

import "log"

//HandleErr calls log.Fatal on err
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
