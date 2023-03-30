# GoGameEngine
This project will be a game engine written in the Go
language. Will implement a graphics agnostic library to use
DirectX, OpenGl or Vulkan.

# Things already implemented

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

## Event System
- The package Event contains all the files and
methods related to the event system.
- The event systems uses an interface `IEvent` to define
the common methods to all type of events.
- There are two enums containing the types of events `EventType`
and categories `Category`
- The concrete events are implemented in witch one of the
categories files. Those event inherits the `Event` struct
and creates its own methods.
- There is an `EventDispatcher` whose job is to dispatch an
event that implements `IEvent` interface to the subsystems
that manage those events. (Still need to implement those subsystems)
- There is `Factory` that can create any concrete event.

### Manager
- The `EventManager` struct uses the `EventDispatcher` and `Factory`
to manage the events.
- Create a new instance of `EventManager` with:
````
eventManager := Events.EventManager{}
````
- Create an event with `CreateEvent(Events."EventType"")` method
passing the event type to be created.
````
event := eventManager.CreateEvent(Events.WindowResize)
````
- Dispatch that event with `Dispatch("event")` passing the
created event
````
eventManager.Dispatch(event)
````

### Notice
Develop of a game engine is really complicated and even more
if you try to use a high or mid level language like go.
There are the bindings from de c++ language but develop becomes
really hard due to lack of documentation, hard in the debug process
and the differences of the languages.
So I will stop developing of this and focus in other projects 
that are more suitable to the Go language, however I've learned a lot of things 
like: the architecture of a game engine, the event system, platforms, renderers,
layers, ImGui, entry point, to read c++ code and some particularities of the Go
language. I really enjoyed working on this. 