1. This program should take a text file path as an argument and generate a JSON file using the same data or an error if the text file has incorrect json format.
2. All the text data is of string type.
3. We will be able to generate a “well-formatted” (indented) JSON file.
4. We will be using stacks for validation and some string validation
5. This program must have a “help” option, to know how to use this tool.

# Usage
Usage: ./jsonvalidate [options] <textfile>
Options:
  -pretty
        1 for Generating pretty JSON file and 0 for only validation
  -verbose
        0 for no error description and 1 for error description
