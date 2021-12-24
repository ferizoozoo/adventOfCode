#include <iostream>
#include <string>
#include <bits/stdc++.h>
#include <unordered_map>

using namespace std;

int main()
{
    unordered_map<int, vector<string>> mapOfNumberList;
    string alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

    string currentLine;

    while (true)
    {
        vector<string> chars;

        getline(cin, currentLine);
        if (currentLine.empty())
        {
            break;
        }

        stringstream ss(currentLine);

        string word;
        string heightString;

        ss >> heightString;

        while (getline(ss, word, ' '))
        {
            chars.push_back(word);
        }

        int height = stoi(heightString);

        mapOfNumberList[height] = chars;
    }

    unordered_map<string, int> charToNumber;

    for (auto heightList : mapOfNumberList)
    {
        for (auto character : heightList.second)
        {
            charToNumber[character] += heightList.first;
        }
    }

    int increases = 0;
    int previousNumber = INT_MIN;

    for (auto character : alphabet)
    {
        string characterString(1, character);
        int sumOfHeights = charToNumber[characterString];

        if (sumOfHeights > previousNumber)
        {
            increases++;
            cout << characterString << ": " << sumOfHeights << endl;
            previousNumber = sumOfHeights;
        }
    }

    cout << increases - 1;
}