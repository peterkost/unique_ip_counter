# Reading 8 Billion Lines in Go - From 13 minutes to 58 seconds

This is my solution to the take home assessment for a SWE role at Lightspeed. The problem statement can be found in [this file](IP-Addr-Counter-GO.md) or in [their repo](https://github.com/Ecwid/new-job/blob/master/IP-Addr-Counter-GO.md).

## Summary

Time complexity: O(n)

Space complexity: O(1)

Runtime on M1 MacBook Pro with 16 GB of RAM is about 60 seconds.

Number of unique IPs in the test file 1,000,000,000 (1 billion).

## Execution

```bash
go run uniqueip.go -f ip_addresses
```

## Dev Log

### 01 - Scanning the input to count the number of lines

The first challenge is efficiently reading the massive input file. Initially, using a naive file scanner, it took about thirteen minutes for the script to count the number of lines, which totalled eight billion.

### 02 - Adding Benchmarking

When performing any kind of optimization, it's essential to benchmark your results to ensure that your changes are actually having the desired effect. I created a basic benchmark using 10% of the full input file (800 million lines), which resulted in a runtime of around 16 seconds. The runtime was further improved by running the benchmark file from my internal SSD.

### 03 - Naive solution with map

This solution passed my test cases with ten IPs, but when I ran the test file with 8 million lines, it timed out after taking more than eleven minutes in the benchmark. I suspect this issue is related to macOS's memory management. At one point, I noticed the benchmark using over 25GB of RAM. Since my machine only has 16GB of RAM, I assume the OS is writing some of the data to disk, and when we try to access it, the OS needs to move it back into memory.

### 04 - Use array instead of map

Currently, each IP address is stored as a string and an integer in memory, which is wasteful since all we need to know is whether we've seen the IP before. Go's smallest type, bool, which takes only a single byte, provides all the information we need for each IP address. Since we're only working with IPv4 addresses, which are limited to just over 4 billion, we can store all the necessary information in memory using a bit over 4GB.

After implementing this change, my test file executed in around 97 seconds. Although I haven't set up proper RAM benchmarking, I observed the memory usage peaking at around 8GB according to the Activity Monitor—a significant improvement over the naive solution.

We are using a fixed boolean array of size 4,294,967,296. This approach does result in some wasted space for smaller input files, but since the problem statement mentions that "the file is unlimited in size," this tradeoff is justified. At this point, I believe this is the optimal space complexity for the problem.

### 05 - Bytes instead of Strings

To explore potential optimizations, I read through Renato Pereira's [article on his solution to the billion row challenge](https://r2p.dev/b/2024-03-18-1brc-go/). I learned that using bytes instead of strings significantly improves the performance of the file scanner. This seemed like a good first step, especially since we ultimately want integers rather than strings. By applying this single optimization, I was able to reduce the runtime of our sample file from 97 seconds to 54 seconds.

### 06 - Scan file by bytes instead of lines

Building on the previous optimization, I decided to read the file byte by byte instead of line by line to avoid multiple reads of the same line. This further reduced the runtime to 44 seconds.

### 07 - Adding concurrency

I had heard that Go makes concurrency really easy, and now I finally had the chance to try it. Full disclosure: my implementation was inspired by Pereira's article mentioned earlier. Before implementing concurrency, I was keeping track of the total number of IPs and the number of IPs I had already seen. However, using an array in a concurrent environment would create race conditions. To avoid this, I opted to simply mark each IP as "seen" and then loop over the array once at the end to tally up the number of seen IPs.

This change caused my unit tests to slow down from around one second to seven seconds, as the sample size was only ten IPs. However, since my implementation is designed to handle an unlimited number of IPs, as stated in the problem, the tradeoff of this additional loop is well worth the speed increase from concurrency. After implementing concurrency, the runtime of my sample file dropped from 44 seconds to around 19 seconds. Note that about 7 of those seconds are now due to the test cases, so the actual gains on the input file are quite substantial.

Running the full file from the internal SSD I got a runtime of around 60 seconds

### 08 - Unexpected behaviour

Upon further testing it seems like there is unexpected bahaviour from accessing
the same array from multiple threads. The answer is actually an even 1 billion I
think, but due to me accessing using one array I am getting unexpected
behaviour.
