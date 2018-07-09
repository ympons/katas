def minion_game(string):
    vowels = 'AEIOU'
    n = len(string)
    kevin = 0
    stuart = 0
    for i in range(n):
        if string[i] in vowels:
            kevin += (n - i)
        else:
            stuart += (n - i)
    if kevin == stuart:
        print('Draw')
    elif kevin > stuart:
        print('Kevin', kevin)
    else:
        print('Stuart', stuart)


if __name__ == '__main__':
    s = input()
    minion_game(s)
