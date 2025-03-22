import { createContext, useContext, useState, ReactNode } from "react";

type TickerContextType = {
    selectedTicker: string;
    setSelectedTicker: (ticker: string) => void;
};

const TickerContext = createContext<TickerContextType | undefined>(undefined);

export const TickerProvider = ({ children }: { children: ReactNode }) => {
    const [selectedTicker, setSelectedTicker] = useState("");

    return (
        <TickerContext.Provider value={{ selectedTicker, setSelectedTicker }}>
            {children}
        </TickerContext.Provider>
    );
};

export const useTicker = () => {
    const context = useContext(TickerContext);
    if (!context)
        throw new Error("useTicker must be used within TickerProvider");
    return context;
};
