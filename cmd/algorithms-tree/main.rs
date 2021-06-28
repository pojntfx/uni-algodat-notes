#[path = "../../pkg/algorithms/tree.rs"]
mod tree;

use tree::Node;

fn main() {
    let mut n1 = Node::new(10);
    let mut n2 = Node::new(11);
    let n3 = Node::new(12);
    let n4 = Node::new(13);
    let n5 = Node::new(14);

    n1.bin(None, Some(n5));
    n2.bin(Some(n4), None);
    n1.bin(Some(n2), Some(n3));

    dbg!(n1);
}
