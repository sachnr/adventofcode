package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:os"
import "core:strconv"
import "core:strings"
import "core:testing"

TEST_INPUT_DAY5 :: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`


TEST_OUTPUT_DAY5_PART1 :: 3
TEST_OUTPUT_DAY5_PART2 :: 43

@(test)
test_day5_part1 :: proc(t: ^testing.T) {
	inputs := parse_input_day5(TEST_INPUT_DAY5)

	total := day5_part1(inputs)
	testing.expect(t, total == TEST_OUTPUT_DAY5_PART1)
}


Fresh_Range :: struct {
	start: int,
	end:   int,
}

Database :: struct {
	fresh_ranges:  []Fresh_Range,
	available_ids: []int,
}

parse_input_day5 :: proc(input: string) -> Database {
	line := strings.split(strings.trim_space(input), "\n\n")

	fresh_str := line[0]
	available_str := line[1]

	fresh_ranges := make([dynamic]Fresh_Range)
	fresh_ranges_str := strings.split(fresh_str, "\n")
	for range_str in fresh_ranges_str {
		parts := strings.split(range_str, "-")
		start, ok1 := strconv.parse_int(parts[0], 10)
		ensure(ok1)
		end, ok2 := strconv.parse_int(parts[1], 10)
		ensure(ok2)

		append(&fresh_ranges, Fresh_Range{start, end})
	}

	available_ids := make([dynamic]int)
	available_ids_str := strings.split(available_str, "\n")
	for id_str in available_ids_str {
		id, ok := strconv.parse_int(id_str, 10)
		ensure(ok)
		append(&available_ids, id)
	}

	return Database{fresh_ranges = fresh_ranges[:], available_ids = available_ids[:]}
}

day5_part1 :: proc(db: Database) -> int {
	count := 0
	for id in db.available_ids {
		for range in db.fresh_ranges {
			if id >= range.start && id <= range.end {
				count += 1
				break
			}
		}
	}
	return count
}

day5_part2 :: proc(db: Database) -> int {
	count := 0

	fresh_ranges_no_overlapp := make([dynamic]Fresh_Range)

	// 246125950983468-246389816279678
	// 249959673124024-250732034749615
	// ???
	for range in db.fresh_ranges {
		for our_range in fresh_ranges_no_overlapp {

		}
	}

	return count
}

main :: proc() {
	arena: vmem.Arena
	arena_err := vmem.arena_init_growing(&arena)
	ensure(arena_err == nil)
	defer vmem.arena_free_all(&arena)

	arena_alloc := vmem.arena_allocator(&arena)
	context.allocator = arena_alloc

	file, file_ok := os.read_entire_file("inputs/day5.txt", arena_alloc)
	ensure(file_ok)

	input_txt := string(file)
	inputs := parse_input_day5(input_txt)

	count_p1 := day5_part1(inputs)
	fmt.printf("final_count_part1: %d\n", count_p1)
}

