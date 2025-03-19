We are collaborating on a Go program, codename `monogram`. The `monogram` tool
translates from `monogram` notation into XML, JSON and other formats. The 
notation is designed to represent program-like texts. However it is just a
notation and not a programming language, although it does have an opinionated
grammar. Consequently it has no built-in variables, no built-in operators and
even the reserved words are dynamically discovered during the parse.

The program has several phases, which we have completed:

- an initial ingestion phase in which the input is tokenised.
- a pass to discover and mark the identifiers that are used as keywords.
- a parsing of the tokens to form an internal AST.
- walking the AST to generate output.

We are now improving the solution and backfilling tests. Before we complete
on that, it would be very cool if we could implement a VSCode extension that
supported syntax highlighting for Monogram. To assist with this we can
provide a function that analyses a whole file and classifies every token and
records the start and end of the token.