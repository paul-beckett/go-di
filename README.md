# Dependency Injection in Go
The aim of this project is to provide a few samples of how application code can be improved through DI (Dependency Injection) and IoC (Inversion of Control) concepts.

Something to be mindful of is Go semantics in passing copy of a value or the pointer. It can be easy to get undesired behaviour if not careful.

# Examples
All of the examples follow the same pattern. A small function chain (Function1-3) that returns the result from a "DBClient".

"DBClient" is used as a representative example. It could easily be substituted for any other functionality.

## example1
In this example the DBClient is instantiated within Function3.

As a result, DBClient is hardcoded and the caller of the function has no control over what DBClient is or how it behaves.

Code like this becomes hard to test and/or change. We can't test Function1-3 without also testing DBClient.

## example2
This example improves on example 1 by injecting the DBClient into the functions.

The immediate improvement can be seen in the test. As we now have control over the DBClient we can test its behaviour with different inputs.

We still have the undesired behaviour of testing DBClient. Also notice how each of our function calls have to declare DBClient, even though they may not use it other than pass through. What would happen if we needed to add another dependency? Having to change multiple functions like this isn't great and eventually the function parameters lists are going to get unwieldy.

Note: if you do not have control over the code to modify existing packages as seen in example3-4 then this may be your best option.

## example3
This example improves on example 2 by using constructor based injection.

We've introduced a Service struct containing the DBClient dependency and refactored the functions to be methods on Service.

Notice how Function1 and Function2 now have no reference to DBClient. That should limit the scope of any changes when adding or changing any dependency.

At this point we are still testing DBClient.

## example4
This example improves on example 3 by allowing DBClient to be mocked.

It is usually a good practice to code to interfaces instead of concrete implementations. In examples 1-3 DBClient is a concrete implementation and therefore difficult to change on the fly.
Note: the concepts added in example4 could also be added to example2.

The main change here is to change DBClient to be an interface defining the public functionality, and introduce a concrete implementation called BasicDBClient. Note that on Service we no longer need to declare the client property as a pointer.

Now in the tests we can create a mock instance, thus allowing us to isolate our tests on the behaviour of the Service. I've used testify/mock for the example but any mocking functionality could have been used.

In real code the Service should probably also be an interface with separate concrete implementation but I've kept that out to keep the changes from example3-4 a little bit simpler.

## End result
Whilst some of the benefits these improvements make has been described in terms of testing, its far from the only benefit.

Comparing the code from example1 to example4 it should be clear to see that the code in example4 is much easier to maintain, change, and also understand. Combined with good package structuring, particularly in a larger or more complex app, the improvements will be vast.

# Real code
The examples demonstrate the benefits of DI, but how should all of the dependencies be created and injected?

There are a number of DI frameworks out there but its often easy to construct your own app without the added complexity of a DI framework.

The process of injecting dependencies is often referred to as wiring. e.g. a Service may depend on a DBClient, therefore it can be said that the DBClient is wired into the Service. Typically the wiring process is handled by the applications main function. In a fictional app that may look something like the following:

```
func main(){
	//create and wire the dependencies
	dbClient := db.NewDyanmoDBClient("<connection string goes here>")
	restClient := rest.NewRestClient("<connection string goes here>")
	userService := user.NewUserService(dbClient)
	orderService := order.NewOrderService(dbClient, restClient)
	
	server := server.NewServer(userService, orderService)
	
	http.listenAndServe("8080", server)	
}
```