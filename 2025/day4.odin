package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:os"
import "core:strings"
import "core:testing"

TEST_INPUT_DAY4 :: `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`


TEST_OUTPUT_DAY4_PART1 :: 13
TEST_OUTPUT_DAY4_PART2 :: 43

GridType :: enum {
	EMPTY,
	PAPER,
	REMOVE,
}

EightWayDirections :: enum {
	TOP,
	TOP_RIGHT,
	RIGHT,
	BOTTOM_RIGHT,
	BOTTOM,
	BOTTOM_LEFT,
	LEFT,
	TOP_LEFT,
}

Grid :: struct {
	data:          []GridType,
	width, height: int,
}

Vec2 :: struct {
	x: int,
	y: int,
}


parse_input_day4 :: proc(input: string) -> Grid {
	lines := strings.split(strings.trim_space(input), "\n")
	defer delete(lines)

	if len(lines) == 0 do return {}

	height := len(lines)
	width := len(lines[0])

	data := make([dynamic]GridType, 0, width * height)

	for line in lines {
		for char in line {
			switch char {
			case '.':
				append(&data, GridType.EMPTY)
			case '@':
				append(&data, GridType.PAPER)
			}
		}
	}

	return Grid{data[:], width, height}
}

@(test)
test_day4_part1 :: proc(t: ^testing.T) {
	inputs := parse_input_day4(TEST_INPUT_DAY4)

	total := day4_part1(inputs)
	testing.expect(t, total == TEST_OUTPUT_DAY4_PART1)
}

@(test)
test_day4_part2 :: proc(t: ^testing.T) {
	inputs := parse_input_day4(TEST_INPUT_DAY4)

	total := day4_part2(inputs)
	testing.expect(t, total == TEST_OUTPUT_DAY4_PART2)
}

day4_part1 :: proc(input: Grid) -> int {
	rolls := 0

	for data, i in input.data {
		curr := Vec2 {
			x = i % input.width,
			y = i / input.width,
		}

		if data == GridType.EMPTY {
			continue
		}

		rolls_around_curr := 0
		for direction in EightWayDirections {
			coords := Vec2{}
			switch direction {
			case EightWayDirections.TOP:
				if curr.y == 0 {
					continue
				}
				coords = Vec2 {
					x = curr.x,
					y = curr.y - 1,
				}
			case EightWayDirections.TOP_RIGHT:
				if curr.y == 0 || curr.x >= input.width - 1 {
					continue
				}
				coords = Vec2 {
					x = curr.x + 1,
					y = curr.y - 1,
				}
			case EightWayDirections.RIGHT:
				if curr.x >= input.width - 1 {
					continue
				}
				coords = Vec2 {
					x = curr.x + 1,
					y = curr.y,
				}
			case EightWayDirections.BOTTOM_RIGHT:
				if curr.x >= input.width - 1 || curr.y >= input.height - 1 {
					continue
				}
				coords = Vec2 {
					x = curr.x + 1,
					y = curr.y + 1,
				}
			case EightWayDirections.BOTTOM:
				if curr.y >= input.height - 1 {
					continue
				}
				coords = Vec2 {
					x = curr.x,
					y = curr.y + 1,
				}
			case EightWayDirections.BOTTOM_LEFT:
				if curr.y >= input.height - 1 || curr.x == 0 {
					continue
				}
				coords = Vec2 {
					x = curr.x - 1,
					y = curr.y + 1,
				}
			case EightWayDirections.LEFT:
				if curr.x == 0 {
					continue
				}
				coords = Vec2 {
					x = curr.x - 1,
					y = curr.y,
				}
			case EightWayDirections.TOP_LEFT:
				if curr.x == 0 || curr.y == 0 {
					continue
				}
				coords = Vec2 {
					x = curr.x - 1,
					y = curr.y - 1,
				}
			}
			index := coords.y * input.width + coords.x
			if input.data[index] == GridType.PAPER {
				rolls_around_curr += 1
			}
		}
		if rolls_around_curr < 4 {
			rolls += 1
		}
	}

	return rolls
}

day4_part2 :: proc(input: Grid) -> int {
	rolls_removed := 0

	for {
		rolls := 0

		for data, i in input.data {
			curr := Vec2 {
				x = i % input.width,
				y = i / input.width,
			}

			if data == GridType.EMPTY {
				continue
			}

			rolls_around_curr := 0
			for direction in EightWayDirections {
				coords := Vec2{}
				switch direction {
				case EightWayDirections.TOP:
					if curr.y == 0 {
						continue
					}
					coords = Vec2 {
						x = curr.x,
						y = curr.y - 1,
					}
				case EightWayDirections.TOP_RIGHT:
					if curr.y == 0 || curr.x >= input.width - 1 {
						continue
					}
					coords = Vec2 {
						x = curr.x + 1,
						y = curr.y - 1,
					}
				case EightWayDirections.RIGHT:
					if curr.x >= input.width - 1 {
						continue
					}
					coords = Vec2 {
						x = curr.x + 1,
						y = curr.y,
					}
				case EightWayDirections.BOTTOM_RIGHT:
					if curr.x >= input.width - 1 || curr.y >= input.height - 1 {
						continue
					}
					coords = Vec2 {
						x = curr.x + 1,
						y = curr.y + 1,
					}
				case EightWayDirections.BOTTOM:
					if curr.y >= input.height - 1 {
						continue
					}
					coords = Vec2 {
						x = curr.x,
						y = curr.y + 1,
					}
				case EightWayDirections.BOTTOM_LEFT:
					if curr.y >= input.height - 1 || curr.x == 0 {
						continue
					}
					coords = Vec2 {
						x = curr.x - 1,
						y = curr.y + 1,
					}
				case EightWayDirections.LEFT:
					if curr.x == 0 {
						continue
					}
					coords = Vec2 {
						x = curr.x - 1,
						y = curr.y,
					}
				case EightWayDirections.TOP_LEFT:
					if curr.x == 0 || curr.y == 0 {
						continue
					}
					coords = Vec2 {
						x = curr.x - 1,
						y = curr.y - 1,
					}
				}
				index := coords.y * input.width + coords.x
				if input.data[index] == GridType.PAPER || input.data[index] == GridType.REMOVE {
					rolls_around_curr += 1
				}
			}
			if rolls_around_curr < 4 {
				rolls += 1
				input.data[i] = GridType.REMOVE
			}
		}

		// reset the gridtype
		for data, i in input.data {
			if data == GridType.REMOVE {
				input.data[i] = GridType.EMPTY
			}
		}

		if rolls == 0 {
			break // reached the end
		}

		rolls_removed += rolls
	}

	return rolls_removed
}

main :: proc() {
	arena: vmem.Arena
	arena_err := vmem.arena_init_growing(&arena)
	ensure(arena_err == nil)
	defer vmem.arena_free_all(&arena)

	arena_alloc := vmem.arena_allocator(&arena)
	context.allocator = arena_alloc

	file, file_ok := os.read_entire_file("inputs/day4.txt", arena_alloc)
	ensure(file_ok)

	input_txt := string(file)
	inputs := parse_input_day4(input_txt)

	count_p1 := day4_part1(inputs)
	fmt.printf("final_count_part1: %d\n", count_p1)

	count_p2 := day4_part2(inputs)
	fmt.printf("final_count_part2: %d\n", count_p2)
}

