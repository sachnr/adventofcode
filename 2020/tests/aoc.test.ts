import { describe, expect, test } from "@jest/globals";
import { DayOne } from "../src/day1";
import { DayTwo } from "../src/day2";

describe("Day One", () => {
    const input = `1721
979
366
299
675
1456`;
    const day1 = new DayOne();
    expect(() => day1.parse(input)).not.toThrow();

    test("part1", () => {
        let part1 = day1.calculatePartOne();
        expect(part1).toBe(514579);
    });

    test("part2", () => {
        let part2 = day1.calculatePartTwo();
        expect(part2).toBe(241861950);
    });
});

describe("Day Two", () => {
    const input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`;
    const day2 = new DayTwo();
    expect(() => day2.parse(input)).not.toThrow();

    test("part1", () => {
        const part1 = day2.partOne();
        expect(part1).toBe(2);
    });
    test("part2", () => {
        const part2 = day2.partTwo();
        expect(part2).toBe(1);
    });
});
