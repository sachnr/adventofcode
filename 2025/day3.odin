package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:os"
import "core:strings"
import "core:testing"

TEST_INPUT_DAY3 :: `987654321111111
811111111111119
234234234234278
818181911112111`


TEST_OUTPUT_DAY3_PART1 :: 357
TEST_OUTPUT_DAY3_PART2 :: 3121910778619

Bank :: [dynamic]int

parse_input_day3 :: proc(input: string) -> []Bank {
	lines := strings.split(input, "\n")
	defer delete(lines)

	n := len(lines)
	banks: [dynamic]Bank
	reserve(&banks, n)

	for line, i in lines {
		bank := make([dynamic]int, 0, len(line))
		for char in line {
			if char < '0' || char > '9' {
				panic(fmt.tprintf("out of range char found: %v (%r)", char, char))
			}
			append(&bank, int(char - '0'))
		}
		append(&banks, bank)
		banks[i] = bank
	}

	return banks[:]
}

day3_part1 :: proc(input: []Bank) -> int {
	joltage := 0

	for bank in input {
		max := 0
		for i := 0; i < len(bank); i += 1 {
			for j := i + 1; j < len(bank); j += 1 {
				num := bank[i] * 10 + bank[j]
				if num > max {
					max = num
				}
			}
		}
		joltage += max
	}

	return joltage
}

ipow10 :: proc(exp: int) -> u64 {
	result: u64 = 1
	for _ in 0 ..< exp {
		prev := result
		result *= 10
	}
	return result
}

day3_part2 :: proc(input: []Bank) -> u64 {
	joltage: u64 = 0
	for bank in input {
		result: [12]int = {}
		remaining := 12
		loop_start := 0
		result_index := 0

		for remaining > 0 {
			loop_end := len(bank) - remaining

			max := 0
			max_index := loop_start
			for i in loop_start ..= loop_end {
				if max < bank[i] {
					max = bank[i]
					max_index = i
				}
			}

			result[result_index] = max
			loop_start = max_index + 1
			result_index += 1
			remaining -= 1
		}

		bank_joltage: u64 = 0
		pow := 11
		for max in result {
			bank_joltage += u64(max) * ipow10(pow)
			pow -= 1
		}
		joltage += bank_joltage
	}
	return joltage
}

@(test)
test_day3_part1 :: proc(t: ^testing.T) {
	inputs := parse_input_day3(TEST_INPUT_DAY3)
	defer delete(inputs)

	total := day3_part1(inputs)
	testing.expect(t, total == TEST_OUTPUT_DAY3_PART1)
}

@(test)
test_day3_part2 :: proc(t: ^testing.T) {
	inputs := parse_input_day3(TEST_INPUT_DAY3)
	defer delete(inputs)

	total := day3_part2(inputs)
	testing.expect(t, total == TEST_OUTPUT_DAY3_PART2)
}


main :: proc() {
	arena: vmem.Arena
	arena_err := vmem.arena_init_growing(&arena)
	ensure(arena_err == nil)
	defer vmem.arena_free_all(&arena)

	arena_alloc := vmem.arena_allocator(&arena)
	context.allocator = arena_alloc

	file, file_ok := os.read_entire_file("inputs/day3.txt", arena_alloc)
	ensure(file_ok)

	input_txt := string(file)
	inputs := parse_input_day3(input_txt)

	count_p1 := day3_part1(inputs)
	fmt.printf("final_count_part1: %d\n", count_p1)

	count_p2 := day3_part2(inputs)
	fmt.printf("final_count_part2: %d\n", count_p2)
}

