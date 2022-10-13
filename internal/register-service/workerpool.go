package registerservice

import (
	"context"
	httpClient "generate-deveui-cli/internal/http-client"
	"sync"

	log "github.com/sirupsen/logrus"
)

type WorkerPool interface {
	Run(context.Context)
	AddTask(task string)
	GetResults() chan string
}

type workerPool struct {
	maxWorker    int
	queuedDevEui chan string
	results      chan string
	client       httpClient.Client
}

func (wp *workerPool) Run(ctx context.Context) {
	wg := sync.WaitGroup{}
	wg.Add(wp.maxWorker)
	wp.results = make(chan string, len(wp.queuedDevEui))
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int, ctx context.Context, wg *sync.WaitGroup) {
			for {
				select {
				case <-ctx.Done():
					{
						log.Printf("Stopping worker %d gracefully", workerID)
						wg.Done()
						return
					}
				case devEui := <-wp.queuedDevEui:
					{
						log.Printf("worker: %d processing deveui: %s", workerID, devEui)
						err := wp.RegisterToLorawan(devEui)
						if err != nil {
							log.Printf("error in RegisterToLorawan(), err: %s", err.Error())
							continue
						}
						wp.results <- devEui
						log.Printf("worker: %d completed processing deveui: %s", workerID, devEui)
					}
				default:
					{
						wg.Done()
						return
					}
				}
			}
		}(i+1, ctx, &wg)
	}
	wg.Wait()
	close(wp.results)
}

func (wp *workerPool) AddTask(task string) {
	wp.queuedDevEui <- task
}

func (wp *workerPool) GetResults() chan string {
	return wp.results
}

func NewRegisterService(maxWorker int, devEuiQueue chan string, client httpClient.Client) WorkerPool {
	return &workerPool{
		maxWorker:    maxWorker,
		queuedDevEui: devEuiQueue,
		client:       client,
	}
}
