# Pricing Service (/services/pricing)

This package contains two major responsibilities:

-   **Routing**: Handles HTTP and SSE requests, directing them to the appropriate business logic layer.
-   **Controller**: Contains the core logic for simulating and updating ticker prices.

# Routing Package (/services/pricing/router)

## Responsibilities

-   Handle incoming HTTP and Server-Sent Events (SSE) requests.
-   Route requests to the appropriate controller logic.
-   Manage API endpoints for price updates and ticker metadata.

---

## Key Files

| File                  | Description                                                                                  |
| --------------------- | -------------------------------------------------------------------------------------------- |
| `router.go`           | Main entry point for routing logic.                                                          |
| `handleCheckOpen.go`  | Handles requests to check if the market is currently open.                                   |
| `handleConfig.go`     | Manages API endpoints for retrieving and updating ticker configuration.                      |
| `handlePrices.go`     | Handles price-related endpoints, including fetching historical prices and streaming updates. |
| `handleServerTime.go` | Streams the current server time to clients using Server-Sent Events (SSE).                   |
| `handleTickerList.go` | Provides a list of all available tickers.                                                    |

---

## API Endpoints

The following endpoints are available in the **Routing Package**:

| Endpoint                       | Method | Description                                                                |
| ------------------------------ | ------ | -------------------------------------------------------------------------- |
| `/tickers`                     | GET    | Retrieves a list of all available tickers.                                 |
| `/config/:ticker`              | GET    | Fetches the configuration (drift and volatility) for a specific ticker.    |
| `/config`                      | PUT    | Updates the configuration (drift and volatility) for a specific ticker.    |
| `/check-open`                  | GET    | Checks if the market is currently open.                                    |
| `/prices/:ticker`              | GET    | Retrieves historical prices for a specific ticker.                         |
| `/stream/server-time`          | GET    | Streams the current server time to clients using Server-Sent Events.       |
| `/stream/update-price/:ticker` | GET    | Streams live price updates for a specific ticker using Server-Sent Events. |

---

# Controller Package (/services/pricing/controller)

## Responsibilities

-   Initialized per-ticker using Go routines (concurrency)
-   Perform GBM-based price simulation
-   Manage ticker metadata, update frequency, and control flow
-   Parse and manage initial config from JSON

---

## Key Files

| File              | Description                                                 |
| ----------------- | ----------------------------------------------------------- |
| `controller.go`   | Entry point for ticker controller logic                     |
| `checkUpdated.go` | Verifies if prices need update and triggers re-computation  |
| `config.go`       | Loads per-ticker drift/volatility config                    |
| `const.go`        | Stores all constants                                        |
| `data.go`         | Loads initial parameters from JSON file                     |
| `simulate.go`     | Implements Geometric Brownian Motion (GBM) simulation logic |
| `ticker.go`       | Maps tickers to data/config and handles selection           |
| `time.go`         | Utility time-specific functions                             |
| `updatePrice.go`  | Handles per-minute price updates                            |

---

## Price Simulation Overview

-   GBM formula: S(t+1) = S(t) _ exp((μ - 0.5 _ σ²) _ dt + σ _ √dt \* Z)

    -   `μ` = drift,
    -   `σ` = volatility,
    -   `dt = 1 / (365 * 24 * 60)`

-   Uses `decimal.Decimal` to avoid floating point (IEEE-754) error
-   Prices updated once per minute

---

## Concurrency Notes

-   Goroutines launched for when service is boosted
-   Uses `sync.WaitGroup` and `sync.Mutex` to ensure task completion and prevent race conditions.

---

## External Config

-   `data/initParam.json` used to load:
    -   Ticker names
    -   Initial price
    -   Drift / Volatility values
