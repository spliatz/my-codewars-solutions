from itertools import combinations


def choose_best_sum(t, k, ls):
    result = 0
    for i in combinations(ls, k):
        if result < sum(i) <= t:
            result = sum(i)
            print(i)
    return result if result != 0 else None
