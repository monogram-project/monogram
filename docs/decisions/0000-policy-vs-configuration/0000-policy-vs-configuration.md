# 0000 - Policy vs configuration, 2025-02-16

## Issue

Monogram is intended to be a notation for writing program-structures analogous 
to the way that XML and JSON are notations for data-structures. It has no 
semantics but is easy to layer further structure and semantics on top. We
envisage applications such as configuration and domain-specific languages 
(DSLs).

Should it be a single notation, like JSON, or should it be extensible or
even configurable in the way XML is via entities? At one extreme it is a
single opinionated notation. In the middle it supports simple quality of
life options (trailing commas, newlines as terminators, predefined keywords).
And at the other end, it is straightforward to design a configuration-driven 
parser for instance.

### Tree-building is excluded

Note that this issue is confined to the surface syntax. The generation of the
syntax tree must be under programmer control for Monogram to be useful at all.

The `monogram` tool will simply take a `--format` option that tells it what the
output format should be and potential values might be `xml`, `json`, `yaml`,
`toml`, `gv` (graphviz) etc. The `monogram` library will allows programmers to
supply a tree-builder, that conforms to a pre-defined interface, and parsing the
input will drive the supplied tree-builder.

## Factors

- Flexibility
- Teachability
- Grammatical ambuiguity
- Where do we find the configuration/schema for a monogram file?
- Efficiency (but difficult to assess)


## Options and Outcome

- Single opinionated notation
- Configuration-driven parser
- Largely opinionated notation but with simple options giving extra flexibility

Outcome: *Single opinionated notation*

## Consequences

- Configuration-free
- Easy to use
- Best for teaching
- Defined syntax tree
- Least flexible

## Pros and Cons of Options

### Option 1: Single opinionated notation

Being opinionated means that Monogram is configuration-free, which leads to the
following benefits.

- Pros 
    - Easy to use.
    - Easy to teach.
    - Can be scanned/linted by third parties with zero knowledge.
    - No impact on build processes.
    - Defined syntax tree.

- Cons
  - Least flexible.
  - Compensating means using heuristics and tokenisation 'hacks'.

### Option 2: Configuration-driven parser

- Pros
    - Most flexible and widest potential flexibility.

- Cons
    - A mechanism to link monogram text to their schema (configuration) would be
      needed.
    - Most demanding on learners.
    - Designing the configuration notation so it is accessible is non-trivial.
    - Configuration notation would also need to specify how to build the tree.

- Interesting
    - Would this be competitive with other parser generation technologies?
    - Maybe a hybrid that locked down tokenisation but left the grammar rules
      open (or vice versa) would be a cool trade-off?

### Option 3: Largely opinionated notation but with simple options giving extra flexibility

- Pros
    - Allows the programmer to choose a sweet spot between flexibility and
      grammatical robustness.
    - Simple options keep it teachable and accessible.
    - Single syntax tree is still available.

- Cons
    - A mechanism to link monogram text to their options would be
      needed. Command-line options and environment variables would be bad
      ideas as then we have no ability to inspect and scan monogram.
    - There will never be enough options.

- Interesting
    - Macro notation is an interesting option here. Several languages show
      how this can be done: Lisp, C, Rust. We should definitely leave this
      door open.

