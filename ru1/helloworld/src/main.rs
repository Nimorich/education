use std::io;

fn main() {
    println!("Пожалуйста, введите ваше имя:");
    let mut name = String::new();

    io::stdin()
        .read_line(&mut name)
        .expect("Не удалось прочитать строку");
    println!("Привет, {}", name);
    
    println!("Пожалуйста, введите ваш возраст:");
    let mut age = String::new();

    io::stdin()
        .read_line(&mut age)
        .expect("Не удалось прочитать строку");
    let age: u32 = age.trim().parse().expect("Введите число!");
    println!("Через год вам будет: {}", age + 1);
}
