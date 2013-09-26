# gorelic

New Relic agent for Go lang

Development started again from scratch



### Metrics reported by plugin
1. Runtime/General/NOGoroutines - number of runned go routines, as it reported by NumGoroutine() from runtime package
2. Runtime/General/NOCgoCalls - number of runned cgo calls, as it reported by NumCgoCall() from runtime package

3. Runtime/GC/NumberOfGCCalls - Nuber of GC calls, as it reported by ReadGCStats() from runtime/debug 
4. Runtime/GC/PauseTotalTime - Total pause time diring GC calls, as it reported by ReadGCStats() from runtime/debug (in nanoseconds)

5. Runtime/GC/GCTime/Max - max GC time
6. Runtime/GC/GCTime/Min - min GC time
7. Runtime/GC/GCTime/Mean - GC mean time
8. Runtime/GC/GCTime/Percentile95 - 95% percentile of GC time

Metrics 5-8 are measured in nanoseconds, and they can be inaccurate if GC called more often then once in GC_POLL_INTERVAL_IN_SECONDS. 
If in your workload GC is called more often - you can consider decreasing value of this constant. But ReadGCStats() blocks garbage collection, so its not good idea to call it very often.

TODO:
Alloc      uint64 // bytes allocated and still in use
Sys        uint64 // bytes obtained from system

Lookups    uint64 // number of pointer lookups

Mallocs    uint64 // number of mallocs
Frees      uint64 // number of frees

