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
BenchmarkIsCPF-8                      	25677463	        42.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28286192	        42.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28258260	        42.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18228805	        64.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18481212	        65.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18352725	        65.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	19349203	        61.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	19419975	        61.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	19190661	        61.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsImage-8                    	  134359	      8461 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  129480	      8612 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  142246	      8553 ns/op	   50112 B/op	      17 allocs/op
BenchmarkGetOnlyEmail-8               	14125465	        82.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14394710	       133.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14462578	        82.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmails-8              	 2417076	       491.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2434869	       497.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2427826	       493.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkIsCreditCard-8               	26451318	        44.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26562358	        44.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26358364	        44.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	48838003	        23.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	48945838	        23.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	49459468	        23.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	17628608	        66.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	17565877	        81.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	 8341159	       129.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7401 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7414 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7396 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7401 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7408 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7399 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	173199450	         6.837 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	174977613	         6.855 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	173821504	         6.855 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	202068315	        10.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	193950901	         6.808 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	193506793	         5.944 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	150038029	         8.002 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	150134925	         8.010 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	150251272	         8.483 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	90253243	        12.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	211677510	         5.704 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	212017075	         5.672 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	14448262	        79.66 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14560807	        80.73 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14432476	        80.53 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsMACAddress-8               	37458418	        31.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	38033506	        31.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	37792764	        31.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	41760457	        28.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	41100772	        46.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	36322651	        28.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	30626709	        37.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	30947910	        38.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	22771251	        46.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	40967320	        29.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	41577283	        28.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	41113022	        28.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	62435730	        31.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	61564410	        17.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	63570004	        17.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         1.056 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	754214071	         1.340 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7398 ns/op	       0 B/op	       0 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	20202388	        57.45 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	20643710	        57.25 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	20120043	        57.31 ns/op	       2 B/op	       1 allocs/op
BenchmarkValidateEmail-8              	12642088	        93.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12655480	        93.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12594968	        93.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	39199057	        28.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	40588633	        28.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	41523330	        28.68 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Multiform-Validator/go	104.269s
```
