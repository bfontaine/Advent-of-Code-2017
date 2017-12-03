#include <stdio.h>
#include <string.h>

const char ASCII_ZERO = '0';

int solve_captcha(char *captcha) {
        int size = strlen(captcha),
            sum = 0,
            i, j, c1, c2;

        for (i=0; i<size; i++) {
                j = (i+1) % size;
                c1 = captcha[i];
                c2 = captcha[j];

                if (c1 == c2) {
                        sum += c1 - ASCII_ZERO;
                }
        }

        return sum;
}

int main(int argc, char **argv) {
        int solution;

        if (argc != 2) {
                printf("Usage:\n    %s <number>\n", argv[0]);
                return 1;
        }

        solution = solve_captcha(argv[1]);
        printf("Solution: %d\n", solution);

        return 1;
}
