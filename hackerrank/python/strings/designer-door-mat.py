def printer(n, m):
    k = n//2

    # top
    j = 1
    for i in range(k):
        print(('.|.' * j).center(m, '-'))
        j += 2

    # center
    print('WELCOME'.center(m, '-'))

    # bottom
    j -= 2
    for i in range(k):
        print(('.|.' * j).center(m, '-'))
        j -= 2


if __name__ == '__main__':
    n, m = map(int, input().split())
    printer(n, m)
