---
title: "Cloud Infrastructure as Code (infra)"
description: "Declare, plan, and provision cloud resources directly in Go with packages/cloud/infra, deterministic execution plans, and zero external DSL bloat."
date: "2026-10-08"
---

# Cloud Infrastructure as Code (`infra`)

The **`infra`** workload allows engineering teams to define cloud topology, container orchestration, networking, and storage directly using **typed Go structures** instead of disjointed HCL or complex YAML templates.

Powered by `packages/cloud/infra`, the engine compiles Go declarations into deterministic execution plans and provisions cloud targets safely.

---

## Architectural Characteristics

1. **Go as the Unified Language:** Refactor and test your infrastructure using standard Go compiler guarantees, type checking, and unit testing (`go test`).
2. **Deterministic Execution Plans:** Run `kiw deploy --target infra --plan` to preview added, modified, or destroyed cloud resources before any mutations take place.
3. **Provider-Agnostic Topologies:** Abstract declarations for compute (containers, serverless), networking (VPC, load balancers), and persistence (PostgreSQL, object stores).
4. **State Integrity:** Stores encrypted state snapshots in remote backends (AWS S3, Google Cloud Storage, or Kubernetes Secrets).

---

## Scaffolding an `infra` Project

```bash
kiw new cloud-env --infra
cd cloud-env
```

Directory layout:

```text
cloud-env/
├── krewire.yaml          # Project configuration
├── go.mod
└── main.go               # Infrastructure topology definition
```

---

## Topology Declaration (`main.go`)

```go
package main

import (
	"context"
	"os"

	"github.com/krewire/krewire/packages/cloud/infra"
	"github.com/krewire/krewire/packages/cloud/infra/providers/aws"
)

func main() {
	stack := infra.NewStack("production-ecommerce")

	// Declare VPC
	network := aws.NewVPC(stack, "ecommerce-vpc", aws.VPCConfig{
		CIDRBlock: "10.0.0.0/16",
		Subnets: []string{
			"10.0.1.0/24", // public
			"10.0.2.0/24", // private
		},
	})

	// Declare Managed PostgreSQL Database
	db := aws.NewRDSInstance(stack, "ecommerce-db", aws.RDSConfig{
		VPC:           network,
		Engine:        "postgres",
		Version:       "16.2",
		InstanceClass: "db.t4g.medium",
		AllocatedGB:   50,
	})

	// Declare Container Service
	aws.NewECSCluster(stack, "ecommerce-compute", aws.ECSConfig{
		VPC:         network,
		Database:    db,
		DesiredTasks: 4,
		CPU:         "512",
		Memory:      "1024",
	})

	// Execute Plan or Apply based on CLI flags
	if err := infra.Execute(context.Background(), stack); err != nil {
		os.Exit(1)
	}
}
```

---

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "cloud-env"
  kind: infra
  version: "0.1.0"

infra:
  provider: "aws" # aws, gcp, kubernetes
  region: "ap-southeast-1"
  state:
    backend: "s3"
    bucket: "krewire-infra-state"
    lock_table: "krewire-locks"
```

---

## CLI Commands for Infrastructure

### Previewing Changes (Dry Run Plan)

```bash
kiw deploy --target infra --plan
```

*Output:*
```text
Plan: 3 to add, 0 to change, 0 to destroy.
+ aws_vpc.ecommerce-vpc (10.0.0.0/16)
+ aws_rds_instance.ecommerce-db (postgres 16.2)
+ aws_ecs_cluster.ecommerce-compute (4 tasks)
```

### Applying Changes

```bash
kiw deploy --target infra
```
