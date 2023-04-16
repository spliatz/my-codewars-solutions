def sort_array(source_array):
    odds = [elem for elem in source_array if elem % 2 != 0]
    odds.sort()
    odds_index = 0
    for i in range(len(source_array)):
        if source_array[i] % 2 != 0:
            source_array[i] = odds[odds_index]
            odds_index += 1
    return source_array