---
title: "Background Worker (worker)"
description: "Asynchronous job processing, recurring cron schedules, retry policies with backoff, and graceful shutdown in Go with packages/cloud/worker."
date: "2026-10-08"
---

# Background Worker (`worker`)

The **`worker`** workload is designed for running long-lived background daemons, asynchronous job consumers, and scheduled cron workflows.

Powered by `packages/cloud/worker` and `packages/resilience`, workers operate reliably under high load with automatic retry policies, exponential backoff, dead-letter queues (DLQ), and graceful shutdown handling.

---

## Architectural Characteristics

1. **Pluggable Queue Backends:** Standard queue interfaces with drivers for in-memory channels (for testing/development), Redis, NATS, and AWS SQS.
2. **Resilience & Retry Policies:** Configurable max attempts, jittered exponential backoff, and DLQ dispatch for poison pill jobs.
3. **Cron Scheduler:** Precision scheduled job triggers with standard 5-part cron syntax.
4. **Clean OS Signal Handling:** Drains in-flight jobs gracefully when receiving `SIGINT` or `SIGTERM`, preventing data corruption.

---

## Scaffolding a `worker` Project

```bash
kiw new email-worker --worker
cd email-worker
```

Directory structure:

```text
email-worker/
├── krewire.yaml          # Project configuration
├── go.mod
├── cmd/
│   └── worker/
│       └── main.go       # Worker entry point
└── internal/
    ├── jobs/             # Job payload definitions & handlers
    └── services/         # Third-party email API clients
```

---

## Worker Implementation

```go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/krewire/krewire/packages/cloud/worker"
	"github.com/krewire/krewire/packages/resilience"
)

type SendEmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Initialize the worker dispatcher
	dispatcher := worker.NewDispatcher(worker.Config{
		Concurrency: 10,
		Logger:      logger,
	})

	// Register a job handler with retry policy
	dispatcher.RegisterHandler("send_email", func(ctx context.Context, job worker.Job) error {
		var payload SendEmailPayload
		if err := job.Decode(&payload); err != nil {
			return err
		}

		logger.Info("Processing email job", "recipient", payload.Recipient)

		// Simulate business logic
		if payload.Recipient == "" {
			return fmt.Errorf("invalid recipient")
		}

		return nil
	}, worker.WithRetryPolicy(resilience.RetryPolicy{
		MaxAttempts: 3,
		InitialBackoff: 500 * time.Millisecond,
		MaxBackoff:     5 * time.Second,
	}))

	// Register a recurring cron task (every 5 minutes)
	dispatcher.RegisterCron("*/5 * * * *", func(ctx context.Context) error {
		logger.Info("Running 5-minute health check")
		return nil
	})

	// Start consuming tasks until OS interrupt
	if err := dispatcher.Run(context.Background()); err != nil {
		logger.Error("Worker terminated with error", "err", err)
		os.Exit(1)
	}
}
```

---

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "email-worker"
  kind: worker
  version: "0.1.0"

worker:
  concurrency: 8
  queue: "redis://127.0.0.1:6379/0"
  dead_letter_queue: "failed_jobs"
```

---

## Running Workers

Run the worker daemon locally:

```bash
kiw worker
```

Run with custom concurrency or queue:

```bash
kiw worker --concurrency 16
```
