#include <stdio.h>
#include <string.h>
#include <assert.h>

const char ASCII_ZERO = '0';

int solve_captcha(char *captcha, int size, int step) {
        int sum = 0,
            i, j, c1, c2;

        for (i=0; i<size; i++) {
                j = (i+step) % size;
                c1 = captcha[i];
                c2 = captcha[j];

                if (c1 == c2) {
                        sum += c1 - ASCII_ZERO;
                }
        }

        return sum;
}

int solve_captcha_part_one(char *captcha) {
        return solve_captcha(captcha, strlen(captcha), 1);
}

int solve_captcha_part_two(char *captcha) {
        int size = strlen(captcha);

        assert(size % 2 == 0);
        return solve_captcha(captcha, size, size/2);
}

int main(int argc, char **argv) {
        int solution;

        if (argc != 2) {
                printf("Usage:\n    %s <number>\n", argv[0]);
                return 1;
        }

        printf("Solution 1: %d\n", solve_captcha_part_one(argv[1]));
        printf("Solution 2: %d\n", solve_captcha_part_two(argv[1]));

        return 1;
}
