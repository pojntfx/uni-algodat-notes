#[path = "../../pkg/algorithms/binary_tree.rs"]
mod binary_tree;

use binary_tree::Node;

fn main() {
    let mut n1 = Node::new(15);

    n1.insert(11);
    n1.insert(24);
    n1.insert(5);
    n1.insert(19);
    n1.insert(16);

    dbg!(n1);
}
