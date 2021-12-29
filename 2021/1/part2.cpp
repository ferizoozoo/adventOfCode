#include <iostream>
#include <string>
#include <bits/stdc++.h>
#include <unordered_map>

using namespace std;

bool is_number(const string& s)
{
    string::const_iterator it = s.begin();
    while (it != s.end() && isdigit(*it)) ++it;
    return !s.empty() && it == s.end();
}

int main()
{
    // unordered_map<int, vector<string>> mapOfNumberList;
    // unordered_map<string, int> charToNumber;
    // string alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

    // string currentLine;

    // while (true)
    // {
    //     vector<string> inputChars;

    //     getline(cin, currentLine);
    //     if (currentLine.empty())
    //         break;

    //     stringstream ss(currentLine);
    //     string word;

    //     while (ss >> word)
    //     {
    //         inputChars.push_back(word);
    //     }

    //     int height = stoi(inputChars[0]);

    //     vector<string> characters;

    //     for (auto ch : inputChars)
    //     {
    //         if (!is_number(ch))
    //         {
    //             characters.push_back(ch);
    //         }
    //     }

    //     for (auto character : characters)
    //     {
    //         charToNumber[character] += height;
    //     }
    // }

    // int increases = 0;
    // int previousNumber = INT_MIN;

    // for (auto character : alphabet)
    // {
    //     string characterString(1, character);
    //     int sumOfHeights = charToNumber[characterString];

    //     if (sumOfHeights > previousNumber)
    //     {
    //         increases++;
    //         cout << characterString << ": " << sumOfHeights << endl;
    //         previousNumber = sumOfHeights;
    //     }
    // }

    // cout << increases - 1;




    int previousNumber = INT_MIN;
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

    cout << increases - 1;
}