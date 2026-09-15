fn main() {

    let n = 97;
    let mut primo = true;

    if n < 2 { // Si el número es menor a 2, no es primo
        primo = false;
    } else {
        for i in 2..n { // Probamos los números desde 2 hasta n-1
            if n % i == 0 { // Si es divisible, no es primo
                primo = false;
                break;
            }
        }
    }

    println!("{}", primo);
}