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
BenchmarkIsCPF-8                      	23126463	        49.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	26171588	        43.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	27602522	        48.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	  206311	      6029 ns/op	   13493 B/op	      10 allocs/op
BenchmarkIsCNPJ-8                     	  315790	      4564 ns/op	   13493 B/op	      10 allocs/op
BenchmarkIsCNPJ-8                     	  309770	      7888 ns/op	   13493 B/op	      10 allocs/op
BenchmarkIsEmail-8                    	 4590588	       340.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsEmail-8                    	 4279288	       373.5 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsEmail-8                    	 4201359	       252.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsImage-8                    	   38868	     35517 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	   72355	     15164 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	   83514	     12909 ns/op	   50112 B/op	      17 allocs/op
BenchmarkGetOnlyEmail-8               	  341948	      3072 ns/op	     289 B/op	       6 allocs/op
BenchmarkGetOnlyEmail-8               	  486812	      2462 ns/op	     289 B/op	       6 allocs/op
BenchmarkGetOnlyEmail-8               	  183435	      6924 ns/op	     289 B/op	       6 allocs/op
BenchmarkGetOnlyEmails-8              	  275668	      5350 ns/op	     289 B/op	       6 allocs/op
BenchmarkGetOnlyEmails-8              	  356125	      4599 ns/op	     290 B/op	       6 allocs/op
BenchmarkGetOnlyEmails-8              	  184429	      6185 ns/op	     289 B/op	       6 allocs/op
BenchmarkIsCreditCard-8               	26529297	        47.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	25910809	       102.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26357106	        55.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	32618916	        38.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	39992862	        30.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	39279424	        54.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	 4291180	       249.6 ns/op	      48 B/op	       4 allocs/op
BenchmarkIsTelephone-8                	 2797058	       551.8 ns/op	      48 B/op	       4 allocs/op
BenchmarkIsTelephone-8                	 4713727	       250.7 ns/op	      48 B/op	       4 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.8484 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.8892 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.8202 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7979 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7707 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.9907 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	161523198	        17.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	160892792	         7.324 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	151657674	        11.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	96703677	        11.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	100000000	        12.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	75722124	        14.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	100000000	        17.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	69659012	        18.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	131347182	         8.218 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	151787788	         8.967 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	156962594	         6.713 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	183756038	         7.712 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	 9738200	       119.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	 8725184	       169.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	 9674611	       117.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsMACAddress-8               	11872922	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	24098409	        45.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	27762655	        47.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	  569832	      3105 ns/op	    6648 B/op	       5 allocs/op
BenchmarkIsCEP-8                      	  776797	      1483 ns/op	    6648 B/op	       5 allocs/op
BenchmarkIsCEP-8                      	  751266	      2128 ns/op	    6648 B/op	       5 allocs/op
BenchmarkIsPostalCode-8               	 3410662	       344.0 ns/op	      24 B/op	       2 allocs/op
BenchmarkIsPostalCode-8               	 3200193	       646.1 ns/op	      24 B/op	       2 allocs/op
BenchmarkIsPostalCode-8               	 3345950	       348.2 ns/op	      24 B/op	       2 allocs/op
BenchmarkIsMD5-8                      	39872613	        33.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	32271849	        31.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	39421736	        51.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	51923690	        26.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	25813126	        41.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	57661208	        23.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	690171984	         1.546 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7621 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         1.301 ns/op	       0 B/op	       0 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	  249001	      4505 ns/op	    6740 B/op	       5 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	  550153	      1937 ns/op	    6740 B/op	       5 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	  576297	      1871 ns/op	    6740 B/op	       5 allocs/op
BenchmarkValidateEmail-8              	 3765711	       349.6 ns/op	      32 B/op	       1 allocs/op
BenchmarkValidateEmail-8              	 4362338	       262.6 ns/op	      32 B/op	       1 allocs/op
BenchmarkValidateEmail-8              	 4243994	       389.4 ns/op	      32 B/op	       1 allocs/op
BenchmarkValidatePassword-8           	37874306	        71.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	37120987	        31.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	41285370	        40.21 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Multiform-Validator/go	137.802s
```
