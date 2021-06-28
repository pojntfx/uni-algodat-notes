#[path = "../../pkg/algorithms/tree.rs"]
mod tree;

use tree::Node;

fn main() {
    let mut n1 = Node::new(10);
    let mut n2 = Node::new(11);
    let n3 = Node::new(12);
    let n4 = Node::new(13);
    let n5 = Node::new(14);

    n1.insert_right(n5);
    n2.insert_left(n4);
    n1.insert_left(n2);
    n1.insert_right(n3);

    dbg!(n1);
}
