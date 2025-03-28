import { useEffect, useState } from "react";
import { API_URL } from "../configs/config";
import { useTicker } from "../contexts/SelectedTickerContext";

interface ConfigData {
    open: number;
    drift: string; // in 0.0001
    volatility: string; // in 0.01
}

interface ConfigPayload {
    ticker: string;
    config: ConfigData;
}

interface UpdateConfigData {
    ticker: string;
    drift: string;
    volatility: string;
}

/**
 * TickerSelector component allows users to select a stock ticker and configure its drift and volatility.
 * It fetches the list of available tickers and their configurations from the server.
 * Users can update the drift and volatility values, which are sent back to the server.
 *
 * Features:
 * - Fetches and displays a list of available tickers.
 * - Fetches and displays the configuration (drift and volatility) for the selected ticker.
 * - Allows users to update the configuration and sends the updates to the server.
 *
 * Dependencies:
 * - React hooks: `useState`, `useEffect`
 * - Context: `useTicker` for managing the selected ticker.
 *
 * @component
 * @returns {JSX.Element} A form for selecting a ticker and configuring its drift and volatility.
 */
export default function TickerSelector() {
    const { selectedTicker, setSelectedTicker } = useTicker();

    const [tickerList, setTickerList] = useState<string[]>([]);
    const [drift, setDrift] = useState<string>("0");
    const [volatility, setVlatility] = useState<string>("0");

    const DRIFT_MULTIPLE = 10000;
    const VOLATILITY_MULTIPLE = 100;

    useEffect(() => {
        // Fetch the list of available tickers from the server
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

        // Fetch the configuration (drift and volatility) for the selected ticker
        const fetchConfig = async (ticker: string) => {
            const response = await fetch(`${API_URL}/config/${ticker}`, {
                method: "GET",
            });
            const data: ConfigPayload = await response.json();

            setDrift(data.config.drift);
            setVlatility(data.config.volatility);
        };

        fetchConfig(selectedTicker);
    }, [selectedTicker]);

    useEffect(() => {
        if (
            drift.length === 0 ||
            volatility.length === 0 ||
            selectedTicker.length === 0
        )
            return;

        // Send updated configuration (drift and volatility) to the server
        const fetchUpdateConfig = async (payload: UpdateConfigData) => {
            const response = await fetch(`${API_URL}/config`, {
                method: "PUT",
                body: JSON.stringify(payload),
            });
            const result = await response.status;

            if (result !== 201) {
                const msg = await response.json();
                console.warn("udpate fail:", msg.message);
            }
        };

        const data: UpdateConfigData = {
            ticker: selectedTicker,
            drift: drift,
            volatility: volatility,
        };
        fetchUpdateConfig(data);
    }, [drift, volatility]);

    return (
        <>
            <div className="flex flex-col items-center space-y-2 max-w-1/2">
                {/* Dropdown for selecting a ticker */}
                <select
                    defaultValue="Pick a Ticker"
                    className="select select-neutral text-center w-max"
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

                {/* Input fields for configuring drift and volatility */}
                <div className="flex flex-col space-y-0.5 max-w-full">
                    <label className="input">
                        <span>Drift (μ, mu):</span>
                        <input
                            type="number"
                            className="input input-neutral input-ghost text-right"
                            value={drift}
                            onChange={(e) => setDrift(e.target.value)}
                        />
                        <span>
                            x 0.0001 = {parseFloat(drift) / DRIFT_MULTIPLE}
                        </span>
                    </label>
                    <label className="input">
                        <span>Volatility (σ, sigma):</span>
                        <input
                            type="number"
                            className="input input-neutral input-ghost text-right"
                            value={volatility}
                            onChange={(e) => setVlatility(e.target.value)}
                        />
                        <span>
                            x 0.01 ={" "}
                            {parseFloat(volatility) / VOLATILITY_MULTIPLE}
                        </span>
                    </label>
                </div>
            </div>
        </>
    );
}
