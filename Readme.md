## Learn Go with Testing

### If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway. go test -cover


### Compile time errors are our friends because they help us write software that works


### functions in Go are values.

### Table Driven Tests are useful when you want to build a list of test cases. They are a great fit when you wish to test various implementations of an interface, or if the data being passed in to a function has lots of different requirements that need testing.


### Note:- By convention you should keep your method receiver types the same for consistency.


#### Errors can be converted to a string with the .Error() method.


### An interesting property of maps is that you can modify them without passing as an address.A map value is a pointer to a runtime.hmap structure.


### The Buffer type from the bytes package implements the Writer interface, buffer := bytes.Buffer{}


### if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design.
