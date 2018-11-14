package main

import (
	"log"
	"net"
	"regexp"
	"strings"
)

const NUM_FIELDS int = 6

var Address string
var Port string
var Online bool            // online or offline?
var Version string         // server version
var Motd string            // message of the day
var Current_players string // current number of players online
var Max_players string     // maximum player capacity

func Init(given_address string, given_port string) {
	Address = given_address
	Port = given_port
	// ToDo: Add timeout
	conn, err := net.Dial("tcp", Address+":"+Port)
	if err != nil {
		Online = false
		return
	}

	_, err = conn.Write([]byte("\xFE\x01"))
	if err != nil {
		Online = false
		return
	}

	raw_data := make([]byte, 512)
	_, err = conn.Read(raw_data)
	if err != nil {
		Online = false
		return
	}
	conn.Close()

	if raw_data == nil || len(raw_data) == 0 {
		Online = false
		return
	}

	data := strings.Split(string(raw_data[:]), "\x00\x00\x00")
	if data != nil && len(data) >= NUM_FIELDS {
		Online = true
		Version = data[2]
		Motd = data[3]
		Current_players = data[4]

		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		Max_players = reg.ReplaceAllString(data[5], "")
	} else {
		Online = false
	}
}
