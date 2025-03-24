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

export default function Chart() {
    const { selectedTicker } = useTicker();

    const [isMarketOpen, setIsMarketOpen] = useState<boolean>(true);
    const [prices, setPrices] = useState<number[]>([]);
    const [times, setTimes] = useState<string[]>([]);

    useEffect(() => {
        const checkMarketOpen = async () => {
            const response = await fetch(`${API_URL}/check-open`, {
                method: "GET",
            });

            const marketOpen: MarketOpen = await response.json();
            setIsMarketOpen(marketOpen.isopen);
        };

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
