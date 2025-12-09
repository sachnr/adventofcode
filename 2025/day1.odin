package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:os"
import "core:strconv"
import "core:strings"
import "core:testing"

Direction :: enum {
	Right,
	Left,
}

Input :: struct {
	direction: Direction,
	count:     int,
}

count_zeroes_part1 :: proc(inputs: []Input) -> int {
	count := 0
	dial := 50

	for input in inputs {
		normalized_count := input.count % 100
		switch input.direction {
		case Direction.Right:
			dial = (dial + normalized_count) % 100
		case Direction.Left:
			dial = (dial - normalized_count + 100) % 100
		}
		if dial == 0 {
			count += 1
		}
	}
	return count
}

// how many times it hit zero
count_zeroes_part2 :: proc(inputs: []Input) -> int {
	count := 0
	dial := 50

	for input in inputs {
		normalized_count := input.count % 100
		laps := input.count / 100
		count += laps
		switch input.direction {
		case Direction.Right:
			if dial != 0 && (dial + normalized_count) >= 100 {
				count += 1
			}
			dial = (dial + normalized_count) % 100
		case Direction.Left:
			if dial != 0 && (dial - normalized_count + 100) <= 100 {
				count += 1
			}
			dial = (dial - normalized_count + 100) % 100
		}
	}
	return count
}

@(test)
test_count_zeroes :: proc(t: ^testing.T) {
	inputs := []Input {
		{direction = Direction.Left, count = 68},
		{direction = Direction.Left, count = 30},
		{direction = Direction.Right, count = 48},
		{direction = Direction.Left, count = 5},
		{direction = Direction.Right, count = 60},
		{direction = Direction.Left, count = 55},
		{direction = Direction.Left, count = 1},
		{direction = Direction.Left, count = 99},
		{direction = Direction.Right, count = 14},
		{direction = Direction.Left, count = 82},
	}

	count := count_zeroes_part1(inputs)
	testing.expectf(t, count == 3, "day1p1: the value is not three, value = %i", count)

	countp2 := count_zeroes_part2(inputs)
	testing.expectf(t, countp2 == 6, "day1p2: the value is not six, value = %i", countp2)
}

main :: proc() {
	arena: vmem.Arena
	arena_err := vmem.arena_init_growing(&arena)
	ensure(arena_err == nil)
	defer vmem.arena_free_all(&arena)

	arena_alloc := vmem.arena_allocator(&arena)
	context.allocator = arena_alloc

	file, file_ok := os.read_entire_file("inputs/day1.txt", arena_alloc)
	ensure(file_ok)

	input_txt := string(file)

	lines := strings.split(input_txt, "\n")

	inputs := make([dynamic]Input)

	for line in lines {
		if line == "" {
			continue
		}
		part1 := line[0]
		part2 := line[1:]

		count, ok := strconv.parse_int(string(part2), 10)
		ensure(ok)

		switch part1 {
		case 'R':
			append(&inputs, Input{count = count, direction = Direction.Right})
		case 'L':
			append(&inputs, Input{count = count, direction = Direction.Left})
		}
	}

	count_p1 := count_zeroes_part1(inputs[:])
	fmt.printf("final_count_part1: %d", count_p1)

	count_p2 := count_zeroes_part2(inputs[:])
	fmt.printf("final_count_part2: %d", count_p2)
}

