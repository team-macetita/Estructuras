#pragma once
#include "Node.h"
#include <iostream>

struct Tree {
    Node *root;
    void insert(int number);
    void preorder();
    void inorder();
    void postorder();

    private:
    Node *insert(int number, Node *node);
    void preorder(Node *node1);
    void inorder(Node *node1);
    void postorder(Node *node1);
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

void Tree::preorder(){
    this->preorder(this->root);
    std::cout << "" << std::endl;
}

void Tree::preorder(Node *node1){
    if (node1 != nullptr){
        std::cout << node1->number << std::endl;
        this->preorder(node1->left);
        this->preorder(node1->right);
    }
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

void Tree::postorder(){
    this->postorder(this->root);
    std::cout << "" << std::endl;
}

void Tree::postorder(Node *node1){
    if (node1 != nullptr){
        this->postorder(node1->left);
        this->postorder(node1->right);
        std::cout << node1->number << std::endl;
    }
}