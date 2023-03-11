#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

bool isVowel(char x) {
    switch (tolower(x)) {
        case 'a':
        case 'e':
        case 'i':
        case 'u':
        case 'o':
            return true;
        default: return false;
    }
}

char *disemvowel(const char *str)
{
    int length = (int) strlen(str);
    char *array = malloc(length);
    int j = 0;
    for (int i = 0; i < length; i++) {
        if (!isVowel(str[i])) {
            array[j] = str[i];
            j++;
        }
    }
    return array;
}

int main()
{
    char *array = disemvowel("No offense but, Your writing is among the worst I've ever read");
    for (int i = 0; array[i] != '\0'; i++) {
        printf("%c", array[i]);
    }
}
