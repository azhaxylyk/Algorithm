# math-skills

## Objectives

The goal of this project is to perform the following calculations on a given dataset:

- Average
- Median
- Variance
- Standard Deviation

## Usage

First, clone the repository:
```bash
git clone git@github.com:azhaxylyk/math-skills.git
```

Navigate to the project directory:
```bash
cd math-skills/project
```

In the project directory, you will find three files: data_large.txt, data_small.txt, and data_with_negative.txt.

To execute the program with one of the datasets, run the following command:
```bash
go run main.go data_large.txt
```
Replace data_large.txt with the dataset you want to analyze.


After reading the file, the program will execute each of the calculations mentioned above and print the results in the following format (rounded to the nearest integer):
```bash
Average: 5075
Median: 4942
Variance: 8039115
Standard Deviation: 2835
```


Please note that the values are rounded to the nearest integer by default. If you need to round to the smallest integer, you can uncomment the appropriate lines in the Go program:
```bash
fmt.Printf("Average: %.0f\n", math.Floor(mean))
fmt.Printf("Median: %.0f\n", math.Floor(median))
fmt.Printf("Variance: %.0f\n", math.Floor(variance))
fmt.Printf("Standard Deviation: %.0f\n", math.Floor(stddev))
```
