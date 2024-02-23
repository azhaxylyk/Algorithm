import sys
import statistics

numbers = []

try:
    while True:
        line = input()
        if line:
            number = int(line)
            numbers.append(number)

            if len(numbers) > 1:
                mean = statistics.mean(numbers)
                stddev = statistics.stdev(numbers)

                lower_bound = max(0, int(mean - stddev))
                upper_bound = int(mean + stddev)

                print(f"{lower_bound} {upper_bound}")
            else:
                print("100 200")
except EOFError:
    pass
