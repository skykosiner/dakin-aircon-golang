import axios, { AxiosResponse } from "axios";

type Status = {
    Temp:     string
    Mode:     string
    FanSpeed: string
    Power:    string
    F_dir:    string
    Shum:     string
};

const modeMap = new Map<string, string>([
    ["3", "Cold"],
    ["4", "Hot"],
]);

const powerMap = new Map<string, string>([
    ["1", "On"],
    ["0", "Off"],
]);

const fanMap = new Map<string, string>([
    ["B", "Night"],
    ["3", "1"],
    ["4", "2"],
    ["5", "3"],
    ["6", "4"],
    ["7", "5"],
]);

const apiValues = new Map<string, string>([
        // Modes
		["Cold", "3"],
		["Heat", "4"],
		// Power
		["On", "1"],
		["Off", "0"],
		// Fan speed
		["Night", "B"],
		["1", "3"],
		["2", "4"],
		["3", "5"],
		["4", "6"],
		["5", "7"],
]);

async function sendRequest(ip: string, power: string, mode: string, temp: string, fanRate: string): Promise<void> {
    axios.get(`http://${ip}/aircon/set_control_info?pow=${power}&mode=${mode}&stemp=${temp}&shum=0&f_rate=${fanRate}&f_dir=3`)
    .catch((err) => {
        console.error("There was an error sending the request", err);
    });
}

async function getState(ip: string): Promise<AxiosResponse | void> {
    return axios.get(`http://${ip}/aircon/get_control_info`).catch((err) => console.error(err));
}

export async function getCurrentStatus(ip: string): Promise<Status> {
    let status: Status = {Temp: "0", Mode: "0", Shum: "0", F_dir: "0", Power: "0", FanSpeed: "0"};
        const resp = await getState(ip);
        let respArr: string[];
        //@ts-ignore
        const data = resp.data as string;
        respArr = data.split(",")

        for (const item of respArr) {
            const parts = item.split("=")

            switch (parts[0]) {
                case "stemp":
                    status.Temp = parts[1];
                    break;
                case "mode":
                    status.Mode = modeMap.get(parts[1]) as string;
                    break;
                case "pow":
                    status.Power = powerMap.get(parts[1]) as string;
                    break;
                case "f_rate":
                    status.FanSpeed = fanMap.get(parts[1]) as string;
                    break;
                case "f_dir":
                    status.F_dir = parts[1]
                    break;
                case "shum":
                    status.Shum = parts[1]
                    break;
            };
        };

    return status;
};

export async function toggleAircon(ip: string, state: boolean): Promise<void> {
    const status = getCurrentStatus(ip);
    const currState = await status;
    const onOrOff = new Map<boolean, string>([
        [false, "0"],
        [true, "1"]
    ]);

    const stateFinal = onOrOff.get(state) as string;

    sendRequest(ip, stateFinal, apiValues.get(currState.Mode) as string, currState.Temp, apiValues.get(currState.FanSpeed) as string);
}

export async function setHotOrCool(ip: string, cool: boolean): Promise<void> {
    const status = getCurrentStatus(ip);
    const currState = await status;
    const hotOrCool = new Map<boolean, string>([
        [true, "3"],
        [false, "4"],
    ]);

    const hotOrCoolFinal = hotOrCool.get(cool) as string;
    sendRequest(ip, apiValues.get(currState.Power) as string, hotOrCoolFinal, currState.Temp, apiValues.get(currState.FanSpeed) as string);
};

export async function setFanRate(ip: string, rate: string): Promise<void> {
    rate = rate.split("-")[1]
    const status = getCurrentStatus(ip);
    const currState = await status;

    sendRequest(ip, apiValues.get(currState.Power) as string, apiValues.get(currState.Mode) as string, currState.Temp, apiValues.get(rate) as string);
};

export async function setTemp(ip: string, temp: string): Promise<void> {
    const status = getCurrentStatus(ip);
    const currState = await status;

    sendRequest(ip, apiValues.get(currState.Power) as string, apiValues.get(currState.Mode) as string, temp, apiValues.get(currState.FanSpeed) as string);
};
