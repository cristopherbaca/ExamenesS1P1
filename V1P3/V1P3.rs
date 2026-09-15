fn main() {

    let a = [7, 8, 7, 9, 7, 8, 10];

    let mut mejor_valor = a[0];
    let mut mejor_frecuencia = 0;

    for i in 0..a.len() { // Recorremos el arreglo
        let mut frecuencia = 0;

        for j in 0..a.len() { // Buscamos cuántas veces aparece a[i]
            if a[i] == a[j] { // Si son iguales, aumentamos la frecuencia
                frecuencia += 1;
            }
        }

        if frecuencia > mejor_frecuencia { // Si encontramos una frecuencia mayor
            mejor_frecuencia = frecuencia;
            mejor_valor = a[i];
        }
    }

    println!("Mejor valor: {}", mejor_valor);
    println!("Frecuencia: {}", mejor_frecuencia);
}