fn main() {
    let mut seconds: i32 = 9385;
    let hour: i32 = seconds / 3600;
    seconds = seconds % 3600;
    let minute: i32 = seconds / 60;
    seconds = seconds % 60;
    println!("{hour}:{minute}:{seconds}")
}
