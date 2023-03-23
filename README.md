# indent-js
This is a command-line tool written in Go that beautifies JavaScript code. It allows you to input JavaScript code from either a URL or a local file, and then formats it with consistent indentation and line breaks. The tool uses a simple approach based on counting braces to determine the appropriate level of indentation for each line. The resulting code is printed to the console.

`-url` : This flag is used to specify a URL to a JavaScript file that you want to beautify. You would use this flag if you want to fetch the JavaScript code from a remote location.

`-local`: This flag is used to specify a local file path to a JavaScript file that you want to beautify. You would use this flag if you want to use a JavaScript file that is stored on your computer.
