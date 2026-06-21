# Kubernetes Day 1 - Revision Notes

## What Problem Does Kubernetes Solve?

Imagine you have multiple backend services:

* Order Service
* Payment Service
* Inventory Service
* Notification Service

Initially, you deploy them on a few servers and everything works fine.

As traffic grows, several problems appear:

### Problem 1: Scaling

Suppose Order Service receives heavy traffic.

You need:

```text
Order Service → 10 instances
Payment Service → 3 instances
Inventory Service → 5 instances
```

Who creates these new instances?

Doing it manually becomes difficult.

---

### Problem 2: Server Failure

Suppose one server crashes.

An Order Service instance disappears.

Now users start getting errors.

Who notices the failure and starts a replacement instance?

---

### Problem 3: Service Discovery

Payment Service needs to call Order Service.

Order Service might be running on multiple machines.

Its IPs keep changing.

How does Payment Service find it?

---

### Problem 4: Deployments

You want to deploy version 2 of Order Service.

You don't want downtime.

You also don't want to manually update dozens of servers.

---

Docker helps package applications, but Docker alone doesn't solve these operational problems.

Kubernetes was built to solve them.

---

# The Most Important Kubernetes Idea

Kubernetes is based on a simple philosophy:

> Tell me what you want, and I will continuously try to make reality match it.

This is called the **Desired State Model**.

Suppose you tell Kubernetes:

```yaml
replicas: 3
```

You are saying:

> I always want 3 copies of this application running.

This becomes the desired state.

Initially:

```text
Desired Pods = 3
Actual Pods = 3
```

Everything is healthy.

Now one Pod crashes.

```text
Desired Pods = 3
Actual Pods = 2
```

Kubernetes notices the mismatch.

It automatically creates another Pod.

Eventually:

```text
Desired Pods = 3
Actual Pods = 3
```

again.

This continuous correction process is called the:

## Reconciliation Loop

This is the heart of Kubernetes.

---

# Understanding Pods

Most beginners think Kubernetes manages containers.

Technically that's not true.

Kubernetes manages Pods.

A Pod is the smallest deployable unit in Kubernetes.

Most commonly:

```text
Pod
 └── Container
```

For example:

```text
Pod
 └── nginx container
```

Sometimes a Pod may contain multiple containers:

```text
Pod
 ├── Application Container
 └── Logging Sidecar Container
```

Containers inside a Pod:

* Share networking
* Share storage
* Start and stop together

Kubernetes treats them as a single unit.

---

# Understanding Deployments

Pods are temporary.

If a Pod dies, Kubernetes may create a completely new Pod with a different identity.

Because Pods are disposable, we normally don't create Pods directly.

Instead we create a Deployment.

A Deployment says:

```text
I want 3 replicas.
```

Kubernetes then creates and manages the Pods.

Think:

```text
Deployment
      ↓
Creates and manages Pods
```

For example:

```text
Deployment
      ↓
Pod A
Pod B
Pod C
```

If Pod B crashes:

```text
Deployment
      ↓
Pod A
Pod C
```

The Deployment notices:

```text
Desired = 3
Actual = 2
```

and creates:

```text
Pod D
```

bringing the count back to 3.

Notice:

```text
Pod B is not repaired.
Pod D is created.
```

This is an important Kubernetes mindset.

**Pods are cattle, not pets.**

---

# Kubernetes Architecture

A Kubernetes cluster consists of two major parts:

```text
Control Plane
Worker Nodes
```

---

## Control Plane

Think of the Control Plane as the brain.

It contains several important components.

### API Server

The API Server is the front door of Kubernetes.

Every command goes through it.

Example:

```bash
kubectl apply -f deployment.yaml
```

The request first reaches the API Server.

---

### etcd

etcd is the cluster database.

It stores:

* Deployments
* Services
* Pods
* Configuration
* Desired State

Whenever you create a Deployment, Kubernetes stores that information in etcd.

---

### Controller

The Controller continuously compares:

```text
Desired State
vs
Actual State
```

When they don't match, it takes corrective action.

Example:

```text
Desired = 3 Pods
Actual = 2 Pods
```

