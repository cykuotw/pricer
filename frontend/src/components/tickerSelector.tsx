import { useEffect, useState } from "react";
import { API_URL } from "../configs/config";

export default function TickerSelector() {
    const [tickerList, setTickerList] = useState<string[]>([]);
    const [selectedTicker, setSelectedTicker] = useState("");

    useEffect(() => {
        const fetchTickerList = async () => {
            const response = await fetch(`${API_URL}/tickers`, {
                method: "GET",
            });
            const data = await response.json();

            setTickerList(data.tickers);
        };

        fetchTickerList();
    }, []);

    useEffect(() => {
        if (selectedTicker.length === 0) return;

        console.log(selectedTicker);
    }, [selectedTicker]);

    return (
        <>
            <select
                defaultValue="Pick a Ticker"
                className="select select-neutral text-center"
                onChange={(e) => {
                    setSelectedTicker(e.target.value);
                }}
            >
                <option disabled={true}>Pick a Ticker</option>
                {tickerList.map((ticker) => {
                    return (
                        <option key={ticker} value={ticker}>
                            {ticker}
                        </option>
                    );
                })}
            </select>
        </>
    );
}
