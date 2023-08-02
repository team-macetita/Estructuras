#pragma once
#include "Node.h"
#include <iostream>

struct Tree {
    Node *root;
    void insert(int number);
    void inorder();

    private:
    Node *insert(int number, Node *node);
    void inorder(Node *node1);
};

void Tree::insert(int number){
    this->root = this->insert(number, this->root);
}

Node *Tree::insert(int number, Node *node){
    if (node == nullptr){
        return new Node(number);
    }
    if (number < node->number) {
        node->left = this->insert(number, node->left);
    }
    else if (number > node->number) {
        node->right = this->insert(number, node->right);
    }
    return node;
}

void Tree::inorder(){
    this->inorder(this->root);
    std::cout << "" << std::endl;
}

void Tree::inorder(Node *node1){
    if (node1 != nullptr) {
        this->inorder(node1->left);
        std::cout << node1->number << std::endl;
        this->inorder(node1->right);
    }
}