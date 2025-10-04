# 97 - Go Extension

It provides a comprehensive set of features that include syntax checking, linting, and error reporting that align closely with the go compiler and tools.

Key features of the Go extention for VS code include:

- Syntax Highlighting
  The purpose of Syntax Highlighting is to help in visually differentiating parts of the code, such as keywords, variables, and types, which aids in readability and spotting Syntax errors.

- We have error reporting and diagonistics through the Go extention.
  The purpose of error reporting and diagonistics is to provide real time feedback on errors and potential issues in the ocde before running it.

- Next feature is Linting, Linting identifies potential issues that may not be strictly errors, but could lead to problems or are against go conventions. And that's how we get warnings. The warnings are also squiggly lines, but they are highlighted in yellow and they are something that we can avoid and still run our code if we have an error in our code, then it will definitely generate an error and the code will not compile. But warnings is something that can be avoided, that can be discarded, and the code will still run. So, we can run the code even if we have warning message in our code base.

- Another feature is that code formatting. Code formatting automatically formats the code to follow Go's formatting conventions, even if we do not have a leading space.

- Another feature is that code navigation, it provides features like go to definition, find references and code suggestions.

So, what go vs code extention does is that we get a go compiler before we run the compiler, because go vs code extention analyzes our code based on the instructions that are fed to the go compiler.
However, yes, there is certain errors that are only caught when we run the code, But yes, there are many errors that we do avoid even running the code.

The go extention does certain a lot of rules that the go compiler follows and generates error in our code based on those rules.