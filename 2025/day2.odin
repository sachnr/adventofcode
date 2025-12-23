package main

import "core:fmt"
import "core:log"
import vmem "core:mem/virtual"
import "core:os"
import "core:strconv"
import "core:strings"
import "core:testing"


TEST_INPUT_DAY2 :: "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
TEST_OUTPUT_DAY2_PART1 :: 1227775554
TEST_OUTPUT_DAY2_PART2 :: 4174379265


Range :: struct {
	start: int,
	end:   int,
}

parse_input_day2 :: proc(input: string) -> []Range {
	ranges := strings.split(input, ",")
	defer delete(ranges)

	output := make([dynamic]Range)

	for range in ranges {
		cleaned := strings.trim_space(range)
		parts := strings.split(cleaned, "-")
		defer delete(parts)
		start, ok1 := strconv.parse_int(parts[0], 10)
		ensure(ok1)
		end, ok2 := strconv.parse_int(parts[1], 10)
		ensure(ok2)
		append(&output, Range{start = start, end = end})
	}

	return output[:]
}

day2_part1 :: proc(inputs: []Range) -> int {
	invalids := 0

	for input in inputs {
		for value in input.start ..= input.end {
			digits, count, val := [20]int{}, 0, value
			for val > 0 {
				digits[count] = val % 10
				val = val / 10
				count += 1
			}
			// odd numbers are valid no repeating pattern
			if count % 2 != 0 {
				continue
			}
			half_count := count / 2
			valid := false
			for i := 0; i < half_count; i += 1 {
				if digits[i] != digits[half_count + i] {
					valid = true
				}
			}
			if !valid {
				invalids += value
			}
		}
	}

	return invalids
}

day2_part2 :: proc(inputs: []Range) -> int {
	invalids := 0

	for input in inputs {
		for value in input.start ..= input.end {
			digits, count, val := [20]int{}, 0, value
			for val > 0 {
				digits[count] = val % 10
				val = val / 10
				count += 1
			}

			invalid := false
			for chunk in 1 ..= count {
				if count / chunk < 2 {
					break
				}
				if count % chunk != 0 {
					// not a perfect sequence
					continue
				}
				prev := 0
				next := prev + chunk
				all_equal := true
				for next < count {
					for i := 0; i < chunk; i += 1 {
						if digits[prev + i] != digits[next + i] {
							all_equal = false
						}
					}
					prev = next
					next = prev + chunk
				}
				if all_equal {
					invalids += value
					break
				}
			}
		}
	}

	return invalids
}

@(test)
test_day2_part1 :: proc(t: ^testing.T) {
	inputs := parse_input_day2(TEST_INPUT_DAY2)
	defer delete(inputs)

	total := day2_part1(inputs)
	testing.expect(t, total == TEST_OUTPUT_DAY2_PART1)
}

@(test)
test_day2_part2 :: proc(t: ^testing.T) {
	inputs := parse_input_day2(TEST_INPUT_DAY2)
	defer delete(inputs)

	total := day2_part2(inputs)
	log.infof("output_day2: %d", total)
	testing.expect(t, total == TEST_OUTPUT_DAY2_PART2)
}

main :: proc() {
	arena: vmem.Arena
	arena_err := vmem.arena_init_growing(&arena)
	ensure(arena_err == nil)
	defer vmem.arena_free_all(&arena)

	arena_alloc := vmem.arena_allocator(&arena)
	context.allocator = arena_alloc

	file, file_ok := os.read_entire_file("inputs/day2.txt", arena_alloc)
	ensure(file_ok)

	input_txt := string(file)
	inputs := parse_input_day2(input_txt)

	count_p1 := day2_part1(inputs)
	fmt.printf("final_count_part1: %d\n", count_p1)

	count_p2 := day2_part2(inputs)
	fmt.printf("final_count_part2: %d\n", count_p2)
}

