#include <stdio.h>
#include <stdlib.h>

int main() {
  int guess;
  int seed = rand() % 20;
  int number = rand() % seed;

  while (guess != number) {
    printf("Enter number from 1-100: ");
    scanf("%d", &guess);

    if (number == guess) {
      printf("Correct\n\n");
    } else if (guess > number) {
      printf("Too high.\n\n");
    } else {
      printf("Too low.\n\n");
    }
  }

  return 0;
}
