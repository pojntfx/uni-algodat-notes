use std::convert::TryInto;

pub fn binary_search_functional(haystack: Vec<i32>, needle: i32, low: i32, high: i32) -> i32 {
    // Needle not in haystack
    if low > high {
        return -1;
    }

    /*
        To make this fully side-effect free, replace all references
        to the following with these expressions
    */
    let focus = (low + high) / 2;
    let center = haystack[focus as usize];

    // Needle at focus
    if needle == center {
        return focus;
    }

    // Take left part
    if needle < center {
        return binary_search_functional(haystack, needle, low, focus - 1);
    }

    // Take right part
    return binary_search_functional(haystack, needle, focus + 1, high);
}

pub fn binary_search_imperative(haystack: Vec<i32>, needle: i32) -> i32 {
    let mut low = 0;
    let mut high = haystack.len() - 1;

    loop {
        // Needle not in haystack
        if low > high {
            return -1;
        }

        let focus = (low + high) / 2;
        let center = haystack[focus as usize];

        // Needle at focus
        if needle == center {
            return focus.try_into().unwrap();
        }

        // Take left part
        if needle < center {
            high = focus - 1
        } else {
            // Take right part
            low = focus + 1
        }
    }
}
