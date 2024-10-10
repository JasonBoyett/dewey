# Testing Philosophy

The test suite will create mock file system to test against.
Because of the nature of the application the tests must be run against
mocks and side effects. The best way we could come up with to test the
application was to use a mock file system. The setup function exports a struct
```TestInfo``` This struct is returned from the ```setup()``` function and has
all the nescisary feilds to create tests around the mock file system.
This includes a map with the paths to all test files and their expected contents
allong with a method to open all test files and return a map of the active test
files and their expected content.
