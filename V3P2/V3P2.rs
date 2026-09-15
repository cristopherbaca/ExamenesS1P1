fn main() {
    let mut tabla1 = 252;
    let mut tabla2 = 105;
    let mut swap;
    while tabla1 != 0 {
        swap = tabla2 % tabla1;
        tabla2 = tabla1;
        tabla1 = swap;
    }
    println!("{tabla2}")
}