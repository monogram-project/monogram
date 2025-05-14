# This is a script to set up a Debian development environment.

sudo apt update -qq

# Install the necessary dependencies and tools required for development.
sudo apt install -qq -y build-essential 

# Install podman for container management
sudo apt install -qq -y podman

# Install graphviz for rendering dot notation
sudo apt install -qq -y graphviz

# Install wkhtmltopdf for PDF generation
sudo apt install wkhtmltopdf

# Install current version of Python
sudo apt install -qq -y python3 python3-venv python3-dev
# Add poetry to the path
sudo apt install -qq -y python3-poetry
# Install pipx for Python package management
sudo apt install -qq -y pipx


