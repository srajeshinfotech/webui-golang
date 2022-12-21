package core

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func dummyAPI(w http.ResponseWriter, r *http.Request) {
	v := `
	{
		"list" : [
			{
				"name" : "milan",
				"office" : "elisity-wework-1",
				"position" : "head"
			},
			{
				"name" : "sarath",
				"office" : "elisity-wework-2",
				"position" : "lead"
			},
			{
				"name" : "rajesh",
				"office" : "elisity-wework-3",
				"position" : "lead"
			}
		]
	}
	`
	fmt.Fprintf(w, v)
}

func StartWebServer(s_path string) {

	log.Info("Starting WebUI Server..")

	http.Handle("/", http.FileServer(http.Dir(s_path)))
	http.HandleFunc("/api/emp", dummyAPI)
	for {
		if err := http.ListenAndServe("127.0.0.1:9000", nil); err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}
}
