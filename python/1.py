import numpy as np


def parse():
    fpath = "1.txt"
    global lList, rList

    with open(fpath) as fhand:
        lines = fhand.readlines()

    for line in lines:
        pair = line.split()
        rList.append(int(pair[1]))
        lList.append(int(pair[0]))


def distances():
    global lList, rList
    lList = sorted(lList)
    rList = sorted(rList)
    distance = np.sum(np.abs(np.array(lList) - np.array(rList)))
    print("distance: ", distance)


def similarity():
    global lList, rList
    count_map = dict.fromkeys(lList, 0)
    for num in rList:
        try:
            count_map[num] += 1
        except KeyError:
            continue
    print("similarity: ", np.sum(
        np.array(count_map.values()) * np.array(count_map.keys())))


def main():
    global rList, lList
    lList = []
    rList = []
    parse()
    distances()
    similarity()


if __name__ == "__main__":
    main()
