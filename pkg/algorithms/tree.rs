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

    pub fn insert_left(&mut self, value: Node<T>) {
        self.left = Box::new(Some(value));
    }

    pub fn insert_right(&mut self, value: Node<T>) {
        self.right = Box::new(Some(value));
    }
}
