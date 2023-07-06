use std::{collections::BinaryHeap, num::ParseIntError};

pub fn run() {
    let mut calories = Calories::new();
    let input = include_str!("inputs/input1.txt");
    if let Err(e) = calories.parse(input) {
        println!("Error Parsing: {}", e);
    }
    println!("day1:");
    println!("part1: {}", calories.part1());
    println!("part2: {}", calories.part2());
}

#[derive(Default)]
struct Calories {
    items: BinaryHeap<i32>,
}

impl Calories {
    fn new() -> Self {
        Default::default()
    }

    fn parse(&mut self, input: &str) -> Result<(), ParseIntError> {
        let mut lines = input.split("\n\n");
        while let Some(lines) = lines.next() {
            let mut total = 0;
            for line in lines.lines() {
                total += line.parse::<i32>()?;
            }
            self.items.push(total);
        }
        Ok(())
    }

    fn part1(&self) -> i32 {
        *self.items.peek().unwrap()
    }

    fn part2(&mut self) -> i32 {
        let mut total = 0;
        for _ in 0..3 {
            total += self.items.pop().unwrap();
        }
        total
    }
}
