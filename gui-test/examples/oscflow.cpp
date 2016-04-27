#if defined(_MSC_VER) && !defined(_CRT_SECURE_NO_WARNINGS)
#define _CRT_SECURE_NO_WARNINGS
#endif

#include "imgui.h"
#include <ctype.h>          // toupper, isprint
#include <math.h>           // sqrtf, fabsf, fmodf, powf, cosf, sinf, floorf, ceilf
#include <stdio.h>          // vsnprintf, sscanf, printf
#include <stdlib.h>         // NULL, malloc, free, qsort, atoi


#if defined(_MSC_VER) && _MSC_VER <= 1500 // MSVC 2008 or earlier
#include <stddef.h>         // intptr_t
#else
#include <stdint.h>         // intptr_t
#endif

#ifdef _MSC_VER
#pragma warning (disable: 4996) // 'This function or variable may be unsafe': strcpy, strdup, sprintf, vsnprintf, sscanf, fopen
#define snprintf _snprintf
#endif
#ifdef __clang__
#pragma clang diagnostic ignored "-Wdeprecated-declarations"    // warning : 'xx' is deprecated: The POSIX name for this item.. // for strdup used in demo code (so user can copy & paste the code)
#pragma clang diagnostic ignored "-Wint-to-void-pointer-cast"   // warning : cast to 'void *' from smaller integer type 'int'
#pragma clang diagnostic ignored "-Wformat-security"            // warning : warning: format string is not a string literal
#endif
#ifdef __GNUC__
#pragma GCC diagnostic ignored "-Wint-to-pointer-cast"          // warning: cast to pointer from integer of different size
#pragma GCC diagnostic ignored "-Wformat-security"              // warning : format string is not a string literal (potentially insecure)
#endif

// Play it nice with Windows users. Notepad in 2015 still doesn't display text data with Unix-style \n.
#ifdef _WIN32
#define IM_NEWLINE "\r\n"
#else
#define IM_NEWLINE "\n"
#endif

#define IM_ARRAYSIZE(_ARR)  ((int)(sizeof(_ARR)/sizeof(*_ARR)))

//-----------------------------------------------------------------------------
// DEMO CODE
//-----------------------------------------------------------------------------

#ifndef IMGUI_DISABLE_TEST_WINDOWS

/*
static void ShowExampleAppConsole(bool* opened);
static void ShowExampleAppLog(bool* opened);
static void ShowExampleAppLayout(bool* opened);
static void ShowExampleAppPropertyEditor(bool* opened);
static void ShowExampleAppLongText(bool* opened);
static void ShowExampleAppAutoResize(bool* opened);
static void ShowExampleAppFixedOverlay(bool* opened);
static void ShowExampleAppManipulatingWindowTitle(bool* opened);
static void ShowExampleAppCustomRendering(bool* opened);
static void ShowExampleAppMainMenuBar();
static void ShowExampleMenuFile();

static void ShowHelpMarker(const char* desc)
{
    ImGui::TextDisabled("(?)");
    if (ImGui::IsItemHovered())
        ImGui::SetTooltip(desc);
}*/




    void ImGui::ShowTestWindow(bool* p_opened){


        if (ImGui::CollapsingHeader("Patches")){
          if (ImGui::TreeNode("Basic"))
            {
                ImGui::Columns(4, "mycolumns");
                ImGui::Separator();
                ImGui::Text("ID"); ImGui::NextColumn();
                ImGui::Text("Name"); ImGui::NextColumn();
                ImGui::Text("Path"); ImGui::NextColumn();
                ImGui::Text("Flags"); ImGui::NextColumn();
                ImGui::Separator();
                const char* names[3] = { "One", "Two", "Three" };
                const char* paths[3] = { "/path/one", "/path/two", "/path/three" };
                static int selected = -1;
                for (int i = 0; i < 3; i++)
                {
                    char label[32];
                    sprintf(label, "%04d", i);
                    if (ImGui::Selectable(label, selected == i, ImGuiSelectableFlags_SpanAllColumns))
                        selected = i;
                    ImGui::NextColumn();
                    ImGui::Text(names[i]); ImGui::NextColumn();
                    ImGui::Text(paths[i]); ImGui::NextColumn();
                    ImGui::Text("...."); ImGui::NextColumn();
                }
                ImGui::Columns(1);
                ImGui::Separator();
                ImGui::TreePop();
            }
          }
    }

#else

void ImGui::ShowTestWindow(bool*) {}



#endif
