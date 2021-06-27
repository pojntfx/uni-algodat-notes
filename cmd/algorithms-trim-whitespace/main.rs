#[path = "../../pkg/algorithms/trim_whitespace.rs"]
mod trim_whitespace;

fn main() {
    let inputs = vec![
        "test",
        "test asdf 9d2i",
        "   asdf dasdf
asdf",
        "asdf	asd sad",
        "yay! ",
    ];

    inputs.iter().for_each(|input| {
        let output_imperative = trim_whitespace::trim_whitespace_imperative(input);
        let output_functional =
            trim_whitespace::trim_whitespace_functional(input, String::new(), 0);

        println!(
            "in=\"{}\" out=\"{}\"=\"{}\"",
            input, output_imperative, output_functional
        )
    });
}
