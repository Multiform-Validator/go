# Benchmark Results

Date: 2026-04-28

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
BenchmarkIsCPF-8                      	23985520	        43.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	26547168	        47.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	25967884	        49.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	17112345	        66.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18045824	        65.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18732240	        69.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	15119090	        78.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18056016	        68.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18487969	        64.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsImage-8                    	   97420	     13329 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  102668	     12071 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  116532	     10887 ns/op	   50112 B/op	      17 allocs/op
BenchmarkGetOnlyEmail-8               	12341325	        89.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14333071	       124.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	13643770	       174.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmails-8              	  955064	      1223 ns/op	      48 B/op	       2 allocs/op
BenchmarkGetOnlyEmails-8              	 1465035	       691.1 ns/op	      48 B/op	       2 allocs/op
BenchmarkGetOnlyEmails-8              	 1869627	       713.2 ns/op	      48 B/op	       2 allocs/op
BenchmarkIsCreditCard-8               	21923679	        49.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	24849074	        46.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26955946	        47.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	46126622	        25.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	47860503	        25.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	49711734	        23.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	12297757	       113.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkIsTelephone-8                	11928669	       106.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkIsTelephone-8                	13076581	        93.86 ns/op	      24 B/op	       1 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.8265 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.8054 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         1.244 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7586 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7655 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.8320 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	162356743	         7.807 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	137482558	         9.402 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	41113383	        43.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	94132581	        13.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	168582543	         6.843 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	191076873	        12.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	95240913	        13.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	135490028	         9.550 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	139261137	         8.179 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	209179333	         6.553 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	171139624	         6.096 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	100000000	        10.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	13743694	        87.53 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14550282	        85.66 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	11380976	        91.42 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsMACAddress-8               	33741732	        33.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	38323087	        55.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	35380411	        32.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	20615163	        56.41 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsCEP-8                      	22448655	        54.30 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsCEP-8                      	16981377	        69.98 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsPostalCode-8               	14311070	        80.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	17380232	        64.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	20351547	        60.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	39887242	        40.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	40661565	        50.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	36803774	        29.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	62356310	        18.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	69069613	        18.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	64053025	        18.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.8167 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7515 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7775 ns/op	       0 B/op	       0 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	11464106	       123.4 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	19715216	        55.90 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	21794550	        72.21 ns/op	       2 B/op	       1 allocs/op
BenchmarkValidateEmail-8              	 5747312	       188.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12532660	        94.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12781429	       100.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	43556815	        31.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	43265127	        28.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	42806572	        35.24 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Multiform-Validator/go	116.641s
```
