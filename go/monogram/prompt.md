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

We are now working on a github release pipeline.  The current workflow is this. 

```yaml
name: Update version.go on Tag Push

on:
    push:
      tags:
        - "v*"      # Matches tags starting with "v"

jobs:
  update-version:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Repository
        uses: actions/checkout@v3
        with:
          fetch-depth: 0 # Ensure full history is available, including tags

      - name: Check Out Main Branch
        run: |
          git fetch origin main
          git checkout main

      - name: Extract Git Tag
        id: get_tag
        run: |
          echo "TAG=${GITHUB_REF#refs/tags/}"
          echo "TAG=${GITHUB_REF#refs/tags/}" >> $GITHUB_ENV

      - name: Update version.go
        run: |
          echo "package lib" > go/monogram/lib/version.go
          echo "const Version = \"${TAG}\"" >> go/monogram/lib/version.go

      - name: Commit Changes
        run: |
          git config --global user.name "github-actions[bot]"
          git config --global user.email "actions@github.com"
          git add go/monogram/lib/version.go
          git commit -m "Update version.go to ${TAG}"
          git push origin main
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

We want to extend this pipeline so we automatically publish a release.