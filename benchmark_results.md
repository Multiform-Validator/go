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
BenchmarkIsCPF-8                      	25829748	        42.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	27965709	        42.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCPF-8                      	28407630	        42.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18097070	        64.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18415680	        64.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCNPJ-8                     	18498096	        68.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	17033280	        79.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	18488686	        76.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmail-8                    	13957755	        79.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsImage-8                    	   94929	     20156 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	   99102	     10389 ns/op	   50112 B/op	      17 allocs/op
BenchmarkIsImage-8                    	  121070	     10192 ns/op	   50112 B/op	      17 allocs/op
BenchmarkGetOnlyEmail-8               	10964496	        95.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	14367552	        95.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmail-8               	13514745	       128.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetOnlyEmails-8              	 2079584	       576.1 ns/op	      48 B/op	       2 allocs/op
BenchmarkGetOnlyEmails-8              	 2050911	       575.9 ns/op	      48 B/op	       2 allocs/op
BenchmarkGetOnlyEmails-8              	 2072548	       578.0 ns/op	      48 B/op	       2 allocs/op
BenchmarkIsCreditCard-8               	25808305	        45.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	26597169	        79.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCreditCard-8               	23326381	        45.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	49845396	        25.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	47772262	        23.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkIdentifyFlagCard-8           	50609611	        25.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsTelephone-8                	13380092	        87.98 ns/op	      24 B/op	       1 allocs/op
BenchmarkIsTelephone-8                	13702226	        90.23 ns/op	      24 B/op	       1 allocs/op
BenchmarkIsTelephone-8                	13052440	        87.47 ns/op	      24 B/op	       1 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7480 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.7506 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmpty-8                    	1000000000	         0.8591 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7482 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         0.7465 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmptyBytes-8               	1000000000	         1.531 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	157854105	         6.413 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	180464734	         7.169 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlank-8                    	186416493	         9.657 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	173865622	         6.851 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	162666216	         7.038 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBlankBytes-8               	139408048	        12.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	198535744	         6.498 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	175830313	         7.517 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAscii-8                    	178199204	         6.047 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	141833480	         8.696 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	143815688	        12.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAsciiBytes-8               	140940930	         8.225 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBase64-8                   	13879096	        84.33 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	12596839	       164.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsBase64-8                   	14428858	        99.28 ns/op	      16 B/op	       1 allocs/op
BenchmarkIsMACAddress-8               	16988506	        66.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	37040355	        31.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMACAddress-8               	37761165	        31.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsCEP-8                      	20280700	        49.70 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsCEP-8                      	22888506	        50.09 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsCEP-8                      	23310152	        49.86 ns/op	       8 B/op	       1 allocs/op
BenchmarkIsPostalCode-8               	24509829	        48.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	24149774	        49.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPostalCode-8               	24695474	        48.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	56977618	        20.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	59098382	        21.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsMD5-8                      	27617242	        42.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	58047597	        17.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	67673409	        17.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPort-8                     	67937347	        17.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7449 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7411 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPortNumber-8               	1000000000	         0.7401 ns/op	       0 B/op	       0 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	21544126	        51.78 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	22415222	        51.76 ns/op	       2 B/op	       1 allocs/op
BenchmarkCalculateCNPJCheckDigits-8   	22616444	        51.83 ns/op	       2 B/op	       1 allocs/op
BenchmarkValidateEmail-8              	12719989	       160.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12729625	        93.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateEmail-8              	12845077	        93.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	35005218	        30.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	37907996	        30.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidatePassword-8           	38405642	        30.53 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Multiform-Validator/go	108.359s
```
