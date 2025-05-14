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

