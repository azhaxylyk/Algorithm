import statistics
def main():
    numbers = []
    mean = 150
    variance = 0
    count = 0
    while True:
        try:
            line = input()
            if line:
                number = int(line)
                numbers.append(number)
                count += 1
                delta = number - mean
                mean += delta / count
                variance += delta * (number - mean)
                if count > 1:
                    stddev = (variance / (count - 1)) ** 0.5
                    lower_bound = max(0, int(mean - stddev)) - 100
                    upper_bound = int(mean + stddev) + 100
                    print(f"{lower_bound} {upper_bound}")
                else:
                    print("100 200")
        except EOFError:
            break
if __name__ == "__main__":
    main()

