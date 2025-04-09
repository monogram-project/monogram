[private]
default:
    just --list

# Update the poetry environments, run the first time after cloning the repo.
setup:
    cd functests && poetry update
    cd make_examples && poetry update
    cd make_railroad_diagram && poetry update

# Initialize decision records
init-decisions:
    python3 scripts/decisions.py --init

# Add a new decision record
add-decision TOPIC:
    python3 scripts/decisions.py --add "{{TOPIC}}"

jumpstart:
    sh jumpstart.sh

install: build
    just -f go/monogram/Justfile install

build:
    just -f go/monogram/Justfile build

test: unittest functest

functest:
    cd functests && poetry run python3 functest.py --tests tests.yaml --command "../go/monogram/monogram"

functest-coverage:
    just -f go/monogram/Justfile functest-coverage

unittest:
    just -f go/monogram/Justfile unittest

unittest-coverage:
    just -f go/monogram/Justfile unittest-coverage

