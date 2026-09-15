#include <iostream>

int main() {

    int calificaciones[] = {7, 8, 7, 9, 7, 8, 10};
    int frecuencia[] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0};

    // Contar cuántas veces aparece cada calificación
    int cantidadCalificaciones = sizeof(calificaciones) / sizeof(calificaciones[0]);

    for (int i = 0; i < cantidadCalificaciones; i++) {
        frecuencia[calificaciones[i] - 1]++;
    }

    // Buscar la calificación que aparece más veces
    int mayor = 0;
    int veces = frecuencia[0];

    int cantidadFrecuencias = sizeof(frecuencia) / sizeof(frecuencia[0]);

    for (int i = 1; i < cantidadFrecuencias; i++) {
        if (frecuencia[i] > veces) {
            veces = frecuencia[i];
            mayor = i;
        }
    }

    // Sumamos 1 porque el índice 6 representa la calificación 7
    std::cout << "Calificacion: " << mayor + 1
              << ", veces: " << veces << "\n";

    return 0;
}
