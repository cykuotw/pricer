import Chart from "./components/chart";
import ServerTime from "./components/serverTime";
import TickerSelector from "./components/tickerSelector";
import { TickerProvider } from "./contexts/SelectedTickerContext";

function App() {
    return (
        <div className="flex flex-col items-center mx-auto my-3 space-y-3 lg:max-w-1/2 md:max-w-2/3">
            <ServerTime></ServerTime>

            <TickerProvider>
                <TickerSelector />
                <Chart />
            </TickerProvider>
        </div>
    );
}

export default App;
