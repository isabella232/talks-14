package xmpp

import (
	"log"
	"os"

	"github.com/processone/gox/xmpp"
)

func Connect() string {
	options := xmpp.Options{Address: "test1.process-one.net:5222",
		Jid: "mremond@test1.process-one.net",
		Password: "password",
		PacketLogger: os.Stdout,
	}

	var client *xmpp.Client
	var err error
	if client, err = xmpp.NewClient(options); err != nil {
		log.Fatal("Error: ", err)
	}

	var session *xmpp.Session
	if session, err = client.Connect(); err != nil {
		log.Fatal("Error: ", err)
	}
	return session.StreamId
}
// STOP1 OMIT
