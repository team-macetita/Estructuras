package Java.ListaDoble;

class Nodo {
    int valor;
    Nodo siguiente;
    Nodo anterior;

    public Nodo(int valor) {
        this.valor = valor;
        siguiente = null;
    }
}

class ListaDoble {
    Nodo primero;
    Nodo ultimo;

    public ListaDoble() {
        primero = null;
        ultimo = null;
    }

    public void Insertar(int valor) {
        if (this.primero != null) {
            this.ultimo.siguiente = new Nodo(valor);
            this.ultimo.siguiente.anterior = this.ultimo;
            this.ultimo = this.ultimo.siguiente;
            return;
        }
        this.primero = new Nodo(valor);
        this.ultimo = this.primero;
    }

    public String recorrer() {
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
    public static void main(String [] args) {
        System.out.println("LIsta Doble");
        ListaDoble l1 = new ListaDoble();
        l1.Insertar(1);
        l1.Insertar(2);
        l1.Insertar(3);
        l1.Insertar(4);
        l1.Insertar(5);
        System.out.println(l1.recorrer());
    }
}