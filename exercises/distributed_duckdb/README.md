# Can we challenge the idea of distributed compute, by using golang and duckdb as simple stack for concurrency and optimized olap


## How does Spark work in distributing data per cluster key?


1. I/O - load csv
2. Determine partitioning keys
3. Planning for task
4. Run per executor - using duckdb but sql
