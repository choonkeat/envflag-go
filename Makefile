all:
	make run-example-default-1
	make run-example-default-2
	make run-example-nodefault-1
	make run-example-nodefault-2

run-example-default-1:
	# should succeed
	env \
		MY_BOOL=true \
		MY_DURATION=1m \
		MY_FLOAT64=3.14 \
		MY_INT64=1234 \
		MY_INT=12 \
		MY_STRING=strrr \
		MY_TEXT=info \
		MY_UINT64=12345 \
		MY_UINT=123 \
		go run example/default/main.go

run-example-default-2:
	# should succeed since there are default values
	env \
		go run example/default/main.go

run-example-nodefault-1:
	# should succeed
	env \
		MY_BOOL=true \
		MY_DURATION=1m \
		MY_FLOAT64=3.14 \
		MY_INT64=1234 \
		MY_INT=12 \
		MY_STRING=strrr \
		MY_TEXT=info \
		MY_UINT64=12345 \
		MY_UINT=123 \
		go run example/nodefault/main.go

run-example-nodefault-2:
	# should fail since env vars are not set
	env \
		go run example/nodefault/main.go || echo expected failure
