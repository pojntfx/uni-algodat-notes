use std::convert::TryInto;

pub fn trim_whitespace_imperative(input: String) -> String {
    let mut output = String::new();

    input.split("").for_each(|character| {
        if !is_whitespace(character) {
            output += character
        }
    });

    return output;
}

pub fn trim_whitespace_functional(input: String, output: String, i: i32) -> String {
    if i >= input.len().try_into().unwrap() {
        return output;
    }

    /*
        To make this fully side-effect free, replace all references
        to the following with these expressions
    */
    let curr = input.chars().nth(i.try_into().unwrap()).unwrap();

    if !is_whitespace(&curr.to_string()) {
        return trim_whitespace_functional(input, output + &curr.to_string(), i + 1);
    }

    return trim_whitespace_functional(input, output, i + 1);
}

fn is_whitespace(input: &str) -> bool {
    input.contains(char::is_whitespace)
}
