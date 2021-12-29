
def find_next_position(position_string, hp, d, aim):
    command, step = position_string.split()
    step = int(step)

    if command == 'forward':
        hp += step
        d += aim * step

    elif command == 'down':
        aim += step

    elif command == 'up':
        aim -= step

    return (hp, d, aim)

if __name__ == '__main__':
    horizontal_position = 0
    depth = 0
    aim = 0

    while True:
        input_string = input()
        if input_string == "":
            break
        horizontal_position, depth, aim = find_next_position(input_string, horizontal_position, depth, aim)

    print (horizontal_position * depth)