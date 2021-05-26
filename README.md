# md5hasher

Small tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

- Able to perform the requests in parallel so that the tool can complete sooner.
- Able to limit the number of parallel requests (defaulted to 10), to prevent exhausting local resources.
- Only Go's standard library is used.
- Unit tests included.