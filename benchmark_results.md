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
BenchmarkIsAscii-8                    	125891698	         8.694 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	148528051	         9.102 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	149962657	         8.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	210643074	         5.691 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	210441712	         5.678 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	210903051	         5.701 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	14258272	       140.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14584683	        80.03 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14780448	        80.39 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsCEP-8                      	40087174	        28.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	41480268	        28.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	41308083	        28.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	20312304	        57.40 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	20454328	        57.34 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	20374476	        57.40 ns/op	       2 B/op	       1 allocs/op
BenchmarkIsCNPJ-8                     	18145111	        64.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18289208	        64.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18524145	        64.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28068721	        42.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	27991548	        42.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28041944	        42.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	48537405	        23.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	48766923	        30.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	23720115	        44.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	25821201	        44.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26472031	        44.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26894302	        44.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	13985656	        83.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14158611	        83.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14094722	        84.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmails-8              	 2403391	       502.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2397529	       496.7 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2410417	       496.8 ns/op	      64 B/op	       1 allocs/op
BenchmarkIsEmail-8                    	18803203	        61.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	19562271	        78.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	 9144512	       110.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsImage-8                    	  131452	      8459 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  138614	      8651 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  137293	      8721 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsMACAddress-8               	33921932	        31.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	37421983	        45.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	38110636	        31.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	39649756	        28.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	41067378	        28.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	38665041	        28.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	62912430	        17.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	67956664	        17.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	66088254	        17.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7420 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7411 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7400 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	30942618	        65.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	31340359	        39.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	14537077	        81.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	15165488	        70.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	17443378	        68.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	17560831	        68.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	186874452	         6.652 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	184204279	         6.527 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	183786220	         6.518 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	163738524	        11.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	176926144	         6.828 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	175395644	        11.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7412 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7424 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7406 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7406 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7426 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7415 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12525381	        94.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12692792	        94.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12527179	        93.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	37320942	        28.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	41432389	        29.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	39311390	        28.70 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Multiform-Validator/go	104.509s
```
