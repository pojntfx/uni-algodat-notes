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

    // See https://stackoverflow.com/questions/64043682/how-to-write-a-delete-function-for-a-binary-tree-in-rust
    pub fn delete(mut this: Box<Node<T>>, target: &T) -> Option<Box<Node<T>>> {
        if target < &this.value {
            if let Some(left) = this.left.take() {
                this.left = Self::delete(left, target);
            }
            return Some(this);
        }

        if target > &this.value {
            if let Some(right) = this.right.take() {
                this.right = Self::delete(right, target);
            }
            return Some(this);
        }

        assert!(target == &this.value, "Faulty Ord implementation for T");

        match (this.left.take(), this.right.take()) {
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

    //  Returns the rightmost child, unless the node itself is that child.
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
