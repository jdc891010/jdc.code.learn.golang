


## Interfaces

**Defining one type of bot**
```golang
type englishBot struct
func (englishBot) getGreeting() string
```

**Defining another bot, focused on Spanish**
```golang
type spanishBot struct
func (spanishBot) getGreeting() string
```

**Defining interface associated with bots**
```golang
type bot interface{
    getGreeting() string
}
```

- Since both `englishBot` and `spanishBot` implement the `getGreeting()` method, the `bot` interface is satisfied.
- Therefor, they can be used to pass into func where the type is set to type bot, optimizing code re-use
- We know that, every value has a type
- Every function has to specify the type of its arguments

**Usage**
```golang
func printGreeting(b bot) {
  // do something here
}
```

- Now, both `spanishBot` or/and `englishBot` can be passed into `printGreeting()` func.

```golang
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

printGreeting(englishBot{})
printGreeting(spanishBot{})
```

- Thereby, single func to call the print greeting for each of the input types, which has `printGreeting` function already defined.
- So, no longer are we stuck with every function has to be re-written to accomodate different type even if the logic in it is identical?

**Example: io.Reader interface**
Background:
