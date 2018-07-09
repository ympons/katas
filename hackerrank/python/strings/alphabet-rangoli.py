def print_rangoli(size):
    alpha = list('abcdefghijklmnopqrstuvwxyz')
    L = []
    for i in range(size):
        s = "-".join(alpha[i:size])
        L.append(s[::-1]+s[1:])

    width = len(L[0])

    for i in range(size-1, 0, -1):
        print(L[i].center(width, "-"))

    for i in range(size):
        print(L[i].center(width, "-"))


if __name__ == '__main__':
    n = int(input())
    print_rangoli(n)
