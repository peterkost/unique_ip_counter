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
