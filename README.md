# Lightspeed Take Home Assessment

https://github.com/Ecwid/new-job/blob/master/IP-Addr-Counter-GO.md

# Summary

Time complexity: O(n)
Space complexity: O(1)

Runtime on M1 MacBook Pro with 16 GB of RAM is about 60 seconds.
Number of unique IPs in the test file 1,000,020,936 (1 billion).

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

## Seven - Adding concurrency

I had heard that Go makes concurrency really easy and now I finally get to try
it. Full disclosure, my implementation was inspired by Pereria's article
mentioned above. Before impelmenting concurrency I was keeping track of the
total number of IPs and the number of IPs I had already seen. With concurrency
my use of an array would create race conditions. To avoid this I opted to simply
set each IP to seen and at the end loop over the array once and tally up number
of seen IPs. This caused my unit tests to go from around a second to seven,
since the sample size was only ten IPs. However as mentioned with the memory
usage, my implementation is aiming for handling an unlimited number of IPs as
stated in the problem in which case the tradeoff of this loop is well worth the
speed increase of concurency. With this implemented the runtime of my sample
file went from 44 seconds to around 19 seconds, but note that now my test cases
are responsible for about 7 of those so the gains on the actual input file are
quite substantial.

I ran the full file and got a runtime of a bit over six minutes, but I think the
bottle neck here is the IO. I'm running the full file off a USB SSD and the 10%
file off the internal SSD. THe 10% file runs in 14 seconds so six minutes is
quite a bit more. Should take about half that, but I don't have the storage to
test the full file off of the internal.

## Eight - Final results and thoughts

I decided to transfer the file over to my internal SSD and got a runtime of
around 60 seconds. That seems like a respectable time so I'll leave the
optimization there.