Controller creates a replacement Pod.

---

### Scheduler

The Scheduler decides:

> Which Node should run this Pod?

Suppose there are three Nodes:

```text
Node A
Node B
Node C
```

Scheduler might decide:

```text
Pod 1 → Node A
Pod 2 → Node B
Pod 3 → Node C
```

based on available resources.

---

## Worker Nodes

Worker Nodes are the machines where applications actually run.

Each Node contains a component called:

### Kubelet

Kubelet is the node agent.

It receives instructions and starts containers using the container runtime.

Think:

```text
Control Plane
      ↓
Run Pod
      ↓
Kubelet
      ↓
Starts Container
```

---

# What Happens When We Run kubectl apply?

When we executed:

```bash
kubectl apply -f deployment.yaml
```

the following happened:

### Step 1

kubectl sent the Deployment definition to the API Server.

### Step 2

The API Server validated it and stored it in etcd.

### Step 3

The Deployment Controller noticed:

```text
Desired Pods = 3
Actual Pods = 0
```

and created three Pod objects.

### Step 4

The Scheduler assigned those Pods to Nodes.

### Step 5

The Kubelet on each assigned Node started the containers.

Eventually:

```text
3 Pods Running
```

---

# What Did We Actually Deploy?

We deployed:

```yaml
image: nginx
replicas: 3
```

This resulted in:

```text
Pod A → nginx
Pod B → nginx
Pod C → nginx
```

Three Pods.

Three nginx containers.

---

# Why Do We Need Services?

Now we have three running Pods.

But another application needs to talk to them.

Suppose the Pods have IPs:

```text
Pod A → 10.0.0.1
Pod B → 10.0.0.2
Pod C → 10.0.0.3
```

Which IP should a client call?

Worse:

If Pod B dies, Kubernetes may create:

```text
Pod D → 10.0.0.10
```

The IP changes.

Clients cannot depend on Pod IPs.

---

# Service: Stable Networking Layer

A Service provides a stable endpoint.

Think:

```text
Service
      ↓
Pod A
Pod B
Pod C
```

Clients call:

```text
hello-app-service
```

instead of Pod IPs.

The Service automatically forwards traffic to healthy Pods.

It also load balances requests.

---

# Labels and Selectors

Services find Pods using labels.

Pods have:

```yaml
labels:
  app: hello-app
```

Service has:

```yaml
selector:
  app: hello-app
```

Kubernetes automatically matches them.

This means:

If replicas increase from:

```text
3 → 10
```

the Service automatically discovers the new Pods.

No configuration change is required.

---

# Scaling

We learned:

```bash
kubectl scale deployment hello-app --replicas=10
```

This changes the desired state.

From:

```text
Desired = 3
```

to:

```text
Desired = 10
```

Controller notices:

```text
Actual = 3
Desired = 10
```

and creates seven additional Pods.

Scheduler assigns Nodes.

Kubelets start containers.

Eventually:

```text
Actual = 10
```

---

# Commands Practiced

```bash
# Create cluster
kind create cluster --name learning-k8s

# Verify cluster
kubectl get nodes

# Create deployment
kubectl apply -f deployment.yaml

# View deployments
kubectl get deployments

# View pods
kubectl get pods

# Delete a pod
kubectl delete pod <pod-name>

# Create service
kubectl apply -f service.yaml

# View services
kubectl get svc

# View service endpoints
kubectl get endpoints

# Scale deployment
kubectl scale deployment hello-app --replicas=10

# Port forward
kubectl port-forward service/hello-app-service 8080:80
```

---

# Complete Mental Model

```text
Deployment
      ↓
Creates Pods
      ↓
Pods Run Containers
      ↓
Service Finds Pods
      ↓
Routes Traffic
```

Internally:

```text
kubectl
      ↓
API Server
      ↓
etcd

Controller
      ↓
Creates Missing Pods

Scheduler
      ↓
Assigns Nodes

Kubelet
      ↓
Starts Containers
```

---

# Interview Definition

> Kubernetes is a container orchestration platform that automates deployment, scaling, service discovery, load balancing, and self-healing of containerized applications using a desired-state reconciliation model.
