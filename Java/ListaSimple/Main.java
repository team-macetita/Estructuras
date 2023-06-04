package Java.ListaSimple;

class Nodo {
    int valor;
    Nodo siguiente;

    public Nodo(int valor) {
        this.valor = valor;
        siguiente = null;
    }
}

class ListaEnlazada {
    Nodo primero;
    Nodo ultimo;

    public ListaEnlazada() {
        primero = null;
        ultimo = null;
    }
}

public class Main {
    public static void main(String[] args) {
        // Aquí puedes escribir el código que se ejecutará cuando se ejecute el programa

        // Por ejemplo, imprimir un mensaje en la consola
        System.out.println("Hola, mundo!");
    }

}
