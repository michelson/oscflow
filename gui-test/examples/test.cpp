#include <string>
#include <iostream>
#include <cstdio>
#include <memory>

//compile as : g++ -o test2 test.cpp

std::string exec(const char* cmd) {
    std::shared_ptr<FILE> pipe(popen(cmd, "r"), pclose);
    if (!pipe) return "ERROR";
    char buffer[128];
    std::string result = "";
    while (!feof(pipe.get())) {
        if (fgets(buffer, 128, pipe.get()) != NULL)
            result += buffer;
    }
    return result;
}

int main(int, char**){
    std::cout << exec("../../oscflow"); 

    return 0;
}