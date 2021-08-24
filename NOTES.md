

### Things To Fix

- passwords should be hashed
- service handlers can use CORS and auth via middlewares
- use pkg/errors to wrap errors and make error messages more informative and easier to identify in the code and debug
- either only use the name Todo or Task, don't mix the two words
- ValidateUser should return an error as well as a bool in case the db fails