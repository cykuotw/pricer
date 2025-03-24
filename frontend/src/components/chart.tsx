import { useEffect, useState } from "react";

import { API_URL } from "../configs/config";
import { useTicker } from "../contexts/SelectedTickerContext";

import { Line } from "react-chartjs-2";
import {
    Chart as ChartJS,
    LineElement,
    PointElement,
    CategoryScale,
    LinearScale,
    Title,
    Tooltip,
    Legend,
    ChartData,
    ChartOptions,
} from "chart.js";

ChartJS.register(
    LineElement,
    PointElement,
    CategoryScale,
    LinearScale,
    Title,
    Tooltip,
    Legend
);

interface Prices {
    ticker: string;
    prices: string[];
    times: string[];
}

interface LatestPrice {
    price: string;
    time: string;
    ticker: string;
}

interface MarketOpen {
    isopen: boolean;
}

/**
 * Chart component displays a real-time line chart of stock prices for a selected ticker.
 * It fetches historical prices and updates the chart with live price data using Server-Sent Events (SSE).
 *
 * Features:
 * - Displays historical price data for the selected ticker.
 * - Updates the chart in real-time when the market is open.
 * - Shows whether the market is open or closed.
 *
 * Dependencies:
 * - React hooks: `useState`, `useEffect`
 * - Chart.js for rendering the line chart
 * - Context: `useTicker` for accessing the selected ticker
 *
 * @component
 * @returns {JSX.Element} A line chart with stock prices and a header showing the market status or selected ticker.
 */
export default function Chart() {
    const { selectedTicker } = useTicker();

    const [isMarketOpen, setIsMarketOpen] = useState<boolean>(true);
    const [prices, setPrices] = useState<number[]>([]);
    const [times, setTimes] = useState<string[]>([]);

    useEffect(() => {
        // Fetch whether the market is open
        const checkMarketOpen = async () => {
            const response = await fetch(`${API_URL}/check-open`, {
                method: "GET",
            });

            const marketOpen: MarketOpen = await response.json();
            setIsMarketOpen(marketOpen.isopen);
        };

        // Fetch historical prices for the selected ticker
        const fetchHistorPrices = async () => {
            if (selectedTicker.length === 0) return;

            const response = await fetch(
                `${API_URL}/prices/${selectedTicker}`,
                {
                    method: "GET",
                }
            );

            const status = await response.status;
            if (status === 403) {
                const message = await response.json();
                console.log(message);
            }

            const prices: Prices = await response.json();
            setPrices(prices.prices.map((p) => parseFloat(p)));
            setTimes(
                prices.times.map((t) => {
                    return new Date(t).toLocaleTimeString("en-CA", {
                        hour: "2-digit",
                        minute: "2-digit",
                        hour12: false,
                    });
                })
            );
        };

        checkMarketOpen();
        fetchHistorPrices();
    }, [selectedTicker]);

    useEffect(() => {
        // Subscribe to live price updates via SSE when the market is open
        if (selectedTicker.length === 0 || !isMarketOpen) return;

        const es = new EventSource(
            `${API_URL}/stream/update-price/${selectedTicker}`
        );

        es.onmessage = (e) => {
            const data: LatestPrice = JSON.parse(e.data);

            if (data.ticker != selectedTicker) return;

            setPrices((prev) => [...prev, parseFloat(data.price)]);
            setTimes((prev) => [
                ...prev,
                new Date(data.time).toLocaleTimeString("en-CA", {
                    hour: "2-digit",
                    minute: "2-digit",
                    hour12: false,
                }),
            ]);
        };

        es.onerror = (err) => {
            console.error("EventSource failed:", err);
            es.close();
        };

        return () => {
            es.close();
        };
    }, [selectedTicker]);

    const data: ChartData<"line"> = {
        labels: times,
        datasets: [
            {
                label: "Price",
                data: prices,
                fill: false,
                borderColor: "rgb(75, 192, 192)",
                tension: 0.1,
            },
        ],
    };

    const options: ChartOptions<"line"> = {
        responsive: true,
        plugins: {
            legend: { display: true },
        },
        scales: {
            x: {
                title: { display: true, text: "Time" },
            },
            y: {
                title: { display: true, text: "Price (USD)" },
            },
        },
    };

    return (
        <>
            <h1>
                {!isMarketOpen
                    ? "CLOSED"
                    : selectedTicker.length !== 0
                    ? selectedTicker
                    : "---"}
            </h1>
            <Line data={data} options={options} />
        </>
    );
}
