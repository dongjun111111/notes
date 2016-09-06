#ifndef _WIN32_WINNT
#define _WIN32_WINNT 0x0500
#endif

#include <windows.h>
#include <WinUser.h>


#define WM_RECVDATA WM_USER+1
HHOOK g_hMouse=NULL;
HHOOK g_hKeyboard=NULL;

#pragma data_seg("MySec")
HWND g_hWnd=NULL;
#pragma data_seg()

#pragma comment(linker,"/section:MySec,RWS")

LRESULT CALLBACK MouseProc(int nCode,WPARAM wParam,LPARAM lParam)
{
 return 1;
}

//HC_ACTION 
LRESULT CALLBACK KeyboardProc(int nCode,WPARAM wParam,LPARAM lParam)
{
/*
	// LowLevel Hook
	BOOL fEatKeystroke = FALSE;
	PKBDLLHOOKSTRUCT p = NULL;

	if (nCode == HC_ACTION)
	{
		p = (PKBDLLHOOKSTRUCT) lParam;
		switch (wParam)
		{
			case WM_KEYDOWN:
				if (p->vkCode == VK_F8)
				{
					::MessageBox(NULL,"Let's make things better and better!\n","HQ Tech",MB_OK);
					break;
				}
			case WM_SYSKEYDOWN:
			case WM_KEYUP:
			case WM_SYSKEYUP:
				fEatKeystroke = (p->vkCode == VK_LWIN) || (p->vkCode == VK_RWIN) ||  // ÆÁ±ÎWin
                            // ÆÁ±ÎAlt+Tab
                            ((p->vkCode == VK_TAB) && ((p->flags & LLKHF_ALTDOWN) != 0)) ||
                            // ÆÁ±ÎAlt+Esc
                            ((p->vkCode == VK_ESCAPE) && ((p->flags & LLKHF_ALTDOWN) != 0)) ||
                            // ÆÁ±ÎCtrl+Esc
                            ((p->vkCode == VK_ESCAPE) && ((GetKeyState(VK_CONTROL) & 0x8000) != 0));
				break;
			default:
				break;
		}
	}
*/


/*
    // Global Hook
 	if(VK_F7==wParam)
	{
		SendMessage(g_hWnd,WM_CLOSE,0,0);
		UnhookWindowsHookEx(g_hKeyboard);
		UnhookWindowsHookEx(g_hMouse);
	}

	return 1;
*/


	// LowLevel Hook
	BOOL fEatKeystroke = TRUE;
	PKBDLLHOOKSTRUCT p = NULL;

	if(nCode==HC_ACTION)
	{
		p = (PKBDLLHOOKSTRUCT) lParam;
		if(wParam==WM_KEYDOWN&&p->vkCode==VK_F7)
			fEatKeystroke = FALSE;
	}

	return (fEatKeystroke ? TRUE : CallNextHookEx(g_hKeyboard,nCode,wParam,lParam)); 

}

void SetHook(HWND hwnd)
{
 g_hWnd=hwnd;
 g_hMouse=SetWindowsHookEx(WH_MOUSE,MouseProc,GetModuleHandle("Hook"),0);
 //g_hKeyboard=SetWindowsHookEx(WH_KEYBOARD,KeyboardProc,GetModuleHandle("Hook"),0); // Global Hook
 g_hKeyboard=SetWindowsHookEx(WH_KEYBOARD_LL,KeyboardProc,GetModuleHandle("Hook"),0); // LowLevel Hook
}

void UnHook(void)
{
	UnhookWindowsHookEx(g_hKeyboard);
	UnhookWindowsHookEx(g_hMouse);
}