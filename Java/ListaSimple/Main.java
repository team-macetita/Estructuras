package Java.ListaSimple;

class Nodo {
    int valor;
    Nodo siguiente;

    public Nodo(int valor) {
        this.valor = valor;
        siguiente = null;
    }
}

class ListaSimple {
    Nodo primero;
    Nodo ultimo;
    int longitud;

    public ListaSimple() {
        primero = null;
        ultimo = null;
        longitud = 0;
    }

    public void Insertar(int valor) {
        if (this.primero != null) {
            this.ultimo.siguiente = new Nodo(valor);
            this.ultimo = this.ultimo.siguiente;
            this.longitud += 1;
            return;
        }
        this.primero = new Nodo(valor);
        this.ultimo = this.primero;
        this.longitud += 1;
    }

    public String recorrer () {
        String lista = "";
        Nodo actual = this.primero;
        while (actual != null) {
            lista += " -> " + actual.valor;
            actual = actual.siguiente;
        }
        return lista;
    }
}

public class Main {
    public static void main(String[] args) {
        System.out.println("Lista Simple");
        ListaSimple l1 = new ListaSimple();
        l1.Insertar(1);
        l1.Insertar(2);
        l1.Insertar(2);
        l1.Insertar(3);
        l1.Insertar(4);
        System.out.println(l1.recorrer());
    }

}
