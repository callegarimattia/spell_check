# Spelling Checker
A simple spell checker, built with go. It relies on [bloom filters](https://en.wikipedia.org/wiki/Bloom_filter).

The binary comes with a packaged dictionary of 350k words.

The filter uses 2^25 bytes (8MiB) and 3 hashing functions.

The parameters are calculated using this cool [website](https://hur.st/bloomfilter/?n=350000&p=&m=67108864&k=3)
It loads in ~50ms.
