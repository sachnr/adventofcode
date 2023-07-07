pub fn run() {
    let input = include_str!("inputs/input3.txt");
    let mut rucksack = RuckSack::new();
    rucksack.parse(input);
    rucksack.parse_groups(input);
    println!();
    println!("day 3:");
    println!("part1: {}", rucksack.part1());
    println!("part2: {}", rucksack.part2());
}

#[derive(Default)]
struct RuckSack {
    prioties: Vec<u8>,
    priorti_group: Vec<u8>,
}

impl RuckSack {
    fn new() -> Self {
        Self::default()
    }

    fn parse(&mut self, input: &str) {
        for line in input.lines() {
            let (part1, part2) = line.split_at(line.len() / 2);
            let value =
                part1
                    .chars()
                    .find(|&value| part2.contains(value))
                    .map(|value| match value {
                        'a'..='z' => value as u8 - b'a' + 1,
                        'A'..='Z' => value as u8 - b'A' + 27,
                        _ => panic!("invalid char: {}", value),
                    });
            if let Some(value) = value {
                self.prioties.push(value);
            }
        }
    }

    fn parse_groups(&mut self, input: &str) {
        let vec_of_lines = input.lines().collect::<Vec<_>>();
        for arr in vec_of_lines.chunks(3) {
            let value = arr[0]
                .chars()
                .find(|&item| arr[1].contains(item) && arr[2].contains(item))
                .map(|value| match value {
                    'a'..='z' => value as u8 - b'a' + 1,
                    'A'..='Z' => value as u8 - b'A' + 27,
                    _ => panic!("invalid char: {}", value),
                });
            if let Some(value) = value {
                self.priorti_group.push(value);
            }
        }
    }

    fn part1(&self) -> i32 {
        self.prioties.iter().map(|&val| val as i32).sum()
    }

    fn part2(&self) -> i32 {
        self.priorti_group.iter().map(|&val| val as i32).sum()
    }
}

mod tests {
    #[test]
    fn day3() {
        use super::RuckSack;
        let input = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
        let mut rucksack = RuckSack::new();
        rucksack.parse(input);
        rucksack.parse_groups(input);
        assert_eq!(rucksack.part1(), 157);
        assert_eq!(rucksack.part2(), 70);
    }
}
