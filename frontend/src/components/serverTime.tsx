import { useEffect, useState } from "react";
import { API_URL } from "../configs/config";

export default function ServerTime() {
    const [serverTime, setServerTime] = useState<string>("connecting...");

    useEffect(() => {
        const es = new EventSource(`${API_URL}/stream/server-time`);

        es.onopen = () => {
            console.info("event source open successfully");
        };

        es.onmessage = (e) => {
            const date = new Date(e.data);
            const formattedTime = `${date.toLocaleDateString(
                "en-CA"
            )} ${date.toLocaleTimeString("en-CA", {
                hour12: false,
                timeZoneName: "short",
            })}`;

            setServerTime(formattedTime);
        };

        es.onerror = (err) => {
            console.error("EventSource failed:", err);
            setServerTime("disconnected, please refresh to reconnect.");
            es.close();
        };

        return () => {
            es.close();
        };
    }, []);

    return (
        <div className="flex flex-col items-center p-4">
            <h2 className="text-xl mb-2">Server Time</h2>
            <p className="list-disc pl-5">{serverTime}</p>
        </div>
    );
}
