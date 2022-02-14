use std::io;
use rand::prelude::*;

fn main() {
    println!("Guessing game: Enter a number to start.");
    let mut num = rand::thread_rng().gen_range(1..100);
    let mut guess_amount: u8 = 0;

    loop{
        
        let mut guess: String = String::new();

        
        io::stdin()
            .read_line(&mut guess)
            .expect("Couldn't aquire input.");

        let guess = guess.trim().parse::<i32>().expect("invalid input");



        if num == guess {
            println!("\nCorrect");
            println!("Number of guesses: {}", guess_amount);
            break;
        }else if num < guess {
            println!("Too big");
            guess_amount += 1;
        }else if num > guess {
            println!("Too small");
            guess_amount += 1;
        };

    }

}
