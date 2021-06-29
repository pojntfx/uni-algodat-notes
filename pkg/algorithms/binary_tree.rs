#[derive(Debug, Clone)]
pub struct Node<T: Ord> {
    value: T,

    left: Option<Box<Node<T>>>,
    right: Option<Box<Node<T>>>,
}

impl<T: Ord> Node<T> {
    pub fn new(value: T) -> Self {
        Self {
            value,
            left: None,
            right: None,
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
                self.left = Some(Box::new(Node::new(value)));

                return;
            }

            self.left.as_mut().as_mut().unwrap().insert(value);

            return;
        }

        // Add to right if parent is larger
        if self.right.is_none() {
            self.right = Some(Box::new(Node::new(value)));

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

    pub fn delete(mut self, value: T) -> Option<Box<Node<T>>> {
        if value < self.value {
            if let Some(left) = self.left.take() {
                self.left = left.delete(value);
            }

            return Some(Box::new(self));
        }

        if value > self.value {
            if let Some(right) = self.right.take() {
                self.right = right.delete(value);
            }

            return Some(Box::new(self));
        }

        match (self.left.take(), self.right.take()) {
            (None, None) => None,
            (Some(left), None) => Some(left),
            (None, Some(right)) => Some(right),
            (Some(mut left), Some(right)) => {
                if let Some(mut rightmost) = left.rightmost_child() {
                    rightmost.left = Some(left);
                    rightmost.right = Some(right);
                    Some(rightmost)
                } else {
                    left.right = Some(right);
                    Some(left)
                }
            }
        }
    }

    pub fn rightmost_child(&mut self) -> Option<Box<Node<T>>> {
        match self.right {
            Some(ref mut right) => {
                if let Some(t) = right.rightmost_child() {
                    Some(t)
                } else {
                    let mut r = self.right.take();
                    if let Some(ref mut r) = r {
                        self.right = std::mem::replace(&mut r.left, None);
                    }
                    r
                }
            }
            None => None,
        }
    }
}
