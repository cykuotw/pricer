import { createContext, useContext, useState, ReactNode } from "react";

/**
 * Context for managing the selected stock ticker across the application.
 * Provides the selected ticker and a function to update it.
 *
 * Features:
 * - Stores the currently selected ticker in a React context.
 * - Allows components to access and update the selected ticker.
 * - Ensures the context is used within the `TickerProvider`.
 *
 * Dependencies:
 * - React hooks: `useState`, `useContext`
 *
 * @module SelectedTickerContext
 */

/**
 * Type definition for the TickerContext.
 * @typedef {Object} TickerContextType
 * @property {string} selectedTicker - The currently selected ticker.
 * @property {function(string): void} setSelectedTicker - Function to update the selected ticker.
 */
type TickerContextType = {
    selectedTicker: string;
    setSelectedTicker: (ticker: string) => void;
};

const TickerContext = createContext<TickerContextType | undefined>(undefined);

/**
 * Provider component for the TickerContext.
 * Wraps the application or part of it to provide access to the selected ticker.
 *
 * @param {Object} props - Component props.
 * @param {ReactNode} props.children - Child components that will have access to the context.
 * @returns {JSX.Element} The provider component.
 */
export const TickerProvider = ({ children }: { children: ReactNode }) => {
    const [selectedTicker, setSelectedTicker] = useState("");

    return (
        <TickerContext.Provider value={{ selectedTicker, setSelectedTicker }}>
            {children}
        </TickerContext.Provider>
    );
};

/**
 * Custom hook to access the TickerContext.
 * Ensures the context is used within a `TickerProvider`.
 *
 * @throws {Error} If used outside of a `TickerProvider`.
 * @returns {TickerContextType} The context value.
 */
export const useTicker = () => {
    const context = useContext(TickerContext);
    if (!context)
        throw new Error("useTicker must be used within TickerProvider");
    return context;
};
