/*
Illustration of Checkpoint Synchronization
The checkpoint synchronization is a problem of synchronizing multiple tasks. Consider a workshop 
where several workers assembling details of some mechanism. When each of them completes his work, they put the details together. 
There is no store, so a worker who finished its part first must wait for others before starting another one. 
Putting details together is the checkpoint at which tasks synchronize themselves before going their paths apart.
*/

package checkpointSynchronizationProblem

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func worker(part string) {
	log.Println(part, "worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part")
	wg.Done()
}

var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
	wg          sync.WaitGroup
)

func CheckpointSynchronizationProblem() {
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c)
		wg.Add(len(partList))
		for _, part := range partList {
			go worker(part)
		}
		wg.Wait()
		log.Println("assemble.  cycle", c, "complete")
	}
}

/*
Output
2019/07/15 16:10:32 begin assembly cycle 1
2019/07/15 16:10:32 D worker begins part
2019/07/15 16:10:32 A worker begins part
2019/07/15 16:10:32 B worker begins part
........
2019/07/15 16:10:32 D worker completes part
2019/07/15 16:10:32 C worker completes part
2019/07/15 16:10:32 assemble.  cycle 3 complete
*/