package main

import (
	"fmt"

	"github.com/processone/gomqtt/mqtt"
	"github.com/processone/gomqtt/mqtt/packet"
)

// START1 OMIT
func main() {
	messages := make(chan *mqtt.Message) // HL
	client := mqtt.New("localhost:1883", messages) // HL
	client.ClientID = "MQTT-Sub"

	if err := client.Connect(); err != nil { // HL
		fmt.Printf("Connection error: %q\n", err)
		return
	}

	name := "mremond-osx/cputemp"
	topic := packet.Topic{Name: name, QOS: 1}
	client.Subscribe(topic) // HL

	for m := range messages { // HL
		fmt.Printf("Received message on topic %s: %+v\n", m.Topic, m.Payload)
	}
}
// STOP1 OMIT
