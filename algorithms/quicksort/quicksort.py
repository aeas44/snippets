#!/usr/bin/env python3


def qsort(lst):
    if len(lst) <= 1:
        return lst

    pivot = lst[0]
    return qsort([x for x in lst[1:] if x < pivot]) + [pivot] + qsort([x for x in lst[1:] if x >= pivot])

print(qsort([3, 5, 8, 10, 2, 4, 6, 1, 7, 8]))
