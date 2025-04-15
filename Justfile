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

build-for-docker:
    just -f go/monogram/Justfile build-for-docker

test: unittest functest

functest:
    cd functests && poetry run python3 functest.py --tests tests.yaml --command "../go/monogram/monogram"

functest-coverage:
    just -f go/monogram/Justfile functest-coverage

unittest:
    just -f go/monogram/Justfile unittest

unittest-coverage:
    just -f go/monogram/Justfile unittest-coverage

# Run the latest version of monogram and print the version
get-version:
    (cd go/monogram && go run ./cmd/monogram --version)

# Bump the version of monogram and optionally commit and push
bump BUMP:
    python3 ./scripts/bump.py --bump={{BUMP}} --save 

# Publish a new release
publish:
    python3 ./scripts/bump.py --publish
