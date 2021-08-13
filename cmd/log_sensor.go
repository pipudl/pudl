package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	sensor_base_dir = "/sys/bus/w1/devices"
)

// sudo apt install golang
// file is at: /sys/bus/w1/devices/28-<12-byte hash>/
func main() {
	fg := sensor_base_dir + "/28-*"
	files, err := filepath.Glob(fg)
	if err != nil {
		log.Println(err)
		log.Fatalf("Error processing file glob [%s]", fg)
	}

	for _, f := range files {
		if strings.HasSuffix(f, "/temperature") {
			for {
				b, err := ioutil.ReadFile(f)
				if err != nil {
					log.Println("ERROR: ", err)
				}
				bs, _ := strconv.ParseInt(string(b), 10, 32)
				log.Println("TEMP: ", bs)
			}
		}
	}
}
