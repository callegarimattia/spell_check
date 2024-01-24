# Spelling Checker
A simple spell checker, built with go. It relies on [bloom filters](https://en.wikipedia.org/wiki/Bloom_filter).

The binary comes with a packaged dictionary of 350k words.

The filter uses 2^25 bytes (8Mib) and 3 hashing functions.
It loads in ~50ms.
