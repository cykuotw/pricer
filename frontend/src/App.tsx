import Chart from "./components/chart";
import ServerTime from "./components/serverTime";
import TickerSelector from "./components/tickerSelector";
import { TickerProvider } from "./contexts/SelectedTickerContext";

function App() {
    return (
        <div className="flex flex-col items-center py-3">
            <ServerTime></ServerTime>

            <TickerProvider>
                <TickerSelector></TickerSelector>
                <Chart></Chart>
            </TickerProvider>
        </div>
    );
}

export default App;
