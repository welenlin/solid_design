# SOLID Design Principles by using Shape as Example in Go

According to Wikipedia, SOLID is a mnemonic acronym for five design principles intended to make software designs more understandable, flexible and maintainable. It is not related to the GRASP software design principles.

In this project, I'm using Shape class as an example to structure the project by using SOLID design principles.

### Single Responsibility Principle

A class should have one, and only one, reason to change.

### Open-closed Principle

Entities should be open for extension, but closed for modification.

### Liskov Substitution Principle
Let ϕ(x) be a property provable about objects x of type T.
Then ϕ(y) should be true for objects y of type S, where S is a subtype of T.

### Interface Segregation Principle
A Client should not be forced to implement an interface that it doesn’t use.

### Dependency Inversion Principle
High-level modules should not depend on low-level modules. Both should depend on abstractions.
Abstractions should not depend on details. Details should depend on abstractions.
