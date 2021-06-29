#[path = "../../pkg/algorithms/binary_tree.rs"]
mod binary_tree;

use binary_tree::Node;

fn main() {
    let mut root = Node::new(15);

    root.insert(11);
    root.insert(24);
    root.insert(5);
    root.insert(19);
    root.insert(16);

    root = dbg!(root);

    dbg!(root.contains(5));
    dbg!(root.contains(6));
    dbg!(root.contains(19));

    let root2 = root.delete(11);

    dbg!(root2);
}
