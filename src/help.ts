import fs from "fs";

function doesFileExist(): boolean {
    const path = `${process.env.HOME}/.local/airconhelp.txt`;

    if (!fs.existsSync(path)) {
        return false
    }

    return true
};

export function setupHelp(): void {
    if (!doesFileExist()) {
        const src = "./helptext.txt"
        const dest = `${process.env.HOME}/.local/airconhelp.txt`

        fs.copyFile(src, dest, (err) => {
            if (err) {
                console.log("COCK IN MY ASS", err);
            }
        });
    };
}
