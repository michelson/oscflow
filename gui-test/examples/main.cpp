// ImGui - standalone example application for Glfw + OpenGL 3, using programmable pipeline
// If you are new to ImGui, see examples/README.txt and documentation at the top of imgui.cpp.

#include <imgui.h>
#include "imgui_impl_glfw_gl3.h"
#include <stdio.h>
#include <GL/gl3w.h>
#include <GLFW/glfw3.h>

static void error_callback(int error, const char* description)
{
    fprintf(stderr, "Error %d: %s\n", error, description);
}

int main(int, char**)
{
    // Setup window
    glfwSetErrorCallback(error_callback);
    if (!glfwInit())
        return 1;
    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 3);
    glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);
#if __APPLE__
    glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GL_TRUE);
#endif
    GLFWwindow* window = glfwCreateWindow(320, 240, "OSCFLOW", NULL, NULL);
    glfwMakeContextCurrent(window);
    gl3wInit();

    // Setup ImGui binding
    ImGui_ImplGlfwGL3_Init(window, true);

    // Load Fonts
    // (there is a default font, this is only if you want to change it. see extra_fonts/README.txt for more details)
    //ImGuiIO& io = ImGui::GetIO();
    //io.Fonts->AddFontDefault();
    //io.Fonts->AddFontFromFileTTF("../../extra_fonts/Cousine-Regular.ttf", 15.0f);
    //io.Fonts->AddFontFromFileTTF("../../extra_fonts/DroidSans.ttf", 16.0f);
    //io.Fonts->AddFontFromFileTTF("../../extra_fonts/ProggyClean.ttf", 13.0f);
    //io.Fonts->AddFontFromFileTTF("../../extra_fonts/ProggyTiny.ttf", 10.0f);
    //io.Fonts->AddFontFromFileTTF("c:\\Windows\\Fonts\\ArialUni.ttf", 18.0f, NULL, io.Fonts->GetGlyphRangesJapanese());

    bool show_test_window = true;

    bool no_titlebar = true;
    bool no_border = true;
    bool no_resize = true;
    bool no_move = true;
    bool no_scrollbar = false;
    bool no_collapse = false;
    bool no_menu = false;

    bool activate_play = false;
    bool activate_stop = false;
    bool activate_rec = false;

    bool show_another_window = false;
    ImVec4 clear_color = ImColor(114, 144, 154);

    ImGuiWindowFlags window_flags = 0;
    if (no_scrollbar) window_flags |= ImGuiWindowFlags_NoScrollbar;
    if (no_move)      window_flags |= ImGuiWindowFlags_NoMove;

    // Main loop
    while (!glfwWindowShouldClose(window))
    {
        glfwPollEvents();
        ImGui_ImplGlfwGL3_NewFrame();


        // Demonstrate the various window flags. Typically you would just use the default.
        ImGuiWindowFlags window_flags = 0;
        if (no_titlebar)  window_flags |= ImGuiWindowFlags_NoTitleBar;
        if (!no_border)   window_flags |= ImGuiWindowFlags_ShowBorders;
        if (no_resize)    window_flags |= ImGuiWindowFlags_NoResize;
        if (no_move)      window_flags |= ImGuiWindowFlags_NoMove;
        if (no_scrollbar) window_flags |= ImGuiWindowFlags_NoScrollbar;
        if (no_collapse)  window_flags |= ImGuiWindowFlags_NoCollapse;
        if (!no_menu)     window_flags |= ImGuiWindowFlags_MenuBar;
        ImGui::SetNextWindowSize(ImVec2(550,680), ImGuiSetCond_FirstUseEver);

        // 1. Show a simple window
        // Tip: if we don't call ImGui::Begin()/ImGui::End() the widgets appears in a window automatically called "Debug"
        {
            ImGui::SetNextWindowSize(ImVec2(320,240), ImGuiSetCond_FirstUseEver);

            ImGui::Begin("OscFlow General Settings", &show_another_window, window_flags);

            if (ImGui::BeginMenuBar())
            {
                if (ImGui::BeginMenu("Menu"))
                {
                    //ShowExampleMenuFile();
                    if (ImGui::MenuItem("Undo", "CTRL+Z")) {}
                    if (ImGui::MenuItem("Redo", "CTRL+Y", false, false)) {}  // Disabled item
                    ImGui::Separator();
                    if (ImGui::MenuItem("Cut", "CTRL+X")) {}
                    if (ImGui::MenuItem("Copy", "CTRL+C")) {}
                    if (ImGui::MenuItem("Paste", "CTRL+V")) {}
                    ImGui::EndMenu();
                }
                ImGui::EndMenuBar();
            }

            static float f = 0.0f;
            ImGui::Text("OSCFLOW v.0.0.1");  ImGui::SameLine();
            if (ImGui::Button("Show patches")) show_test_window ^= 1; 

            ImGui::Text("volume");
            ImGui::SliderFloat("", &f, 0.0f, 1.0f);
            //ImGui::ColorEdit3("clear color", (float*)&clear_color);
            
            ImGui::AlignFirstTextHeightToWidgets();
            if (ImGui::Button("Play")) activate_play ^= 1; ImGui::SameLine();
            if (ImGui::Button("Stop")) activate_stop ^= 1; ImGui::SameLine();
            if (ImGui::Button("Rec"))  activate_rec ^= 1;

            //ImGui::Checkbox("No titlebar", &no_titlebar); ImGui::SameLine(150);
            //ImGui::Checkbox("No border", &no_border); ImGui::SameLine(300);
            //ImGui::Checkbox("No resize", &no_resize);
            //ImGui::Checkbox("No move", &no_move); ImGui::SameLine(150);
            //ImGui::Checkbox("No scrollbar", &no_scrollbar); ImGui::SameLine(300);
            //ImGui::Checkbox("No collapse", &no_collapse);
            //ImGui::Checkbox("No menu", &no_menu);

            //if (ImGui::Button("Test Window")) show_test_window ^= 1;
            //if (ImGui::Button("Another Window")) show_another_window ^= 1;
            
            ImGui::Text("Application average %.3f ms/frame (%.1f FPS)", 1000.0f / ImGui::GetIO().Framerate, ImGui::GetIO().Framerate);
            
            // 3. Show the ImGui test window. Most of the sample code is in ImGui::ShowTestWindow()
            if (show_test_window)
            {
                ImGui::ShowTestWindow(&show_test_window);
            }

            ImGui::End();
        }

        //ImGui::Separator();


        // 2. Show another simple window, this time using an explicit Begin/End pair
        /*
        if (show_another_window)
        {
            ImGui::SetNextWindowSize(ImVec2(200,100), ImGuiSetCond_FirstUseEver);
            ImGui::Begin("Another Window", &show_another_window);
            ImGui::Text("Hello");
            ImGui::End();
        }*/

        // Rendering
        int display_w, display_h;
        glfwGetFramebufferSize(window, &display_w, &display_h);
        glViewport(0, 0, display_w, display_h);
        glClearColor(clear_color.x, clear_color.y, clear_color.z, clear_color.w);
        glClear(GL_COLOR_BUFFER_BIT);
        ImGui::Render();
        glfwSwapBuffers(window);
    }

    // Cleanup //
    ImGui_ImplGlfwGL3_Shutdown();
    glfwTerminate();

    return 0;
}
