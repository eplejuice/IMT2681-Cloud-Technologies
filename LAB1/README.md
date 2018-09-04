#Programming Tasks


### 1. Write a function that returns nth element of the Fibonacci sequence. Write version with recursion, and version with a loop.

### 2. Write a function that returns the factorial of a given number. Write version with recursion, and version with a loop. 

### 3. Put those two functions into their own module, and call them from a different module.

### 4. Write a function that takes zero arguments, and returns a function that takes two integers, and returns their sum. Do that using anonymous functions inline.

### 5. Create a struct type representing a Person. Create a method that makes the person older by 10 years, ie. function that takes person as an argument. Do that with a struct method, and as a standalone function. What are the differences in using the instance by value, or by reference? (Use pointers).

### 6. Write a simple web server, that can handle requests to URLs of the form "/hello/" and responds with string "Hello " where name is the parameter that is provided to the request. Use string tokenizer to extract the last element of the path.

### 7. Modify the program from point 6 in such a way, that it gracefully handles all possible errors: e.g. calls to
        /
        /non-existent-path
        /hello (without the trailing slash followed by "name"
        /hello/123  (number instead of a name)


### 8. Create a in-memory database, that stores books in an array. Books have id, year, title and author. Make a simple service that allows you to retrieve a book by the given id, through API like: /books/  Make sure you return appropriate return code for case where books with given id does not exist.
