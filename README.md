# Spelling Checker
A simple spell checker, built with go. It relies on [bloom filters](https://en.wikipedia.org/wiki/Bloom_filter).

The binary comes with a packaged dictionary of 350k words.

The filter uses 2^25 bit (8MB) and 3 hashing functions.
The parameters are calculated using this cool [website](https://hur.st/bloomfilter/?n=350000&p=&m=67108864&k=3)

The whole dictionary gets loaded in ~50ms.

