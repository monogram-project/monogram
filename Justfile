[private]
default:
    just --list

# Initialize decision records
init-decisions:
    python3 scripts/decisions.py --init

# Add a new decision record
add-decision TOPIC:
    python3 scripts/decisions.py --add "{{TOPIC}}"

jumpstart:
    go install github.com/wadey/gocovmerge@latest

install: build
    just -f go/monogram/Justfile install

build:
    just -f go/monogram/Justfile build

functest:
    python3 runtests.py --tests tests.yaml --command "./go/monogram/monogram"

functest-coverage:
    just -f go/monogram/Justfile functest-coverage

unittest:
    just -f go/monogram/Justfile unittest

unittest-coverage:
    just -f go/monogram/Justfile unittest-coverage

