# Larnin' notes

Notes for larning. I know very little, so this is where I disambiguate terms etc I don't already understand while reading.

## Ch.1

"Goroutines are like threads..." ^1 Wait, what's a [thread](#thread)?

- <a name="thread"></a> Thread: In computer science, a thread of execution is the smallest sequence of programmed instructions that can be managed independently by a [scheduler](#scheduler). In most cases a thread is a component of a [process](#process). Multiple threads can exist within one process, executing concurrently and sharing resources such as memory, while different processes do not share these resources. In particular, the threads of a process share its executable code and the values of its variables at any given time. ^2

- <a name="scheduler"></a> Scheduler: a scheduler schedules- it assigns work specified by some means to resources that complete the work. 
- The work may be virtual computation elements such as threads, processes, etc, which are in turn scheduled onto hardware resources such as processors, network links or expansion cards.
- The process scheduler is a part of the operating system that decides which process runs at a certain point in time. It usually has the ability to pause a running process, move it to the back of the running queue and start a new process; such a scheduler is known as preemptive scheduler, otherwise it is a cooperative scheduler.
- Scheduling is fundamental to computation itself, and an intrinsic part of the execution model of a computer system; the concept of scheduling makes it possible to have computer multitasking with a single CPU. ^3


- <a name="process"></a> Process: basically a program in execution. Must progress in a sequential fashion. An entity which represents the basic unit of work to be implemented in the system. ^4 May be made up of multiple [threads](#thread) of execution. 
  - Consists of the following:
	- an image of the executable machine code associated with a program
	- memory (typically some region of virtual memory) that includes:
      - the executable code, process-specific data (input and output)
	  - a call stack (to keep track of active subroutines and/or other events)
	  - and a heap to hold intermediate computation data generated during run time.
    - Operating system descriptors of resources that are allocated to the process, such as file descriptors (Unix terminology) or handles (windows), and data sources and sinks.
    - Security attributes, such as the process owner and the process' set of permissions (allowable operations).
    - Processor state (context), such as the content of registers and physical memory addressing. The state is typically stored in computer registers when the process is executing, and in memory otherwise.
  - The operating system holds most of this information about active processes in data structures called process control blocks. Any subset of the resources, typically at least the processor state, may be associated with each of the process' threads in operating systems that support threads or child (daughter) processes. ^5
  
-------------------------

^1: Go in Action, 1.1.2

^2: <https://en.wikipedia.org/wiki/Thread_(computing)>

^5: <https://en.wikipedia.org/wiki/Scheduling_(computing)>

^4: <https://www.tutorialspoint.com/operating_system/os_processes.htm>

^5: <https://en.wikipedia.org/wiki/Process_(computing)#Representation>
