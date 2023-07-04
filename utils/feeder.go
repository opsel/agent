package utils

import (
	"agent/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Feeder(cfg config.Config, Endpoint string, Data interface{}) error {

	/**
	* JSON Marshal payload before sending it to the upstream
	* and then pass the byte array to net/http post request
	 */
	payload, err := json.Marshal(Data)
	if err != nil {
		return err
	}

	/**
	* Initiate new HTTP post request using net/http to feed
	* data into the upstream gather endpoint
	 */
	Request, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/gather%s", cfg.Server.Schema, cfg.Server.URI, Endpoint), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	/**
	* Configure relevent header values before sending the request
	* to upstream
	 */
	Request.Header.Set("User-Agent", "OpselAgent/0.1-beta (+https://opsel.github.io)")
	Request.Header.Set("Opsel-Agent-ID", cfg.Agent.ID)

	/**
	* Dispatch the actual request using the net/http and wait for the
	* return and pass the error back to handler function
	 */
	client := &http.Client{Timeout: time.Second * 5}
	Response, err := client.Do(Request)
	if err != nil {
		return err
	}
	defer Response.Body.Close()

	// RETURN
	return nil
}
