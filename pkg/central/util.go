package central

import "io"

func closeMe(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Error("Unable to close %v", err.Error())
	}
}
