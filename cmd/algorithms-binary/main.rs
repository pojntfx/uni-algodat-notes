#[path = "../../pkg/algorithms/binary.rs"]
mod binary;

fn main() {
    let inputs = vec![5, 6, 9, 15, 31, 66, 242];

    inputs.iter().for_each(|input| {
        let output_imperative = binary::get_as_binary_imperative(*input);
        let output_functional = binary::get_as_binary_functional(*input, String::new());

        println!(
            "in={:<3} out={:0>8}={:0>8}",
            input, output_imperative, output_functional
        )
    });
}
