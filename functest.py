#!/usr/bin/env python3

"""
runtests.py - Functional Test Runner for the Monogram Command-Line Tool

Overview:
-----------
This script provides a functional test runner for the Monogram tool, or
command-line tools in general, which converts text on the standard input into
structured outputs (like XML or JSON). Test cases are defined in a YAML file
that specifies the command to run, the input data (fed via standard input), and
the expected output. Additionally, the YAML file (or individual tests) may
specify a normalization key (e.g., "xml", "json") to preprocess outputs before
comparison, helping to eliminate discrepancies due to formatting differences.

Usage:
-----------
Run the test runner from the command line by providing the YAML test data file
using the --tests option. For example:

    python3 runtests.py --tests tests.yaml

Security:
-----------
To avoid inadvertent execution of arbitrary binaries, this script checks that
the first token in the provided command resolves to an executable located within
a permitted directory. Use the --check-on-path option to specify the required
base directory for allowed executables. For example:

    python3 runtests.py --tests tests.yaml --check-on-path
    /path/to/allowed/directory

Normalization:
-----------
Normalization functions (such as for XML or JSON) can be specified either
globally in the YAML file (under "normalize") or overridden on a per-test basis.
This helps ensure that output differences in whitespace, attribute order, or key
order do not cause superficial test failures.

Output:
-----------
- Each test outputs a PASS or FAIL message.
- On failure, differences between the expected and actual output are displayed
  via a unified diff.
- A summary is printed at the end.
- Error messages and test failures are printed to stderr; normal pass messages
  are printed to stdout.
- The script exits with status code 0 if all tests pass, or 1 if any test fails.

"""

import argparse
import subprocess
import sys
import difflib
import yaml
import xml.dom.minidom
import json
import os
import shlex
import shutil
from pathlib import Path

def normalize_xml(xml_str):
    """
    Parse and pretty-print XML so that minor differences in whitespace,
    attribute order, etc., are eliminated.
    """
    try:
        dom = xml.dom.minidom.parseString(xml_str)
        normalized = dom.toprettyxml(indent="  ")
        # Remove the XML declaration and any blank lines.
        lines = normalized.splitlines()
        lines = [line for line in lines if line.strip() and not line.startswith("<?xml")]
        return "\n".join(lines)
    except Exception:
        return xml_str

def normalize_json(json_str):
    """
    Load and re-dump JSON so that differences in spacing or key order are normalized.
    """
    try:
        obj = json.loads(json_str)
        return json.dumps(obj, sort_keys=True, indent=2)
    except Exception:
        return json_str

# Mapping of normalization keys to functions.
normalization_functions = {
    "xml": normalize_xml,
    "json": normalize_json,
    # If an unrecognized value (or "none") is provided, no normalization is done.
}

def is_command_valid(command, base_path=None):
    """
    Check that the command (specifically, the first token—the executable)
    resolves to a path that lies inside the provided base_path.
    If base_path is None then the current working directory is used.

    Returns (True, "") if valid, otherwise (False, error_message).
    """
    tokens = shlex.split(command)
    if not tokens:
        return False, "Empty command."

    executable = tokens[0]
    resolved = None

    if os.path.isabs(executable):
        resolved = os.path.realpath(executable)
    else:
        resolved = shutil.which(executable)
        if resolved:
            resolved = os.path.realpath(resolved)

    if not resolved:
        return False, f"Could not resolve executable: {executable}"

    if base_path is None:
        base = Path(os.getcwd()).resolve()
    else:
        base = Path(base_path).resolve()

    exe_path = Path(resolved)
    try:
        if not exe_path.is_relative_to(base):
            return False, f"Executable {resolved} is not located within {base}"
    except AttributeError:
        common = os.path.commonpath([str(base), str(exe_path)])
        if common != str(base):
            return False, f"Executable {resolved} is not located within {base}"

    return True, ""

class Main:

    def __init__(self):
        parser = argparse.ArgumentParser(
            description="Functional test runner for the Monogram command-line tool."
        )
        # Renamed option --file to --tests.
        parser.add_argument(
            "--tests", 
            required=True, 
            help="YAML file containing test data"
        )
        parser.add_argument(
            "--check-on-path",
            help="Path under which the command's executable must reside. "
                 "If not provided, the current working directory is used."
        )
        parser.add_argument(
            "--command",
            default="monogram",
            help="Path to the executable to test."
        )
        self.args = parser.parse_args()

    def run_test(self, tcount, test, default_normalize=None, check_path=None):
        """
        Execute a single test case.
        The normalization setting is determined first by a test-specific flag
        and then falls back to the default. The check_path parameter specifies
        the allowed directory for the command's executable.
        """
        name = test.get("name", "<unnamed>")
        command = test["command"].format(command=Main().args.command).format(count=tcount)
        input_text = test.get("input", "")
        expected_output = test.get("expected_output", "")
        expected_exit_status = int(test.get("expected_exit_status", "0"))

        norm_key = test.get("normalize", default_normalize)
        normalization_func = normalization_functions.get(norm_key, None)

        valid, err_msg = is_command_valid(command, base_path=check_path)
        if not valid:
            return (name, False, f"COMMAND ERROR: {err_msg}", expected_output, "")

        result = subprocess.run(
            command,
            input=input_text,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            shell=True,
            text=True
        )
        actual_output = result.stdout

        if result.returncode != expected_exit_status:
            return (name, False, f"EXIT STATUS {result.returncode}", f"EXPECTED STATUS {expected_exit_status}", result.stderr)

        if normalization_func is not None:
            actual_output = normalization_func(actual_output)
            expected_output = normalization_func(expected_output)

        passed = actual_output.strip() == expected_output.strip()
        return name, passed, actual_output, expected_output, result.stderr

    def run_single_test(self, tcount, test, default_normalize):
        """
        Run a single test and print its result.
        If the test fails or there is a command error, warning messages are sent to stderr.
        Returns True if the test passed, else False.
        """
        name, passed, actual, expected, stderr_text = self.run_test(
            tcount,
            test,
            default_normalize=default_normalize,
            check_path=self.args.check_on_path
        )
        if passed:
            print(f"PASS: {name}")
        else:
            print(f"FAIL: {name}")
            print("Expected:")
            print(expected)
            print("Actual:")
            print(actual)
            diff = difflib.unified_diff(
                expected.splitlines(),
                actual.splitlines(),
                fromfile="expected",
                tofile="actual",
                lineterm=""
            )
            print("Diff:")
            print("\n".join(diff))
            if stderr_text:
                print("Error output:")
                print(stderr_text)
            print("-" * 40)
        return passed
    
    def main(self):
        try:
            with open(self.args.tests, "r", encoding="utf-8") as f:
                data = yaml.safe_load(f)
        except Exception as e:
            print(f"Error reading {self.args.tests}: {e}", file=sys.stderr)
            sys.exit(1)

        default_normalize = data.get("normalize", None)
        tests = data.get("tests", [])
        if not tests:
            print("No tests found in the YAML file!", file=sys.stderr)
            sys.exit(1)

        total = 0
        passed_count = 0

        for tcount, test in enumerate(tests):
            total += 1
            if self.run_single_test(tcount, test, default_normalize):
                passed_count += 1

        print(f"\nSummary: {passed_count} out of {total} tests passed.")
        sys.exit(0 if passed_count == total else 1)

if __name__ == "__main__":
    Main().main()
