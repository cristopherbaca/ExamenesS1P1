use std::io;

fn main() {
    let mut entrada = String::new();

    io::stdin().read_line(&mut entrada).unwrap();

    let c: f64 = entrada.trim().parse().unwrap();

    let f = (c * 9.0 / 5.0) + 32.0;
    let k = c + 273.15;

    println!("Celsius: {}, Fahrenheit: {}, Kelvin: {}", c, f, k);
}
