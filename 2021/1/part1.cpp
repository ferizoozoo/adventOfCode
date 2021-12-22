#include <iostream>
#include <string>
#include <bits/stdc++.h>

using namespace std;

int main()
{
    int previousNumber = INT_MAX;
    int currentNumber = 0;
    int increases = 0;
    string currentNumberString;

    while (true)
    {
        getline(cin, currentNumberString);
        if (currentNumberString.empty())
            break;

        currentNumber = stoi(currentNumberString);
        if (currentNumber > previousNumber)
            increases++;

        previousNumber = currentNumber;
    }

    cout << increases;
}