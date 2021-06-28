#[derive(Debug, Clone)]
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

    pub fn insert(&mut self, value: T) {
        // Prevent duplicates
        if value == self.value {
            return;
        }

        // Add to left if parent is smaller
        if value < self.value {
            if self.left.is_none() {
                self.left = Box::new(Some(Node::new(value)));

                return;
            }

            self.left.as_mut().as_mut().unwrap().insert(value);

            return;
        }

        // Add to right if parent is larger
        if self.right.is_none() {
            self.right = Box::new(Some(Node::new(value)));

            return;
        }

        self.right.as_mut().as_mut().unwrap().insert(value);
    }

    pub fn contains(&mut self, value: T) -> bool {
        // Found
        if value == self.value {
            return true;
        }

        // If parent is smaller, check left
        if value < self.value {
            if self.left.is_none() {
                return false;
            }

            return self.left.as_mut().as_mut().unwrap().contains(value);
        }

        // If parent is larger, check right
        if self.right.is_none() {
            return false;
        }

        return self.right.as_mut().as_mut().unwrap().contains(value);
    }
}
