package goroutines

import (
	creatingGoroutines "GoBasic/goroutines/creating-goroutines"
	fetchValuesFromGoroutines "GoBasic/goroutines/fetch-values-from-goroutines"
	waitingForGoroutinesToFinishExecution "GoBasic/goroutines/waiting-goroutines-to-finish-execution"
	playPauseExecutionOfGouroutine "GoBasic/goroutines/play-pause-execution-of-goroutine"
	fixRaceConditionUsingAtomicFunctions "GoBasic/goroutines/fix-race-condition-using-atomic-functions"
	defineCriticalSectionsUsingMutex "GoBasic/goroutines/define-critical-sections-using-mutex"

)

/*
Goroutines
Concurrency in Golang is the ability for functions to run independent of each other. 
Goroutines are functions that are run concurrently. Golang provides Goroutines as a way to handle operations concurrently.

New goroutines are created by the go statement.

To run a function as a goroutine, call that function prefixed with the go statement. Here is the example code block:

Example
sum()     // A normal function call that executes sum synchronously and waits for completing it
go sum()  // A goroutine that executes sum asynchronously and doesn't wait for completing it
The go keyword makes the function call to return immediately, while the function starts running in the 
background as a goroutine and the rest of the program continues its execution. The main function of every 
Golang program is started using a goroutine, so every Golang program runs at least one goroutine.

*/

func Goroutines() {
	creatingGoroutines.CreatingGoroutines()
	fetchValuesFromGoroutines.FetchValuesFromGoroutines()
	fixRaceConditionUsingAtomicFunctions.FixRaceConditionUsingAtomicFunctions()
	waitingForGoroutinesToFinishExecution.WaitingForGoroutinesToFinishExecution()
	playPauseExecutionOfGouroutine.PlayPauseExecutionOfGouroutine()
	defineCriticalSectionsUsingMutex.DefineCriticalSectionsUsingMutex()
}