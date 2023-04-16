def high_and_low(numbers):
    a = [int(i) for i in numbers.split(" ")]
    return f"{max(a)} {min(a)}"
