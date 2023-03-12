#include <stddef.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

//  return a dynamically allocated int array
//  the return array will be freed by tester
//  set *szout to the length of output array

typedef struct {
    int value;
    int repeat;
} item;

int *delete_nth(size_t szin, const int order[szin], int max_e, size_t *szout) {

    //  <----  hajime!
    int *newOrder = malloc(szin * sizeof(int));
    item *items = malloc(szin * sizeof(int));

    int itemLength = 0;
    int newOrderIndex = 0;

    for (int i = 0; i < (int) szin; i++) {
        bool doItemRepeat = false;
        item current;
        for (int j = 0; j < itemLength; j++) {
            if (items[j].value == order[i]) {
                items[j].repeat++;
                current = items[j];
                doItemRepeat = true;
            }
        }

        if (!doItemRepeat) {
            item newItem = {.value=order[i], .repeat=1};
            current = newItem;
            items[itemLength] = newItem;
            itemLength++;
        }

        if (current.repeat <= max_e) {
            newOrder[newOrderIndex] = current.value;
            newOrderIndex++;
        }
    }

    *szout = newOrderIndex;
    free(items);
    return newOrder;
}

int main() {
    const int order[9] = {1, 1, 3, 3, 7, 2, 2, 2, 2};
    int max_e = 3;
    int expected[8] = {1, 1, 3, 3, 7, 2, 2, 2};
    size_t szout = 0;
    int *submitted = delete_nth(9, order, max_e, &szout);

    printf("was:\n");
    for (int i = 0; i < 9; ++i) {
        printf("%d ", order[i]);
    }

    printf("\nexpected:\n");
    for (int i = 0; i < szout; ++i) {
        printf("%d ", expected[i]);
    }

    printf("\nreally:\n");
    for (int i = 0; i < szout; ++i) {
        printf("%d ", submitted[i]);
    }

    return 0;
}