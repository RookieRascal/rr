let tokenize query =
  query |> String.split_on_char ' '

let () =
  let query = "find the best search engine" in
  let tokens = tokenize query in
  List.iter print_endline tokens
