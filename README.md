# voucher-redeem

## Description

A simple voucher redemption API built with Go (Golang), following clean architecture principles. The service allows users to redeem vouchers using unique codes and provides proper validation and handling for different scenarios.

## Features

- Redeem voucher using a code
- Validate voucher expiration and availability
- Clean architecture (handlers, services, repositories)

## Getting Started

### Prerequisites

- Golang
- MySQL

### Installation

1. Clone the repository:

```bash
git clone https://github.com/fauzan264/voucher-redeem.git
cd voucher-redeem
```

2. Copy the example environment file and adjust it to your local settings:

```bash
cp .env.example .env
```

3. Run the project:

```bash
go run main.go
```

## Folder Structure

- `config/` – Application configuration
- `constants/` – Shared constants
- `docs/` – Documentation and Postman collection
- `domain/` – Business domain entities and interfaces
- `repositories/` – Database access logic
- `services/` – Core application/Business logic
- `handlers/` – HTTP handlers
- `helpers/` – Utility helpers
- `middleware/` – Fiber middleware
- `docs/` – API documentation and Postman collections
- `utils/` – Generic utility functions

## Database Diagram

You can view the database entity-relationship diagram (ERD) by checking the link stored in the following file:  
[`docs/dbdiagram-links.md`](docs/dbdiagram-links.md)

## Postman Collection

To test the `redeem-voucher` endpoint, import the following file into Postman:  
[`docs/Redeem Voucher.postman_collection.json`](docs/Redeem%20Voucher.postman_collection.json)