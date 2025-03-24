# Pricing App

A real-time financial pricing simulation platform built with Go and React.
Supports live price updates via Server-Sent Events (SSE).

---

## Tech Stack

### 🔹 Frontend

-   React (TypeScript)
-   Vite
-   TailwindCSS v4 + DaisyUI v5
-   Chart.js via `react-chartjs-2`

### 🔹 Backend

-   Go 1.23
-   Gin
    -   RESTful API
    -   Server-Sent Events (SSE)

---

## Project Structure

```
.
├── api/                     # API entry point
├── config/                  # Configuration structs (env, market params)
├── data/                    # Static config (e.g. GBM parameters)
├── docker/                  # Dockerfiles for backend & frontend
├── frontend/                # React frontend app (Vite + TailwindCSS)
│   ├── public/              # Static assets (e.g., favicon)
│   │
│   ├── src/                 # Source code
│   │   ├── components/      # Reusable React components (e.g. Chart, TickerSelector)
│   │   ├── configs/         # App-level config (e.g. API base URL)
│   │   ├── contexts/        # React Contexts (e.g. SelectedTickerContext)
│   │   ├── styles/          # Tailwind and custom stylesheets
│   │   ├── App.tsx          # Main app shell
│   │   ├── main.tsx         # Entry point for React
│   │
│   ├── .env                 # Frontend environment variables
│   ├── index.html           # HTML entry point for Vite
│
├── services/                # Core domain and pricing logic
│   ├── middleware/          # Custom middleware (e.g. CORS)
│   └── pricing/             # Pricing module (controller + routing)
│       ├── controller/      # Price simulation logic
│       └── router/          # API routes for tickers, server time, etc.
│
├── .air.toml                # Live reload config for Go
├── .env                     # Backend environment variables
├── main.go                  # Main Go entry point
├── go.mod                   # Go module config
├── go.sum                   # Go dependency checksums
├── docker-compose.yml       # Dev orchestration (frontend + backend)
├── Makefile                 # Developer scripts and build automation

```

---

## Running the Project

### With Docker Compose

```bash
docker-compose up --build
```

> **Warning:** Total image size can be up to **2.5 GB**.

### Dev Without Docker

#### Start backend:

Please have Go 1.23 installed. (Go Official: https://go.dev/doc/install)

```bash
go install github.com/cosmtrek/air@latest  # if you do not have air installed
air
```

#### Start frontend:

```bash
cd frontend
npm install
npm run dev
```

---

## Price Simulation

-   Uses Geometric Brownian Motion (GBM)
-   Parameters from `data/initParam.json`
-   Tickers updated per minute using goroutines
-   Exposed via `/stream/update-price/` (SSE)

---

## Assumptions

-   User count is fewer than 10, so there's no need to scale for a high number of concurrent users.
-   Market is assumed to be open daily from 9:30 AM to 4:00 PM, ignoring holidays and weekends.
-   Users will only access the service during market open hours and will reopen it the next day.
-   Users are experienced traders who are familiar with pricing models such as GBM.
-   Prices are not persisted or stored.

---

## Possible Future Improvements

-   **Improve time synchronization** (e.g., use an NTP server) along with higher-precision data types (e.g., `decimal.Decimal`).
-   **Increase precision in price simulation**:
    -   It's unclear how much floating-point error may affect the accuracy of simulated prices.
    -   If the error proves significant, switching to a decimal-based type should be considered.
-   **Consider a lighter backend framework**:
    -   Gin is a solid choice for building Go backends, but there are lighter and faster alternatives available depending on future scale.
-   **Consider using a DBMS** (e.g., PostgreSQL, Redis) to store prices:
    -   In real-world scenarios, thousands of stock tickers may need to be monitored, and an in-memory buffer could become a bottleneck.
-   **Transition to an event-driven architecture**:
    -   In real-world scenarios, thousands of stock tickers may need to be monitored, which could exceed the limits of a monolithic system.
-   **Enhance observability in a distributed system** (e.g., via OpenTelemetry):
    -   Currently, the system only prints error messages; proper error handling and structured logging are needed for production readiness.
    -   Richer telemetry data enables engineers to make informed decisions, especially during on-call incidents.
