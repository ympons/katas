def count_substring(string, sub_string):
    ls = len(string)
    lsub = len(sub_string)
    count = 0
    for i in range(ls):
        if string[i] == sub_string[0]:
            if string[i:(i+lsub)] == sub_string:
                count += 1
    return count


if __name__ == '__main__':
    string = input().strip()
    sub_string = input().strip()

    count = count_substring(string, sub_string)
    print(count)
