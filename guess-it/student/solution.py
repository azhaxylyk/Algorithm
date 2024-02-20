import sys
import statistics

# Инициализируем список для хранения введенных чисел
numbers = []

try:
    while True:
        # Читаем число из стандартного ввода
        line = input()
        if line:
            number = int(line)
            numbers.append(number)

            # Проверяем, достаточно ли у нас данных для предсказания
            if len(numbers) > 1:
                # Рассчитываем среднее и стандартное отклонение
                mean = statistics.mean(numbers)
                stddev = statistics.stdev(numbers)

                # Определяем диапазон как mean ± stddev
                lower_bound = max(0, int(mean - stddev))  # Делаем нижнюю границу не меньше 0
                upper_bound = int(mean + stddev)

                print(f"{lower_bound} {upper_bound}")
            else:
                # Если данных недостаточно, выводим базовый диапазон
                print("100 200")
except EOFError:
    # Завершаем программу при достижении конца файла ввода
    pass
