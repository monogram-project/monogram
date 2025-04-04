# Change Log for Monogram

Following the style in https://keepachangelog.com/en/1.0.0/

## [0.2.3] Fixes span info

### Added 

- CHANGELOG.md

### Fixed

- Span info for "unit" nodes
- Span info for chained infix operators


## [0.2.2.0] Fixes cosing triple quotes

### Fixes

- Closing triple quotes do not have to be on their own but may be followed by other text. e,g, for infix operators.
- Test recipes working.


## [0.2.1.1] Adds string interpolation and multi-line

### Added

- New `just` recipes `unittest` and `functest` for running the types types of tests.
- String interpolation
- Multi-line strings


## [0.2.0.5] Example implementations: calc and typecalc

### Added

- New --version option, integrated with GH workflow
- Example application calc: using dynamic 'monotype' AST
- Example application typecalc: using typed AST
- More functional tests
- Working GH Codespace
- More examples added to docs
- Status 'badges' added to the README.md

### Fixes

- Subtraction handled correctly in expressions such as `3-2`


## [0.2.0] First version of reference implementation in Golang

## [0.1] Proof of concept implemented in Pop-11


