## Concurrency in Go
- Breaking up a single process into indenpedent components and specifying how these components safely share data
- Most concurrency models, OSlvel thread share data by attempting to acquire locks
- Golang model is different, referred to *CSP* (**C**ommunicating **S**equencing **P**rocesses) -- Tony Hoare [1978]
- 

### When to use Concurrency?

1.
