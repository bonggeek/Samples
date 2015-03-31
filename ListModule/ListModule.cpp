#include <windows.h>
#include <stdio.h>
#include <wchar.h>
#include <psapi.h>

void ShowError(HRESULT hResult, WCHAR* msg)
{
    LPVOID lpMsgBuf;

    ::FormatMessageW( FORMAT_MESSAGE_ALLOCATE_BUFFER | FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_IGNORE_INSERTS,
                   NULL, hResult, MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT), (LPWSTR) &lpMsgBuf, 0, NULL );

    wprintf(L"ERROR!!! %s %08X: %s\n", msg, hResult, lpMsgBuf);

    LocalFree(lpMsgBuf);
}

void Modules(DWORD procId)
{
    HANDLE hProc = ::OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ, FALSE, procId);
    if(hProc == NULL)
    {
        ShowError(::GetLastError(), L"OpenProcess");
        return;
    }
    
    HMODULE hMods[10000];
    DWORD cbNeeded;
    if(::EnumProcessModules(hProc, hMods, sizeof(hMods), &cbNeeded))
    {
        int count = cbNeeded / sizeof(HMODULE);
        for(int i = 0; i <= count; ++i)
        {
            HMODULE hMod = hMods[i];
            WCHAR szFileName[MAX_PATH];
            if(::GetModuleFileNameExW(hProc, hMod, szFileName, sizeof(szFileName)/sizeof(WCHAR)))
            {
                wprintf(L"\t(0x%08X)\t%s\n", (DWORD)hMod, szFileName);
            }
        }
    }

    ::CloseHandle(hProc);
}

void Processes(DWORD *aProc, DWORD nProc, WCHAR *procName)
{
    if(procName != NULL)
        wprintf(L"Searching for %s in %d processes", procName, nProc);

    for(int i = 0; i < nProc; ++i)
    {
        DWORD procId = aProc[i];
        HANDLE hProc = ::OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ, FALSE, procId);

        WCHAR szProcFileName[MAX_PATH] = L"";
        if(::GetProcessImageFileNameW(hProc, szProcFileName, MAX_PATH) != 0)
        {
            if(procName == NULL || FindNLSString (LOCALE_INVARIANT, NORM_IGNORECASE, szProcFileName, -1, procName, -1, NULL) != -1)
            {
                wprintf(L"\n\n%s (%d)\n========================================================\n", szProcFileName, procId);
                Modules(procId);
            }
        }

        CloseHandle(hProc);
    }
}

void usage()
{
    // ok always have 42!!
    wprintf(L"        @@@@   @@@@@@\n");
    wprintf(L"       @@ @@  @@   @@\n");
    wprintf(L"      @@  @@       @@\n");
    wprintf(L"     @@   @@      @@ \n");
    wprintf(L"     @@@@@@@    @@   \n");
    wprintf(L"          @@   @@    \n");
    wprintf(L"          @@   @@@@@@\n");
    wprintf(L"\n\n");
    wprintf(L"ListModule <filter>\n");
    wprintf(L"      List all modules of processes that contain the \"filter\" in its full path\n\n"); 
    wprintf(L"ListModule\n");
    wprintf(L"      List all modules of all processes\n\n"); 
}

int wmain(int argc, WCHAR **argv)
{
    if(argc == 2 && *argv[1] == '-')
    {
        usage();
        return -1;
    }
    
    DWORD aProcesses[4000], cbNeeded, cProcesses;
    WCHAR *procName = NULL;
    if(argc == 2)
        procName = argv[1];

    ::EnumProcesses(aProcesses, sizeof(aProcesses), &cbNeeded);
    cProcesses = cbNeeded / sizeof(DWORD);
    Processes(aProcesses, cProcesses, procName);
}
