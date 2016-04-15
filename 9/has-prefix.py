#!/usr/bin/env python3


def has_prefix(s, prefix):
    return len(s) >= len(prefix) and s[:len(prefix)] == prefix


def has_suffix(s, suffix):
    print(s[len(suffix):])
    return len(s) >= len(suffix) and s[len(s)-len(suffix):] == suffix


def contains(s, substr):
    for i in range(len(s) - len(substr) + 1):
        if has_prefix(s[i:], substr):
            return True
    return False

if __name__ == "__main__":
    s = "Hello, world"
    hello = "Hello"
    world = "world"

    print("expecting: True, False |", has_prefix(s, hello), has_prefix(s, world))
    print("expecting: False, True |", has_suffix(s, hello), has_suffix(s, world))
    print("expecting: true, true, false |", contains(s, hello), contains(s, world), contains(s, "something"))
