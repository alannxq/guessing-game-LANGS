import random

num, guessed = random.randint(1,100), False

while guessed == False:
    guess = int(input("Enter Guess from 1-100: "))

    if num == guess:
        print("Correct.\n")
        guessed = True
    elif guess > num:
        print("Too high.\n")
    else:
        print("Too low.\n")
