# Change Log for Monogram

Following the style in https://keepachangelog.com/en/1.0.0/

## [0.2.7] Bug Fixes

### Fixes

- Inconsistent phrasing in the [README.md](README.md) has been fixed.

- A leading backslash is supposed to introduce an identifer. However this 
  was actually not being recognised as valid syntax. This is now fixed.

```
❯ echo '\_x' | monogram -f xml 
<unit>
  <identifier name="x" />
</unit>
```

- Statements were simple sequences of expressions. This was not intentional
  and the documented grammar is now enforced. This means that the statements 
  of a form are now separated by semi-colons or line breaks.

```
❯ echo 'block a = b * b; x = y + z endblock' | monogram -f xml 
<unit>
  <form syntax="surround">
    <part keyword="block">
      <operator name="=" syntax="infix">
        <identifier name="a" />
        <operator name="*" syntax="infix">
          <identifier name="b" />
          <identifier name="b" />
        </operator>
      </operator>
      <operator name="=" syntax="infix">
        <identifier name="x" />
        <operator name="+" syntax="infix">
          <identifier name="y" />
          <identifier name="z" />
        </operator>
      </operator>
    </part>
  </form>
</unit>

```

Or equivalently

```
❯ cat << "EOF" | ./monogram -f xml
block
    a = b * b
    x = y + z
endblock
EOF
<unit>
  <form syntax="surround">
    ETC ... deleted for brevity. Identical to above.
  </form>
</unit>

```

## [0.2.6] Docker image and Underscores in Numbers

### Added

- Underscores in numbers now supported

- Prebuilt binaries on release for popular platforms:
    - Linux on x86_64 and Arm64
    - MacOS on x86_64 and Arm64
    - Windows on Intel
- Docker image sfkleach/monogram pushed on release to Docker Hub
    - Use with `docker run --rm -i sfkleach/monogram [OPTIONS] < STDIN > STDOUT`

## [0.2.4] Fixes span info

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


