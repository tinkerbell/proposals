package listener

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const conninfo = "dbname=tinkerbell user=tinkerbell password=tinkerbell sslmode=disable"

// StartListener creates a new dedicated connection for LISTEN/NOTIFY
// and starts listening for events.
func StartListener() {
	_, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Error(err)
	}

	listener := pq.NewListener(conninfo, 10*time.Second, 15*time.Second, errorHandler)
	err = listener.Listen("workflow_changed")
	if err != nil {
		log.Error(err)
	}

	log.Info("starting listener")
	for {
		waitForNotification(listener)
	}
}

func errorHandler(ev pq.ListenerEventType, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func waitForNotification(l *pq.Listener) {
	for {
		select {
		case n := <-l.Notify:
			log.Info("Received data from channel [", n.Channel, "] :")

			// Prepare notification payload for pretty print
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
			if err != nil {
				log.Error(err)
				return
			}
			fmt.Println(string(prettyJSON.Bytes()))
			return
		case <-time.After(10 * time.Second):
			log.Info("Received no events for 90 seconds, checking connection")
			go func() {
				l.Ping()
			}()
			return
		}
	}
}
