import { useEffect, useState } from "react";
import { API_URL } from "../configs/config";
import { useTicker } from "../contexts/SelectedTickerContext";

interface Prices {
    ticker: string;
    t0: string;
    prices: string[];
}

export default function Chart() {
    const { selectedTicker } = useTicker();

    const [prices, setPrices] = useState<string[]>([]);

    useEffect(() => {
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
            setPrices(prices.prices);
        };

        fetchHistorPrices();
    }, [selectedTicker]);

    // TODO: Price Chart
    return (
        <>
            <h1>{selectedTicker}</h1>
            <ul>
                {prices.map((price, index) => {
                    return <li key={index}>{price}</li>;
                })}
            </ul>
        </>
    );
}
