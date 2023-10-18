import * as fs from "fs";

class DayOne {
    entries: number[];

    constructor() {
        this.entries = [];
    }

    parse(input: string): Error | void {
        const lines = input.split("\n");
        for (const line of lines) {
            // check for empty lines
            if (line === "") {
                continue;
            }
            const intNo = parseInt(line, 10);
            if (isNaN(intNo)) {
                throw new Error(`Unable to parse line: \n${line}`);
            }
            this.entries.push(intNo);
        }
    }

    calculatePartOne(): number | null {
        const len = this.entries.length;
        for (let i = 0; i < len - 1; i++) {
            for (let j = i; j < len; j++) {
                const sum = this.entries[i] + this.entries[j];
                if (sum === 2020) {
                    return this.entries[i] * this.entries[j];
                }
            }
        }
        return null;
    }

    calculatePartTwo(): number | null {
        const len = this.entries.length;
        for (let i = 0; i < len - 2; i++) {
            for (let j = i; j < len - 1; j++) {
                for (let k = j; k < len; k++) {
                    const sum =
                        this.entries[i] + this.entries[j] + this.entries[k];
                    if (sum === 2020) {
                        return (
                            this.entries[i] * this.entries[j] * this.entries[k]
                        );
                    }
                }
            }
        }
        return null;
    }
}

function DayOneRun(): string {
    const dayOne = new DayOne();
    let output: String[] = [];
    try {
        const file = fs.readFileSync("src/inputs/day1.txt", "utf8");
        dayOne.parse(file);
    } catch (error: any) {
        console.error(`failed to parse the file: ${error.message}`);
    }
    const partOne = dayOne.calculatePartOne();
    output.push("Day One:");
    if (partOne !== null) {
        output.push(`\n\tPartOne: ${partOne}`);
    }
    const partTwo = dayOne.calculatePartTwo();
    if (partTwo !== null) {
        output.push(`\n\tPartTwo: ${partTwo}`);
    }
    return output.join("");
}

export { DayOneRun, DayOne };
