import Chart from "./components/chart";
import ServerTime from "./components/serverTime";
import TickerSelector from "./components/tickerSelector";

function App() {
    return (
        <div className="flex flex-col items-center py-3">
            <ServerTime></ServerTime>
            <TickerSelector></TickerSelector>
            <Chart></Chart>
        </div>
    );
}

export default App;
