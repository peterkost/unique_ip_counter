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

With this implemented my test file executed in around 97 seconds. I haven't
implemented proper RAM benchmarking, but I only saw it go to around 8GB from
just peaking at activity monitor. Big improvment over the naive solution.

We are using a fixed boolean array of size 4294967296. This is a lot of wasted
space if we are using a small input file, however the problem states that "the
file is unlimited in size" so this is a tradeoff worth making. At this time I
believe that this is the optimal space complexity for the problem.

## Five - Bytes instead of Strings

To get an idea of optimizations I can implement I had a read through Renato Pereira's
[article on his solution to the billion row challenge](https://r2p.dev/b/2024-03-18-1brc-go/).
I learned that using bytes instead of strings greatly improves the performance of the file scanner.
This seemed like a good first step since we want ints instead of strings anyways. With this single
optimization I was able to get the runtime of our sample file from 97 seconds to 54.

## Six - Scan file by bytes instead of lines

Going off of the last optimization I figured I might as well just read the file
byte by byte instaead of line by line to avoid multiple reads of the same line.
This optimization got me down to 44 seconds.
