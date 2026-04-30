# Benchmark Results

Date: 2026-04-28 09:31 PM

Command:

```bash
go test -bench=. -benchmem -count=3 .
```

Environment:

```text
goos: linux
goarch: amd64
pkg: github.com/Multiform-Validator/go
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
```

Results:

```text
BenchmarkIsCPF-8                      	26688848	        42.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28264622	        42.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	27905773	        42.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18087316	        65.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18369090	        65.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18375010	        65.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18228982	        65.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18174412	        66.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18071754	        65.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsImage-8                    	  136604	      8482 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  135961	      9141 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  129445	      8691 ns/op	   50112 B/op	      17 allocs/op
BenchmarkGetOnlyEmail-8               	13724342	       150.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	 6590491	       180.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	13052511	        83.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmails-8              	 2333070	       513.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2325650	       512.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2337613	       514.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkIsCreditCard-8               	26503449	        44.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	25069243	        44.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	23177482	        44.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	47750968	        23.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	48662464	        26.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	24006829	        49.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	 8095202	       145.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	16319037	        68.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	17511991	        68.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7491 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7416 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7453 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7448 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7410 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7417 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	190589821	         6.306 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	189598492	         6.280 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	142893375	        13.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	177219766	         6.844 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	176021240	         6.781 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	176317543	         6.755 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	205203882	         5.812 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	206664570	         5.756 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	209208256	         5.786 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	149841196	         8.002 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	149696684	         8.023 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	149930660	         7.998 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	13938824	       114.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14729943	        81.12 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14358390	        81.43 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsMACAddress-8               	28565487	        41.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	28549777	        41.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	28740440	        41.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	20654334	        51.17 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsCEP-8                      	22778565	        50.68 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsCEP-8                      	22360041	        50.71 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsPostalCode-8               	30871758	        37.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	29882276	        37.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	30426804	        38.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	57694756	        30.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	59447718	        20.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	58676397	        20.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	60065464	        18.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	59566054	        18.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	63241695	        17.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7454 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7423 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7413 ns/op	       0 B/op	       0 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	21429632	        61.13 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	10694272	       107.0 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	22837424	        51.62 ns/op	       2 B/op	       1 allocs/op
BenchmarkValidateEmail-8              	12056073	        95.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12581536	       163.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12474750	        95.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	36458324	        30.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	38346880	        30.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	38515418	        30.65 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Multiform-Validator/go	106.594s
```
