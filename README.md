# 📦 WebsiteGoApplication – Microservices-Based Order & Payment System

A Go-based cloud-native microservices system that handles **order processing and payment management** using **Google Pub/Sub**, **JWT authentication**, and **Cloud SQL**. This project demonstrates the shift from a monolith to a decoupled, event-driven architecture.

---

## 🎯 Objective

To build and deploy a secure, scalable, and event-driven microservices system using:

* **Orders and Payment MicroServices**
* **JWT-based authentication**
* **Cloud Pub/Sub messaging**
* **Cloud SQL (MySQL)**
* **Service-to-service communication**

---

## 🏗️ Architecture Overview

```
[ User ] 
   ↓   (JWT Auth)
[ Order Service ] 
   ↓   (Publishes)
[ Google Pub/Sub Topic: payment-events ]
   ↓   (Subscribed by)
[ Payment Service ]
   ↓   (Publishes status to)
[ Google Pub/Sub Topic: order-updates ]
   ↓
[ Order Service updates order status ]
```

---

## 📦 Services

### 🧾 Order Service

Responsible for order lifecycle management.

#### Endpoints:

| Method  | Endpoint              | Description         |
| ------- | --------------------- | ------------------- |
| `POST`  | `/orders`             | Create a new order  |
| `GET`   | `/orders/{id}`        | Fetch order by ID   |
| `PATCH` | `/orders/{id}/status` | Update order status |

🔒 All endpoints require **JWT authentication**.

When an order is created:

* It is stored in the database with status `PENDING`
* A message is **published to `payment-events` Pub/Sub topic**

---

### 💳 Payment Service

Listens to `payment-events` topic.

#### Responsibilities:

* Process incoming order messages
* Simulate or process real payment logic
* If successful, publish a message with updated status to `order-updates` topic

---

## 🔐 Authentication

* JWT-based authentication using `/users/login` and `/users` endpoints.
* JWT token is required for accessing order-related APIs.

#### User API:

| Method | Endpoint       | Description                        |
| ------ | -------------- | ---------------------------------- |
| `POST` | `/users`       | Register a new user                |
| `POST` | `/users/login` | Authenticate and receive JWT token |

---

## ☁️ Google Cloud Setup

### ✅ Topics:

* `payment-events` → Triggered when an order is created.
* `order-updates` → Used by payment service to send back payment results.

### ✅ Subscriptions:

* `payment-service-sub` → Subscribes to `payment-events`
* `order-service-sub` → Subscribes to `order-updates`

---

## 🛠️ Implementation Steps

1. **User Registration & Login**

   * Register via `/users`
   * Login via `/users/login` → store the returned JWT token

2. **Create Order**

   * Use `POST /orders` with JWT token
   * Internally, this publishes to `payment-events` topic

3. **Payment Processing**

   * Payment service listens for `payment-events`
   * After processing, it sends a result to `order-updates`

4. **Order Status Update**

   * Order service listens for updates on `order-updates`
   * Updates order status in database (e.g., `PROCESSING → PAID`)

---

## 🔧 Tech Stack

* **Language**: Golang
* **Framework**: [Iris](https://www.iris-go.com/)
* **Database**: Cloud SQL (PostgreSQL/MySQL)
* **Messaging**: Google Cloud Pub/Sub
* **Auth**: JWT (Go middleware)
* **ORM**: GORM

---

## 🚀 Deployment

* The microservices are deployed on Google Cloud Platform (GCP) using Cloud Run:
    *	🔗 Order Service: https://order-microservice-747901258630.us-central1.run.app
    *	🔗 Payment Service: https://payment-microservice-747901258630.us-central1.run.app
