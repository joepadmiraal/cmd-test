#include <unistd.h>
#include <stdio.h>

int main(int argc, char* argv[])
{
    setvbuf(stdout, NULL, _IOLBF, 0);
    setvbuf(stderr, NULL, _IOLBF, 0);

    char buffer[100];
    fgets(buffer, 100, stdin);
    printf("From stdin: %s\n", buffer);

    printf("C test\n");
    sleep(1);
    printf("1\n");
    sleep(1);
    printf("done\n");
}