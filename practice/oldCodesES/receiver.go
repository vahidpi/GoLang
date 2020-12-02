package main

import (
	"io"
	"log"
	"strings"
	"sync/atomic"
	"time"
)

type (
	messageReceiver struct {
		// pi        *processorInput
		// awsClient *awsClient
		queue     string
		mustStop  int64
		timeout   int64
	}
)

func (r *messageReceiver) run() error {
	for atomic.LoadInt64(&r.mustStop) == 0 {
		messages, err := r.awsClient.sqsReceive(r.queue, 1, r.timeout)
		// TODO: retry?
		if err != nil {
			return err
		}

		if len(messages) > 0 {
			// TODO: pick timestamps from actual messages
			timestamp := time.Now()

			// Although we expect ony one message here...
			for _, msg := range messages {
				if err := r.awsClient.sqsDelete(r.queue, msg.ReceiptHandle); err != nil {
					log.Printf("Failed to delete message: %v\n", err)
				}

				msgChan := make(chan io.Reader)

				go func(msgBody chan io.Reader) {
					if err := processMessage(<-msgBody, r.pi, timestamp); err != nil {
						log.Printf("Failed to process message: %v\n", err)
					}
				}(msgChan)

				// Block until goroutine is actually running
				msgChan <- strings.NewReader(*msg.Body)
			}
		}
	}

	return nil
}