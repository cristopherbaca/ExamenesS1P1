fn main() {
    let arr = [1, 2, 2, 3, 4, 4, 5, 5, 5, 7];
    let mut output = Vec::new();
    let mut encontrado;
    for i in 0..arr.len() {
        encontrado = false;
        for j in 0..output.len() {
            if arr[i] == output[j] {
                encontrado = true;
            }
        }
        if encontrado == false {
            output.push(arr[i])
        }
    }
    for x in &output {
        println!("{}", x);
    }
}
