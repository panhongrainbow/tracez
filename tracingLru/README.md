# TracingLru

## Why use LRU in tracing logs？

### Reasons

1. I plan to use List, because `turning Log data into List`, I can write `an instruction interface to interact`, which can increase interaction.

2. Use `double linked list`, respectively do `two root nodes, root and error`.

   `root` is the source of all nodes, the direction of the key table is from `top to bottom`
   `error` connects to all places where errors occur, the direction is from `bottom up`

3. The problem is, every time a new node is added, `the entire node has to be traversed`.
   The problem is that I `can't guarantee that the logs are in order`.
   `The highly concurrency`, the order of log will `not` meet expectations.
   Directly add `a map` to speed up inserting new nodes.

Overall, I need to use a list to create an instruction interface and a map to speed up inserting new nodes.

The combination of `map and list` achieves `an LRU cache`.





























































