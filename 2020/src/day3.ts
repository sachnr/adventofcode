import * as fs from "fs";

enum Move {
    Right,
    Down,
}

class DayThree {
    entries: boolean[][];
    posx: number;
    posy: number;

    constructor() {
        this.entries = [];
        this.posx = 0;
        this.posy = 0;
    }

    parse(input: string): Error | void {
        const lines = input.split("\n");
        const rows = lines[0].length;
        const cols = lines.length;
        this.entries = new Array(cols)
            .fill(null)
            .map(() => new Array(rows).fill(false));

        lines.forEach((line, index) => {
            if (line === "") {
                return;
            }
            for (let i = 0; i < line.length; i++) {
                switch (line[i]) {
                    case ".":
                        this.entries[index][i] = false;
                        break;
                    case "#":
                        this.entries[index][i] = true;
                        break;
                    default:
                        throw new Error(`unknown value: ${line[i]}`);
                }
            }
        });
    }

    Move(direction: Move) {
        const rows = this.entries[0].length;
        const cols = this.entries.length;
        switch (direction) {
            case Move.Right:
                if (this.posx < rows - 1) {
                    this.posx += 1;
                } else if (this.posx >= rows - 1) {
                    this.posx = 0;
                }
                break;
            case Move.Down:
                if (this.posy < cols) {
                    this.posy += 1;
                }
                break;
        }
    }

    slopeMovement(right: number, down: number) {
        for (let i = 0; i < right; i++) {
            this.Move(Move.Right);
        }
        for (let i = 0; i < down; i++) {
            this.Move(Move.Down);
        }
    }

    reset() {
        this.posx = 0;
        this.posy = 0;
    }

    partOne(): number {
        this.reset();
        let treesFound = 0;
        const cols = this.entries.length;
        this.slopeMovement(3, 1);
        while (this.posy < cols) {
            if (this.entries[this.posy][this.posx]) {
                treesFound += 1;
            }
            this.slopeMovement(3, 1);
        }
        return treesFound;
    }

    debug() {
        this.entries.forEach((row) => console.log(row.join(" ")));
    }

    partTwo(): number {
        const slopes = [
            [1, 1],
            [3, 1],
            [5, 1],
            [7, 1],
            [1, 2],
        ];

        let output = 1;
        const cols = this.entries.length;
        slopes.forEach((slope) => {
            this.reset();
            let treesFound = 0;
            this.slopeMovement(slope[0], slope[1]);
            while (this.posy < cols) {
                if (this.entries[this.posy][this.posx]) {
                    treesFound += 1;
                }
                this.slopeMovement(slope[0], slope[1]);
            }
            output *= treesFound;
        });
        return output;
    }
}

function RunDayThree(): string {
    const dayThree = new DayThree();
    let output: string[] = [];
    try {
        const file = fs.readFileSync("src/inputs/day3.txt", "utf8");
        dayThree.parse(file);
    } catch (error: any) {
        console.error(`Failed to parse the file: ${error.message}`);
        process.exit(1);
    }
    output.push(`Day Two: `);
    const partOne = dayThree.partOne();
    output.push(`\n\tPartOne: ${partOne}`);
    const partTwo = dayThree.partTwo();
    output.push(`\n\tPartTwo: ${partTwo}`);
    return output.join("");
}

export { DayThree, RunDayThree };
