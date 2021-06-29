#[path = "../../pkg/algorithms/tree.rs"]
mod tree;

use tree::Node;

fn main() {
    let mut root = Node::new(10);
    let mut n1 = Node::new(11);
    let n2 = Node::new(12);
    let n3 = Node::new(13);
    let n4 = Node::new(14);

    root.insert_right(n4);
    n1.insert_left(n3);
    root.insert_left(n1);
    root.insert_right(n2);

    dbg!(root);
}
