# testing-handlers

You're probably at the point where you are writing test code for your handlers. This short example 
will covers how to "mock" the behavior of a data store (MySQLStore in your case). 

Brief breakdown of each file:
- `handler.go`: Contains two handlers that writes (Insert) and reads (Get) from a data store called `MyMockStore`.
- `handler_test.go`: Contains test code on how to test handlers given a context and a "mock" data store that 
similarly represents your actual data store.
- `user.go`: Contains a `User` struct used for this example.
- `myMockStore.go`: Contains two methods that "mocks" the behavior or your ACTUAL data store (MySQLStore). The two main 
behaviors covered are for successful and unsuccessful use of the methods given some input. Since the purpose of
testing your HANDLERS is to test its functionality, you can assume any dependencies behave as they should. Therefore,
we are allowed to "mock" the behavior of the data store to test whether our handlers conditional checks are implemented 
correctly.


Feel free to play around with the code by `forking` or `cloning` it down. You can even run the
tests as they should all pass. 