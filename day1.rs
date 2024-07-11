use std::env;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn char_is_integer(ch: char) -> bool {
  return ch.is_ascii_digit()
}

fn parse_first_and_last_number(numbers: String) -> String {
  let mut chars = numbers.chars();

  let first = chars.next().unwrap();
  let last = chars.next_back().unwrap_or(first);

  let mut first_and_last = String::new();
  first_and_last.push(first);
  first_and_last.push(last);

  return first_and_last;
}

fn main() {
  let args: Vec<String> = env::args().collect();
  let path = Path::new(&args[1]);
  let file = match File::open(&path) {
    Err(why) => panic!("Could not read path: {:?}. Why: {:?}", path, why),
    Ok(f) => f
  };

  let lines = io::BufReader::new(file).lines();

  let mut numbers = Vec::new();

  for line in lines {

    let mut extracted_numbers = String::new();

    for ch in line.unwrap().chars() {
      if char_is_integer(ch) {
        extracted_numbers.push(ch)
      }
    }

    numbers.push(extracted_numbers);
  }

  let mut integers: Vec<usize> = Vec::new();

  for num in numbers {
    let numbers_as_string = parse_first_and_last_number(num);

    integers.push(numbers_as_string.parse::<usize>().unwrap());
  }

  println!("{:?}", integers.into_iter().sum::<usize>());
}