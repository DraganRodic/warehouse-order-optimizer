# Warehouse Order Optimizer

Warehouse Order Optimizer is a Go backend application designed to automate warehouse picking operations.

The application allows warehouse staff to:

* Import inventory data from Excel files
* Store product locations in MySQL
* Upload purchase orders from Excel
* Automatically lookup product locations
* Sort products by warehouse location
* Manage products and locations through REST APIs

This eliminates the need for manual VLOOKUP operations in Excel and significantly improves warehouse picking efficiency.

---

## Features

### Inventory Import

Import inventory data from Excel files containing:

| SKU     | Location |
| ------- | -------- |
| WKD-001 | 7R       |
| WKD-002 | 12L      |

The application parses the Excel file and stores all products in MySQL.

---

### Purchase Order Processing

Upload a purchase order Excel file containing product codes.

Example:

| SKU     |
| ------- |
| WKD-001 |
| WKD-002 |
| WKD-003 |

The application will:

1. Read all product codes
2. Lookup product locations
3. Sort products by warehouse location
4. Return an optimized picking list

Example output:

| SKU     | Location |
| ------- | -------- |
| WKD-001 | 1L       |
| WKD-005 | 1R       |
| WKD-002 | 2L       |
| WKD-003 | 2R       |

---

### Product Management

Manage warehouse products using REST APIs.

Supported operations:

* Create Product
* Get Product
* Get All Products
* Update Product
* Delete Product

This allows warehouse staff to:

* Add new SKUs
* Change product locations
* Remove discontinued products

---

## Tech Stack

* Go
* Gin
* GORM
* MySQL
* Excelize
* REST API

---

## Project Structure

```text
cmd/
└── server/

internal/
├── config/
├── database/
├── handler/
├── model/
├── repository/
├── router/
└── service/

uploads/

migrations/
```

---

## API Endpoints

### Inventory

#### Upload Inventory File

```http
POST /api/inventory/upload
```

#### Import Inventory

```http
POST /api/inventory/import
```

---

### Orders

#### Process Purchase Order

```http
POST /api/orders/process
```

---

### Products

#### Get All Products

```http
GET /api/products
```

#### Get Product By ID

```http
GET /api/products/:id
```

#### Create Product

```http
POST /api/products
```

#### Update Product

```http
PUT /api/products/:id
```

#### Delete Product

```http
DELETE /api/products/:id
```

---

## Installation

Clone the repository:

```bash
git clone https://github.com/DraganRodic/warehouse-order-optimizer.git
```

Move into the project directory:

```bash
cd warehouse-order-optimizer
```

Install dependencies:

```bash
go mod tidy
```

---

## Environment Variables

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=warehouse_optimizer
SERVER_PORT=8080
```

---

## Run Application

Start MySQL.

Run database migrations.

Start the server:

```bash
go run cmd/server/main.go
```

The API will be available at:

```text
http://localhost:8080
```

---

## Example Workflow

### Step 1

Import warehouse inventory:

```http
POST /api/inventory/import
```

Upload inventory Excel file.

---

### Step 2

Process purchase order:

```http
POST /api/orders/process
```

Upload purchase order Excel file.

---

### Step 3

Receive sorted picking list.

Warehouse staff can now pick products in optimized shelf order.

---

## Future Improvements

* Web dashboard
* Authentication & authorization
* Excel export of picking lists
* Warehouse analytics
* Inventory reporting
* Docker support

---

## Author

Dragan Rodić

Backend project built with Go, Gin, GORM, MySQL and Excel processing.
