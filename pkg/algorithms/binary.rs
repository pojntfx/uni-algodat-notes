pub fn get_as_binary_imperative(num: i32) -> String {
    let mut output = String::new();

    let mut i = num;
    while i > 0 {
        output = (i % 2).to_string() + &output;

        i = i / 2;
    }

    return output;
}

pub fn get_as_binary_functional(num: i32, output: String) -> String {
    if num == 0 {
        return output;
    }

    return get_as_binary_functional(num / 2, (num % 2).to_string() + &output);
}
