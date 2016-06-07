// Can be launched with:
//   ./xmpp_jukebox -jid=test@localhost/jukebox -password=test -address=localhost:5222
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/processone/gox/xmpp"
	"github.com/processone/gox/xmpp/iot"
	"github.com/processone/mpg123"
	"github.com/processone/soundcloud"
)

// Get the actual song Stream URL from SoundCloud website song URL and play it with mpg123 player.
const scClientID = "ReplaceWithYourSoundCloudID"

func main() {
	jid := flag.String("jid", "", "jukebok XMPP JID, resource is optional")
	password := flag.String("password", "", "XMPP account password")
	address := flag.String("address", "", "If needed, XMPP server DNSName or IP and optional port (ie myserver:5222)")
	flag.Parse()

	var client *xmpp.Client
	var err error
	if client, err = connectXmpp(*jid, *password, *address); err != nil { // HL
		log.Fatal("Could not connect to XMPP: ", err)
	}

	p, err := mpg123.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}

	// Iterator to receive packets coming from our XMPP connection
	for packet := range client.Recv() { // HL
		switch packet := packet.(type) {
		case *xmpp.ClientMessage: // HL
			processMessage(client, p, packet)
		case *xmpp.ClientIQ: // HL
			processIq(client, p, packet)
		case *xmpp.ClientPresence:
			// Do nothing with received presence
		default:
			fmt.Printf("Ignoring packet: %T\n", packet)
		}
	}
	// STOPFOR OMIT
}

func processMessage(client *xmpp.Client, p *mpg123.Player, packet *xmpp.ClientMessage) {
	command := strings.Trim(packet.Body, " ")
	if command == "stop" {
		p.Stop()
	} else {
		playSCURL(p, command)
	}
}
// STOPMSG OMIT

func processIq(client *xmpp.Client, p *mpg123.Player, packet *xmpp.ClientIQ) {
	switch payload := packet.Payload.(type) {
	// We support IOT Control IQ
	case *iot.ControlSet:
		var url string
		for _, element := range payload.Fields {
			if element.XMLName.Local == "string" && element.Name == "url" {
				url = strings.Trim(element.Value, " ")
				break
			}
		}

		playSCURL(p, url)
		setResponse := new(iot.ControlSetResponse)
		reply := xmpp.ClientIQ{Packet: xmpp.Packet{To: packet.From, Type: "result", Id: packet.Id}, Payload: setResponse}
		client.Send(reply.XMPPFormat())
	default:
		fmt.Printf("Other IQ Payload: %T\n", packet.Payload)
	}
}
// STOPIQ OMIT

func playSCURL(p *mpg123.Player, rawURL string) {
	songID, _ := soundcloud.GetSongID(rawURL)
	url := soundcloud.FormatStreamURL(songID)
	p.Play(url)
}

func connectXmpp(jid string, password string, address string) (client *xmpp.Client, err error) {
	xmppOptions := xmpp.Options{Address: address,
		Jid: jid, Password: password, PacketLogger: os.Stdout,
		Retry: 10}

	if client, err = xmpp.NewClient(xmppOptions); err != nil {
		return
	}

	if _, err = client.Connect(); err != nil {
		return
	}
	return
}
// STOPCONNECT OMIT

// TODO
// - Have a player API to play, play next, or add to queue
// - Have the ability to parse custom packet to play sound
// - Use PEP to display tunes status
// - Ability to "speak" messages
