# GOLANG CLI OPERATIONS

## GOLANG CLI TOOL FOR DYNAMIC JSON PARSING DIGGING

I found not much resources on Dynamic JSON Parsing. Even if we use JSON.Unmarshall, it is a pain to access the elements of the interface(if you unmarshall into an empty interface). Moreover, Structs cant be used for Dynamic Parsing as we have no idea what the JSON file will look like. So, I wrote a digger in which we easily query the JSON interface not having to use reflect. This supports Nested JSON and Arrays. This is as long as the JSON file has no empty lines, the arrays are of string type and written in one line and integers are not supported yet. \[A limitation which can be improved!\]

1. This program takes JSON file to read in form of User input. It Provides a low level JSON Unmarshall and works on files Specifically formatted in a way currently. I am open to anyone who wants to contribute for this.
2. The Program Reads the JSON file line by line and appends the objects into map\[string\]interface{}. This is done using our own logic and I have NOT used json.Unmarshall for this. 
3. The Next Step is it lists out the parsed JSON in an indented fashion to give user an idea of the JSON file.
4. Next, the User is asked to enter a query and if the query is correct, the value is fetched and shown.


### Usage
Usage:  >>> jsondigger
        >>> ./check.json

## GOLANG CLI TOOL FOR JSON VALIDATION

1. This program should take a text file path as an argument and generate a JSON file using the same data or an error if the text file has incorrect json format.
2. All the text data is of string type.
3. We will be able to generate a “well-formatted” (indented) JSON file.
4. We will be using stacks for validation and some string validation
5. This program must have a “help” option, to know how to use this tool.

### Usage
Usage:  >>> jsonvalidator 

Options:
1. pretty <br>
            1 for Generating pretty JSON file and 0 for only validation
2. verbose <br>
            0 for no error description and 1 for error description

# Installation and Running
1. Clone this Project.
2. Make Sure Golang and Make are Installed and Open the root directory of the project.
3. Run `$ make install` in the CLI
4. Run `$ make digger`
5. Run `$ make validator` \[If you want to run jsonvalidator\]