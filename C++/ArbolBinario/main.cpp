#include "Tree.h"

Tree *arbolBinario = new Tree();

int main() {
    arbolBinario->insert(7);
    arbolBinario->insert(1);
    arbolBinario->insert(18);
    arbolBinario->insert(27);
    arbolBinario->insert(6);
    arbolBinario->insert(5);
    std::cout << " --Preorder-- " << std::endl;
    arbolBinario->preorder();
    std::cout << " --Inorder-- " << std::endl;
    arbolBinario->inorder();
    std::cout << " --Postorder-- " << std::endl;
    arbolBinario->postorder();
    return 0;
}