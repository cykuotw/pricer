import { useEffect, useState } from "react";
import { API_URL } from "../configs/config";

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

export default function TickerSelector() {
    const [tickerList, setTickerList] = useState<string[]>([]);
    const [selectedTicker, setSelectedTicker] = useState<string>("");
    const [drift, setDrift] = useState<string>("0");
    const [volatility, setVlatility] = useState<string>("0");

    const DRIFT_MULTIPLE = 10000;
    const VOLATILITY_MULTIPLE = 100;

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
        if (drift.length === 0 || volatility.length === 0) return;

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
                <div className="flex flex-col space-y-0.5 max-w-4/5">
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
