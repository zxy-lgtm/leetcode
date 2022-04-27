#include <stdio.h>
static char buff[256];
char *string;
int main()
{
	int i,n;
    char c;
    string = buff;
	printf("Please input a string:");
	gets(string);
	printf("\nYour string is: %s\n", string);
n=strlen(string);
	for(i=0; i<n/2; i++)
{
        c=string[i];
 		string[i]=string[n-i-1];
 		string[n-i-1]=c;
}
printf("\n After conversion, the string is:%s\n", string);
return 0;
}