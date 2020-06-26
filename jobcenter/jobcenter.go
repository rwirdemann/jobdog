package jobcenter

import (
	"fmt"
	"time"

	"github.com/rwirdemann/jobdog/domain"
)

var feeds []domain.Feed

// Publisher abstracts the functions of the exported by the jobcenter package.
type Publisher interface {
	Publish(job domain.Job)
}

func Publish(job domain.Job) {
	for _, f := range feeds {
		post(f.URL, job)
	}
}

type PublisherV1 struct {
}

func (p PublisherV1) Publish(job domain.Job) {
	Publish(job)
}

func post(url string, job domain.Job) {
}

type PublisherV2 struct {
}

var queue []domain.Job

func init() {
	go process()
}

func process() {
	var failed []domain.Job
	for {
		for _, j := range queue {
			if !push(j) {
				failed = append(failed, j)
			}
		}
		time.Sleep(2 * time.Second)
		queue = failed
		failed = nil
	}
}

func (p PublisherV2) Publish(job domain.Job) {
	queue = append(queue, job)
}

func push(job domain.Job) bool {
	fmt.Printf("Pushing job '%s'\n", job.ID)
	return true
}
