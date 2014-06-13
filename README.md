# anomoly 

    əˈnäməlē/ 
    noun
    something that deviates from what is standard, normal, or expected.

anomaly formalizes an approach to dealing with errors in Go programs using the `panic` and `recover` mechanisms of Go langauge.  

It is minimally non-invasive; can be used in an opt-in fashion; and plays nicely with standard error and panics.

Use anamoly if you (a) beleive every error deserves attention; and (b) can't deal with the code noise that results from meticulous error checking.
    
## build and install

    cd <repo root>
    go build
    
# usage & guidline

Go functions indicating anomalous conditions typically return either an `error` or `bool`. In some cases e.g. `Reader.ReadLine` both `error` and `bool` types are used.

This is typical of call sites that can generate errors where the calling function has no means or intention of dealing with the error:

     result, ok, e := someFunc(...)
     if e != nil {
         return e
     }
     if !ok {
         return errors.New("someFunc returned false")
     }
     
With anomaly, we could (minimally) write:

     result, ok, e := someFunc(...)
     anomaly.PanicOnError(e)
     anomaly.PanicOnError(ok)

When using the error handling mechanism of anomaly, you have the option of retaining function signatures e.g. still return errors. You can also opt for treating them as panics that can be trapped at an appropriate top-level call site, typically the API boundary for libraries, or interface(s) to sub-systems.

When using anomaly to simply cleanup noise code in a function body, use the following pattern. (Note that use of named returned arg is an invasive requirement):

    func YourFunction(...) (err error) {
        // will trap all panics and return an error
        defer anomaly.Recover(&err)
       
        ...
       
        result, ok, e := someFunc(...)
        anomaly.PanicOnError(e)
        anomaly.PanicOnError(ok)
       
       ...
       
        result, e = someOtherFunc(...)
        anomaly.PanicOnError(e)
       
       ...
       
       // you can still return errors as normal
       e = Foobar()
       if e != nil {
          return e
       }
    }

If you omit the `defer anomaly.Recover(&err)`, the function will clearly throw panics in certain conditions. This is perfectly fine as long as you insure that at some higher level in the call chain you have a `anomaly.Recover(&err)` deferred block.

**recommendation**

Name functions and methods that use `anomaly.PanicOnError` and/or `anomaly.PanicOnFalse` with a leading underscore e.g. `_AssertNotNil` as a mnemonic for Go call pattern `res, _ := someFuncWithErr()`.

    
