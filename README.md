# Lightspeed Take Home Assessment

https://github.com/Ecwid/new-job/blob/master/IP-Addr-Counter-GO.md

# Dev Log

## One - Scanning the input to count the number of lines

The first challenege to takle is reading in the massive input file. With a naive
file scanner it took about thirteen minutes for the script to count up the number of
lines. The result was eight billion.

## Two - Adding Benchmarking

When doing any sort of performance optimization it is essential to benchmark
your results to ensure that what your optimizations are actually having the
result you expect. I created a basic benchmark that uses 10% of the full input
file (800 million lines). The resulting runtime is around 16 seconds.

## Three - Naive solution with map

This solution passed my test cases with ten ips, but our test file with 8MM
lines timed out for taking more then eleven minutes in the benchmark. I'm
guessing this has something to do with Mac OS's memory managment. At one point I
saw the benchmark using over 25GB of RAM. I only have 16GB so I'm assuming it
writes some of it to disk and then when we try to access it the OS need's to
move it back into memory.

## Four - Use array instead of map

Currently each IP uses a string and an int in memory. This wastes a lot of
memory since all we care about for each IP is if we've seen it before. Go's
smallest type `bool` which only takes a single byte will provide us all of the
information we need for a given IP address. We are also only working with IPv4
addresses which is limited to a bit over 4 billion. Therfore to store all of the
information we need in memory we will a bit over 4 GB.
