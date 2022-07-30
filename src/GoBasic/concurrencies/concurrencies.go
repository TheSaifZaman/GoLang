/*
Concurrency is an ability of a program to do multiple things at the same time. This means a program that have two or more tasks that run individually of each other, 
at about the same time, but remain part of the same program. Concurrency is very important in modern software, due to the need to execute independent pieces of code as
 fast as possible without disturbing the overall flow of the program.

Concurrency in Golang is the ability for functions to run independent of each other. A goroutine is a function that is capable of running concurrently with other functions. 
When you create a function as a goroutine, it has been treated as an independent unit of work that gets scheduled and then executed on an available logical processor. 
The Golang runtime scheduler has feature to manages all the goroutines that are created and need processor time. The scheduler binds operating system's threads to logical 
processors in order to execute the goroutines. By sitting on top of the operating system, scheduler controls everything related to which goroutines are running on which logical processors at any given time.

Popular programming languages such as Java and Python implement concurrency by using threads. Golang has built-in concurrency constructs: goroutines and channels. Concurrency 
in Golang is cheap and easy. Goroutines are cheap, lightweight threads. Channels, are the conduits that allow for communication between goroutines.

Communicating Sequential Processes, or CSP for short, is used to describe how systems that feature multiple concurrent models should interact with one another. 
It typically relies heavily on using channels as a medium for passing messages between two or more concurrent processes, and is the underlying mantra of Golang.

Goroutines — A goroutine is a function that runs independently of the function that started it.
Channels — A channel is a pipeline for sending and receiving data. Channels provide a way for one goroutine to send structured data to another.

Concurrency and parallelism comes into the picture when you are examining for multitasking and they are often used interchangeably, concurrent and parallel refer to related but different things.

Concurrency - Concurrency is about to handle numerous tasks at once. This means that you are working to manage numerous tasks done at once in a given period of time. However, you will only be doing 
a single task at a time. This tends to happen in programs where one task is waiting and the program determines to drive another task in the idle time. It is an aspect of the problem domain — where your 
program needs to handle numerous simultaneous events.

Parallelism - Parallelism is about doing lots of tasks at once. This means that even if we have two tasks, they are continuously working without any breaks in between them. 
It is an aspect of the solution domain — where you want to make your program faster by processing different portions of the problem in parallel.

A concurrent program has multiple logical threads of control. These threads may or may not run in parallel. A parallel program potentially runs more quickly than a sequential 
program by executing different parts of the computation simultaneously (in parallel). It may or may not have more than one logical thread of control.
*/

package concurrencies

func Concurrencies() {
	//ToDo

}