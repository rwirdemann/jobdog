package main

import (
	"github.com/rwirdemann/jobdog/domain"
	"github.com/rwirdemann/jobdog/jobcenter"
)

func main() {
	stop := make(chan bool)
	var jc jobcenter.Publisher = jobcenter.PublisherV2{}
	jc.Publish(domain.Job{ID: "1"})
	jc.Publish(domain.Job{ID: "2"})
	<-stop
}
