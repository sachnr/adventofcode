import * as fs from "fs";

type Entry = {
    password: string;
    letter: string;
    min: number;
    max: number;
};

class DayTwo {
    entries: Entry[];

    constructor() {
        this.entries = [];
    }

    parse(input: string): void | Error {
        const lines = input.split("\n");
        for (const line of lines) {
            if (line === "") {
                continue;
            }
            const parts = line.split(" ");
            const min_max = parts[0].split("-");
            const min = parseInt(min_max[0]);
            const max = parseInt(min_max[1]);
            const letter = parts[1].replace(":", "");
            const password = parts[2];

            const data: Entry = {
                min,
                max,
                password,
                letter,
            };
            this.entries.push(data);
        }
    }

    partOne(): number {
        let valid = 0;
        for (const entry of this.entries) {
            const len = entry.password.length;
            let count = 0;
            for (let i = 0; i < len; i++) {
                if (entry.password.charAt(i) === entry.letter) {
                    count += 1;
                }
            }
            if (count <= entry.max && count >= entry.min) {
                valid += 1;
            }
        }
        return valid;
    }

    partTwo(): number {
        let valid = 0;
        for (const entry of this.entries) {
            const firstPositionMatch =
                entry.password.charAt(entry.min - 1) === entry.letter;
            const secondPositionMatch =
                entry.password.charAt(entry.max - 1) === entry.letter;

            if (firstPositionMatch !== secondPositionMatch) {
                valid += 1;
            }
        }
        return valid;
    }
}

function DayTwoRun(): string {
    const daytwo = new DayTwo();
    let output: String[] = [];
    try {
        const file = fs.readFileSync("src/inputs/day2.txt", "utf8");
        daytwo.parse(file);
    } catch (error: any) {
        console.error(`failed to parse the file: ${error.message}`);
    }
    output.push("Day Two: ");

    const partOne = daytwo.partOne();
    if (partOne !== null) {
        output.push(`\n\tPartOne: ${partOne}`);
    }
    const partTwo = daytwo.partTwo();
    if (partTwo !== null) {
        output.push(`\n\tPartTwo: ${partTwo}`);
    }
    return output.join("");
}

export { DayTwoRun, DayTwo };
