use rand::Rng;

use clap::{Arg, Command};

fn create_numbers(max : u16, count : u16) -> Vec<u16> {
    let mut balls : Vec<u16> = (1..=max).into_iter().collect();
    let mut rng = rand::thread_rng();
    let mut result : Vec<u16> = Vec::new();


    for _ in 0..count {
        let i = rng.gen_range(0..balls.len());
        let ball = balls.remove(i);
        result.push(ball);
    }

    result
}

fn main() {
    let matches = Command::new("sixoutof49")
        .version("0.1.0")
        .about("Feeling lucky? Play lotto with 6 out of 49!")
        .arg(Arg::new("count").long("count").default_value("6").value_parser(clap::value_parser!(u16).range(1..=1000)).help("Number of draws"))
        .arg(Arg::new("max").long("max").default_value("49").value_parser(clap::value_parser!(u16).range(1..=1000)) .help("Number of balls")).get_matches();

    let count : &u16 = matches.get_one("count").unwrap();
    let max : &u16 = matches.get_one("max").unwrap();

    let numbers = create_numbers(*max, *count);
    let numbers_str : Vec<_>= numbers.iter().map(|v| format!("{}", *v)).collect();


    println!("sixoutof49: your lucky numbers are: {}", numbers_str.join(" "));
}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(1, 1)]
    #[case(1000, 1000)]
    fn test_create_numbers(#[case] max : u16, #[case] count: u16) {
        let numbers = create_numbers(max, count);

        assert_eq!(numbers.len() as u16, count);
        assert!(*numbers.iter().min().unwrap() >= 1);
        assert!(*numbers.iter().max().unwrap() <= max);
    }
}