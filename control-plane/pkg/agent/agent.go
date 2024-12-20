package agent

import "time"

type Agent struct {
	Id        string
	Context   string
	URL       string
	Heartbeat time.Time
}

// here after we establish connection we should recive heartbeats
// and 
