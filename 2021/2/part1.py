
def find_next_position(position_string, hp, d):
    command, step = position_string.split()
    step = int(step)

    if command == 'forward':
        hp += step

    elif command == 'down':
        d += step

    elif command == 'up':
        d -= step

    return (hp, d)

if __name__ == '__main__':
    horizontal_position = 0
    depth = 0

    while True:
        input_string = input()
        if input_string == "":
            break
        horizontal_position, depth = find_next_position(input_string, horizontal_position, depth)

    print (horizontal_position * depth)