//  do not allocate memory for the result
//  write to pre-allocated pointer *camel

#include <stdbool.h>
#include <ctype.h>
#include <stdio.h>
#include <string.h>

void to_camel_case(const char *text, char *camel) {

    //  <----  hajime!

    int length = strlen(text);
    int j = 0;

    if (length == 0) {
        *camel = '\0';
    }

    bool isNewWord = false;
    for (int i = 0; text[i] != '\0'; i++) {
        if (i == 0) {
            camel[j] = text[i];
            j++;
            continue;
        }
        if (text[i] == '-' || text[i] == '_') {
            isNewWord = true;
            continue;
        }
        if (isNewWord) {
            camel[j] = toupper(text[i]);
            j++;
            isNewWord = false;
            continue;
        }
        camel[j] = tolower(text[i]);
        j++;
    }
}

int main() {
    char camel[100];
    char *text = "sddDWsbd_dfdfDS-ccDS";
    to_camel_case(text, camel);
    printf("%s", camel);
}