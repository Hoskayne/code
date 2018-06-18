# Information

Notes for larning. I know very little, so this is where I disambiguate terms etc I don't already understand while reading. I quote everything that's from the book directly

## Ch.1

### 1.1.2

>Goroutines are like threads..." 

Wait, what's a [thread](#thread)?

>Goroutines are functions that run concurrently with other goroutines, including the entry point of your program. In other languages, you’d use threads to accomplish the same thing, but in Go many goroutines execute on a single thread...Goroutines use less memory than threads and the Go runtime will automatically schedule the execution of goroutines against a set of configured logical processors. Each logical processor is bound to a single OS thread.

![Figure 1.2. Many goroutines execute on a single OS thread](https://i.imgur.com/LVdhpb6.jpg)

>Channels are data structures that enable safe data communication between goroutines.

Basically, the channels make sure that one goroutine doesn't do anything to data before other goroutines are ready.

Example:

![Figure 1.3. Using channels to safely pass data between goroutines](https://i.imgur.com/MmbTi0V.jpg)

>In figure 1.3 you see three goroutines and two unbuffered channels. The first goroutine passes a data value through the channel to a second goroutine that’s already waiting. The exchange of the data between both goroutines is synchronized, and once the hand-off occurs, both goroutines know the exchange took place. After the second goroutine performs its tasks with the data, it then sends the data to a third goroutine that’s waiting. That exchange is also synchronized, and both goroutines can have guarantees the exchange has been made...It’s important to note that channels don’t provide data access protection between goroutines. If copies of data are exchanged through a channel, then each goroutine has its own copy and can make any changes to that data safely. When pointers to the data are being exchanged, each goroutine still needs to be synchronized if reads and writes will be performed by the different goroutines.

### 1.1.3

>Go provides a flexible hierarchy-free type system that enables code reuse with minimal refactoring overhead... Go developers simply embed types to reuse functionality in a design pattern called composition... In Go, types are composed of smaller types, which is in contrast to traditional inheritance-based models.

![Figure 1.4. Inheritance versus composition: Go interfaces model small behaviors](https://i.imgur.com/JTwB86O.jpg)

[Relevant gist](https://gist.github.com/rsperl/01d7a4da7c01706d50f6)

>In Go, if your type implements the methods of an interface, a value of your type can be stored in a value of that interface type. No special declarations are required...In a strictly object-oriented language like Java, interfaces are all-encompassing. You’re often required to think through a large inheritance chain before you’re able to even start writing code.

E.g., in java, implementing an interface requires you to create a class that fulfills all of the promises made in that interface, and explicitly declare that you are implementing it.

>In contrast, a Go interface typically represents just a single action.

The book then uses `io.Reader` as an example. From the [documentation](https://golang.org/pkg/io/#Reader):
>Reader is the interface that wraps the basic Read method.

>Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered. 


[Relevant Medium article](https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b) says the following about `io.Reader`:

>If you’re designing a package or utility (even if it’s an internal thing that nobody will ever see) rather than taking in strings or []byte slices, consider taking in an io.Reader if you can for data sources. Because suddenly, your code will work with every type that implements io.Reader.

See [this](chapter1/toy-readers/first.go) short example, and [this one](chapter1/toy-readers/second.go) for a slightly different way to do the same thing.

Huh!

### 1.1.4

Go has a nice garbage collector hooray

### Extra Larning

<a name="thread"></a> **Thread**: In computer science, a thread of execution is the smallest sequence of programmed instructions that can be managed independently by a [scheduler](#scheduler). In most cases a thread is a component of a [process](#process). Multiple threads can exist within one process, executing concurrently and sharing resources such as memory, while different processes do not share these resources. In particular, the threads of a process share its executable code and the values of its variables at any given time. ^1

<a name="scheduler"></a> **Scheduler**: a scheduler schedules- it assigns work specified by some means to resources that complete the work.

+ The work may be virtual computation elements such as threads, processes, etc, which are in turn scheduled onto hardware resources such as processors, network links or expansion cards.
+ The process scheduler is a part of the operating system that decides which process runs at a certain point in time. It usually has the ability to pause a running process, move it to the back of the running queue and start a new process; such a scheduler is known as preemptive scheduler, otherwise it is a cooperative scheduler.
+ Scheduling is fundamental to computation itself, and an intrinsic part of the execution model of a computer system; the concept of scheduling makes it possible to have computer multitasking with a single CPU. ^2


<a name="process"></a> **Process**: basically a program in execution. Must progress in a sequential fashion. An entity which represents the basic unit of work to be implemented in the system. ^3 May be made up of multiple [threads](#thread) of execution. 

  + Consists of the following:
	+ an image of the executable machine code associated with a program
	+ memory (typically some region of virtual memory) that includes:
      + the executable code, process-specific data (input and output)
	  + a call stack (to keep track of active subroutines and/or other events)
	  + and a heap to hold intermediate computation data generated during run time.
    + Operating system descriptors of resources that are allocated to the process, such as file descriptors (Unix terminology) or handles (windows), and data sources and sinks.
    + Security attributes, such as the process owner and the process' set of permissions (allowable operations).
    + Processor state (context), such as the content of registers and physical memory addressing. The state is typically stored in computer registers when the process is executing, and in memory otherwise.
  + The operating system holds most of this information about active processes in data structures called process control blocks. Any subset of the resources, typically at least the processor state, may be associated with each of the process' threads in operating systems that support threads or child (daughter) processes. ^4
  
## Ch. 2
  
-------------------------

^1: <https://en.wikipedia.org/wiki/Thread_(computing)>

^2: <https://en.wikipedia.org/wiki/Scheduling_(computing)>

^3: <https://www.tutorialspoint.com/operating_system/os_processes.htm>

^4: <https://en.wikipedia.org/wiki/Process_(computing)#Representation>
