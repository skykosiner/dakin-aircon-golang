import fs from "fs";
import { EventEmitter } from "stream";

export type Config = {
    MainIp: string,
    ConflictAirconOne: string,
    ConflictAirconTwo: string,
};

class config extends EventEmitter {
    private conf: Config = {MainIp: "COCK", ConflictAirconOne: "PUSSY", ConflictAirconTwo: "CUM"};

    public readConfig(): void {
        const path = `${process.env.HOME}/.config/aircon/aircon.json`;
        fs.readFile(path, (err, buffer) => {
            if (err) {
                console.log("CUM IN MY ASS", err);
            };

            this.conf = JSON.parse(buffer.toString());
            this.emit("config", this.conf);
        });
    };
}

export function GetConfig(): Config {
    const Config = new config();
    Config.readConfig();
    let conf = Config.on("config", (config: Config): Config => {
        return config
    });

    return conf;
};
