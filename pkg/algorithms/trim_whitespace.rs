pub fn trim_whitespace_imperative(input: &str) -> String {
    return input
        .chars()
        .filter(|character| !character.is_whitespace())
        .collect();
}

// Alternative, less efficient solution
// pub fn trim_whitespace_imperative(input: &str) -> String {
// let mut output = String::new();
//
// input.chars().for_each(|character| {
// if !is_whitespace(character) {
// output.push(character)
// }
// });
//
// return output;
// }

pub fn trim_whitespace_functional(input: &str, output: String, i: usize) -> String {
    if i >= input.len() {
        return output;
    }

    /*
        To make this fully side-effect free, replace all references
        to the following with these expressions
    */
    let curr = input.chars().nth(i).unwrap();

    if !curr.is_whitespace() {
        return trim_whitespace_functional(input, output + &curr.to_string(), i + 1);
    }

    return trim_whitespace_functional(input, output, i + 1);
}
