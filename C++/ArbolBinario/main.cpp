#include "Tree.h"

Tree *arbolBinario = new Tree();

int main() {
    arbolBinario->insert(7);
    arbolBinario->insert(1);
    arbolBinario->insert(18);
    arbolBinario->insert(27);
    arbolBinario->insert(6);
    arbolBinario->insert(5);
    arbolBinario->inorder();
    return 0;
}