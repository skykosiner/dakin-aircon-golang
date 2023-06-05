import { spawn } from "child_process";
import { getCurrentStatus, setFanRate, setHotOrCool, setTemp, toggleAircon } from "./aircon";

async function main() {
    const state = getCurrentStatus("10.0.0.10");
    const currState = await state;
    // Take in input from stdin
    const args = process.argv.slice(2);

    switch (args[0]) {
        case "toggle":
            if (currState.Power == "On") {
                toggleAircon("10.0.0.10", false);
            } else {
                toggleAircon("10.0.0.10", true);
            };
            break;
        case "cold":
            setHotOrCool("10.0.0.10", true);
            break;
        case "hot":
            setHotOrCool("10.0.0.10", false);
            break;
        case "status":
            // If not connected to the correct wifi don't show the status
            spawn("iwgetid", ["-r"]).stdout.on("data", (chunk: Buffer) => {
                const networkName = chunk.toString();

                if (networkName.trim() !== "The Kosiner's wifi") {
                    console.log("Incorrect network name");
                } else {
                    const statusString = `{${currState.Temp} ${currState.Mode} ${currState.FanSpeed} ${currState.Power}}`;
                    console.log(statusString);
                };
            });
            break;
        default:
            if (args[0].includes("fan")) {
                setFanRate("10.0.0.10", args[0]);
            } else {
                setTemp("10.0.0.10", args[0]);
            };
    };
};

main();