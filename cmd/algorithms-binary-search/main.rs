use std::convert::TryInto;

#[path = "../../pkg/algorithms/binary_search.rs"]
mod binary_search;

fn main() {
    let inputs = vec![
        vec![1, 6, 9, 10, 11],
        vec![1, 2, 3, 4, 5],
        vec![3, 4, 5, 7, 12, 50],
        vec![3, 4, 6, 7, 12, 50],
        vec![1, 2, 3, 4, 5, 6, 12, 50],
    ];
    let target = 6;

    inputs.iter().for_each(|input| {
        let output_imperative = binary_search::binary_search_imperative(input.to_vec(), target);
        let output_functional = binary_search::binary_search_functional(
            input.to_vec(),
            target,
            0,
            (input.len() - 1).try_into().unwrap(),
        );

        println!(
            "in={:?} out={}={}",
            input, output_imperative, output_functional
        )
    })
}
