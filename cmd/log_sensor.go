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
// calculation from temp int like:
//        temp_c = float(temp_string) / 1000.0
//        temp_f = temp_c * 9.0 / 5.0 + 32.0
//        return temp_c, temp_f
//
func main() {
	fg := sensor_base_dir + "/28-*/*"
	files, err := filepath.Glob(fg)
	if err != nil {
		log.Println(err)
		log.Fatalf("Error processing file glob [%s]", fg)
	}

	for _, f := range files {
		log.Println("Checking: " + f)
		if strings.HasSuffix(f, "/temperature") {
			for {
				b, err := ioutil.ReadFile(f)
				strings.Replace(string(b), "\n", "", -1)
				log.Println("RAW BYTES: " + string(b))
				if err != nil {
					log.Println("ERROR: ", err)
				}

				bs, err := strconv.ParseInt(string(b[:5]), 10, 32)
				if err != nil {
					log.Println("ERROR: ", err)
				}
				log.Println("TEMP: ", bs)
			}
		}
	}
}
