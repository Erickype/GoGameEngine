# GoGameEngine
This project will be a game engine

## Things that needs research
- Game teams
- Game engine parts

## Things in research

## Things learned
- Make a game only using the defaul Go
- Make a small package to build games

## Reading history
- Runtime engine architecture
- Tools and asset pipeline
- Parallelism and concurrency

# Things done

## Getting started
Use the go get command to use the Game Engine

```
go get github.com/Erickype/GoGameEngine
```

## Entry point
- The basic entry point is in the Core package, 
that creates and run the application. Use the command
`Core.CreateApplication()` after importing the 
package.

```
Core.CreateApplication()
```

## Log System
- The package Log contains the methods and definitions
for the log system. At this time it is only used
for the `Core.CreateApplication()` method.
- It will print info in the client and the engine.
