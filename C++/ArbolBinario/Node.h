#pragma once

struct Node {
    int number;
    Node *left;
    Node *right;
    Node(int number);
};

Node::Node(int number) : number(number), left(nullptr), right(nullptr) {}