# PNLyzer

**PNLyzer** is a powerful tool designed for analyzing spot trading history. It provides a platform for users to manually record spot trading transactions, enabling them to keep track of trades, analyze performance, and generate reports.

## Features

- **User-Friendly Spot Trading History Management:**
  Easily record spot trading transactions, including details such as stock symbol, quantity, price, and timestamp.

- **Performance Analysis:**
  Gain insights into your trading performance with Profit and Loss (P&L) reports. Understand how individual stocks and your overall portfolio are performing.

- **Versioned Application Information:**
  The application exposes a simple JSON endpoint (GET /) that returns information about the application name and version, extracted from a VERSION file in the project root.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (at least version 1.14)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/pnlyzer.git
    ```
2. Change into the project directory:
    ```bash
    cd pnlyzer
    ```
3. Build the project:
    ```bash
    go build
    ```
### Usage
To start the PNLyzer application, run:
```bash
./pnlyzer
```
Visit the application at http://localhost:8080 to access the version information.

