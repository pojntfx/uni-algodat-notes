#[derive(Debug)]
pub struct Node<T: Ord> {
    value: T,

    left: Box<Option<Node<T>>>,
    right: Box<Option<Node<T>>>,
}

impl<T: Ord> Node<T> {
    pub fn new(value: T) -> Self {
        Self {
            value,
            left: Box::new(None),
            right: Box::new(None),
        }
    }

    pub fn bin(&mut self, left: Option<Node<T>>, right: Option<Node<T>>) {
        if left.is_some() {
            self.left = Box::new(left);
        }

        if right.is_some() {
            self.right = Box::new(right);
        }
    }
}
