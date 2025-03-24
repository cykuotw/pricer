import { useEffect, useState } from "react";
import { API_URL } from "../configs/config";

/**
 * ServerTime component displays the current server time in real-time.
 * It uses Server-Sent Events (SSE) to receive live updates from the server.
 *
 * Features:
 * - Connects to the server's `/stream/server-time` endpoint to fetch the current time.
 * - Updates the displayed time in real-time.
 * - Handles connection errors gracefully and displays a message when disconnected.
 *
 * Dependencies:
 * - React hooks: `useState`, `useEffect`
 *
 * @component
 * @returns {JSX.Element} A styled component showing the server's current time.
 */
export default function ServerTime() {
    const [serverTime, setServerTime] = useState<string>("connecting...");

    useEffect(() => {
        // Establish a Server-Sent Events (SSE) connection to the server
        const es = new EventSource(`${API_URL}/stream/server-time`);

        es.onopen = () => {
            console.info("event source open successfully");
        };

        // Update the server time when a message is received
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
