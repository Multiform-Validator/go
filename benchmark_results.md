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
BenchmarkIsAscii-8                    	207678693	         5.745 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	209314459	         5.726 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	209725446	         5.757 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	150144230	         8.010 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	149977216	         8.111 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	150143167	         7.992 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	23590030	        49.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	23991187	        49.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	24132750	        49.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	53366845	        19.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	55702322	        19.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	59683636	        19.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	21311055	        50.70 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	23082531	        50.42 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	23088470	        50.39 ns/op	       2 B/op	       1 allocs/op
BenchmarkIsCNPJ-8                     	18050916	        64.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18416224	        64.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18486698	        64.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	26607733	        42.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28425144	        42.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28274046	        42.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	47470536	        23.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	51602206	        23.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	50623708	        33.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	12564547	        94.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26989096	        44.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26466858	        44.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	13987899	        84.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14218268	        83.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14252133	        83.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmails-8              	 2380080	       503.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2372910	       503.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkGetOnlyEmails-8              	 2366959	       502.7 ns/op	      64 B/op	       1 allocs/op
BenchmarkIsEmail-8                    	18536428	        64.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18794103	        63.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18439892	        63.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsImage-8                    	  983563	      1302 ns/op	    5280 B/op	       5 allocs/op
BenchmarkIsImage-8                    	  978327	      1222 ns/op	    5280 B/op	       5 allocs/op
BenchmarkIsImage-8                    	  868642	      1244 ns/op	    5280 B/op	       5 allocs/op
BenchmarkIsMACAddress-8               	27259950	        42.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	28152488	        42.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	28181170	        42.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	53244010	        26.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	27660805	        37.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	58483915	        20.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	113615320	        10.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	100000000	        10.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	98595433	        10.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7421 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7422 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7396 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	28956529	        38.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	30410846	        39.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	29199415	        39.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	16762950	        68.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	17459583	        68.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	17258431	        90.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	172171321	         6.854 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	172353108	        11.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	175027756	         6.909 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	202127599	        10.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	202149201	         5.947 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	201218842	         5.957 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7407 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7399 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7400 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7420 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7399 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7405 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12827928	        92.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12992626	        93.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12838137	        92.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	35058174	        30.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	36864409	        30.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	37489160	        30.51 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Multiform-Validator/go	104.466s
```
