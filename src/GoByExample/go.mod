module GoByExample

go 1.18

replace helloWorld v0.0.0 => ./01_HelloWorld

replace values v0.0.0 => ./02_Values

require (
	helloWorld v0.0.0
	values v0.0.0
)
