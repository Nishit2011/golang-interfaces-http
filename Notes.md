Previously, we saw that our interface has a method tied to it with a return type and any other structs or other types that uses that method(getGreeting) gets part of the interface. But
the ReaderCloser interface uses Reader and Closer interfaces and not any method with a return type. (Check goland documentation for type ReaderCloser interface)

Why??
As you can see in the "different,sources, returns.png" screenshot, in order to work around with all the diffrent types, we would need a whole host of different functions that while they do exactly the same thing internally, we have to put together all these separate definitions of them because they all have to have these different types in their argument list exactly like we were talking about the problem that interfaces solve.

And that very problem is solved when we use the Reader interface.

type Reader interface {
    //remember request body is implementing the reader interface
    //so create this byte slice and pass it to the Read function and then inside of //the read function. That function takes the byte slice and its source of data
    //So the raw body of response is taken and injected inside or pushed into that //byte slice. 

    //int is the number of integers read into the byte slice
    
    Read(p []byte) (n int, err error)
}

So no matter what the different sources are, when we use reader interface, then we get some very common interface or some common point of contact that we can use to take that input and then pipe it off to different places inside of our codebase without writing a bunch of custom functions to work with each of these individual return types
