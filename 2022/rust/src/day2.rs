pub fn run() {
    let input = include_str!("inputs/input2.txt");
    let mut scores = Scores::new();
    scores.parse(input);
    println!();
    println!("day 2:");
    println!("part1: {}", scores.part1());
    println!("part2: {}", scores.part2());
}

#[derive(Default)]
struct Scores<'a> {
    moves: Vec<(&'a str, &'a str)>,
}

impl<'a> Scores<'a> {
    fn new() -> Self {
        Self::default()
    }

    fn parse(&mut self, inputs: &'a str) {
        self.moves = inputs
            .lines()
            .map(|line| {
                let mut words = line.split(' ');
                let first = words.next().unwrap();
                let second = words.next().unwrap();

                (first, second)
            })
            .collect();
    }

    fn part1(&self) -> i32 {
        let mut total_score = 0;
        for (p1, p2) in self.moves.iter() {
            match (*p1, *p2) {
                ("A", "X") => total_score += 4,
                ("A", "Y") => total_score += 8,
                ("A", "Z") => total_score += 3,
                ("B", "X") => total_score += 1,
                ("B", "Y") => total_score += 5,
                ("B", "Z") => total_score += 9,
                ("C", "X") => total_score += 7,
                ("C", "Y") => total_score += 2,
                ("C", "Z") => total_score += 6,
                _ => {}
            }
        }
        total_score
    }

    fn part2(&self) -> i32 {
        let mut total_score = 0;
        for (p1, p2) in self.moves.iter() {
            match (*p1, *p2) {
                ("A", "X") => total_score += 3,
                ("A", "Y") => total_score += 4,
                ("A", "Z") => total_score += 8,
                ("B", "X") => total_score += 1,
                ("B", "Y") => total_score += 5,
                ("B", "Z") => total_score += 9,
                ("C", "X") => total_score += 2,
                ("C", "Y") => total_score += 6,
                ("C", "Z") => total_score += 7,
                _ => {}
            }
        }
        total_score
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn day2() {
        let input = "A Y
B X
C Z";
        let mut scores = Scores::new();
        scores.parse(input);
        assert_eq!(scores.part1(), 15);
        assert_eq!(scores.part2(), 12);
    }
}
