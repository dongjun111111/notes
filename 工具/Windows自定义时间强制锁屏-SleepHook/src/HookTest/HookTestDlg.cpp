// HookTestDlg.cpp : implementation file
//

#include "stdafx.h"
#include "HookTest.h"
#include "HookTestDlg.h"

#ifdef _DEBUG
#define new DEBUG_NEW
#undef THIS_FILE
static char THIS_FILE[] = __FILE__;
#endif

/////////////////////////////////////////////////////////////////////////////
// CAboutDlg dialog used for App About

class CAboutDlg : public CDialog
{
public:
	CAboutDlg();

// Dialog Data
	//{{AFX_DATA(CAboutDlg)
	enum { IDD = IDD_ABOUTBOX };
	//}}AFX_DATA

	// ClassWizard generated virtual function overrides
	//{{AFX_VIRTUAL(CAboutDlg)
	protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV support
	//}}AFX_VIRTUAL

// Implementation
protected:
	//{{AFX_MSG(CAboutDlg)
	//}}AFX_MSG
	DECLARE_MESSAGE_MAP()
};

CAboutDlg::CAboutDlg() : CDialog(CAboutDlg::IDD)
{
	//{{AFX_DATA_INIT(CAboutDlg)
	//}}AFX_DATA_INIT
}

void CAboutDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialog::DoDataExchange(pDX);
	//{{AFX_DATA_MAP(CAboutDlg)
	//}}AFX_DATA_MAP
}

BEGIN_MESSAGE_MAP(CAboutDlg, CDialog)
	//{{AFX_MSG_MAP(CAboutDlg)
		// No message handlers
	//}}AFX_MSG_MAP
END_MESSAGE_MAP()

/////////////////////////////////////////////////////////////////////////////
// CHookTestDlg dialog

CHookTestDlg::CHookTestDlg(CWnd* pParent /*=NULL*/)
	: CDialog(CHookTestDlg::IDD, pParent)
{
	//{{AFX_DATA_INIT(CHookTestDlg)
		// NOTE: the ClassWizard will add member initialization here
	//}}AFX_DATA_INIT
	// Note that LoadIcon does not require a subsequent DestroyIcon in Win32
	m_hIcon = AfxGetApp()->LoadIcon(IDR_MAINFRAME);
}

void CHookTestDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialog::DoDataExchange(pDX);
	//{{AFX_DATA_MAP(CHookTestDlg)
		// NOTE: the ClassWizard will add DDX and DDV calls here
	//}}AFX_DATA_MAP
}

BEGIN_MESSAGE_MAP(CHookTestDlg, CDialog)
	//{{AFX_MSG_MAP(CHookTestDlg)
	ON_WM_SYSCOMMAND()
	ON_WM_PAINT()
	ON_WM_QUERYDRAGICON()
	ON_WM_CTLCOLOR()
	ON_WM_CLOSE()
	ON_WM_TIMER()
	//}}AFX_MSG_MAP
END_MESSAGE_MAP()

/////////////////////////////////////////////////////////////////////////////
// CHookTestDlg message handlers
_declspec(dllimport) void SetHook(HWND hwnd);
_declspec(dllimport) void UnHook(void);

BOOL CHookTestDlg::OnInitDialog()
{
	CDialog::OnInitDialog();

	// Add "About..." menu item to system menu.

	// IDM_ABOUTBOX must be in the system command range.
	ASSERT((IDM_ABOUTBOX & 0xFFF0) == IDM_ABOUTBOX);
	ASSERT(IDM_ABOUTBOX < 0xF000);
		SetWindowLong(this->GetSafeHwnd(),GWL_EXSTYLE,
		GetWindowLong(this->GetSafeHwnd(),GWL_EXSTYLE)^0x80000);
	HINSTANCE hInst = LoadLibrary("User32.DLL"); //显式加载DLL
	if(hInst) 
	{            
		typedef BOOL (WINAPI *MYFUNC)(HWND,COLORREF,BYTE,DWORD);          
		MYFUNC fun = NULL;     
		fun=(MYFUNC)GetProcAddress(hInst, "SetLayeredWindowAttributes");//取得SetLayeredWindowAttributes函数指针
		if(fun)fun(this->GetSafeHwnd(),0,128,2);     
		FreeLibrary(hInst); 
	}

	CMenu* pSysMenu = GetSystemMenu(FALSE);
	if (pSysMenu != NULL)
	{
		CString strAboutMenu;
		strAboutMenu.LoadString(IDS_ABOUTBOX);
		if (!strAboutMenu.IsEmpty())
		{
			pSysMenu->AppendMenu(MF_SEPARATOR);
			pSysMenu->AppendMenu(MF_STRING, IDM_ABOUTBOX, strAboutMenu);
		}
	}

	// Set the icon for this dialog.  The framework does this automatically
	//  when the application's main window is not a dialog
	SetIcon(m_hIcon, TRUE);			// Set big icon
	SetIcon(m_hIcon, FALSE);		// Set small icon
	
	// TODO: Add extra initialization here
	int cxScreen,cyScreen;
	cxScreen=GetSystemMetrics(SM_CXSCREEN);
	cyScreen=GetSystemMetrics(SM_CYSCREEN);
	SetWindowPos(&wndTopMost,0,0,cxScreen,cyScreen,SWP_HIDEWINDOW&SWP_NOOWNERZORDER&SWP_NOMOVE);//SWP_SHOWWINDOW);
	
	//m_clr=RGB(58,13,140);
	m_clr=RGB(0,0,0);


	m_destroy=FALSE;
	m_hooked=FALSE;
	


	CStdioFile file;
	file.Open("TimeSet.ini",CFile::modeNoTruncate|CFile::modeReadWrite|CFile::modeCreate|CFile::shareDenyRead|CFile::shareDenyWrite);
	
	CString strRead;
	while(file.ReadString(strRead))
	{
		m_mySetTime.push_back(strRead);
	}

	if(m_mySetTime.size()==0)
	{
		m_destroy=TRUE;
		MessageBox("Please Set Your Time In TimeSet.ini, The Format Is: \n		     hh:mm:ss \n		     hh:mm:ss \n          Then, Run This Software Again! ");
		AfxGetMainWnd()->SendMessage(WM_CLOSE);
	}
	else
	{
		m_showed=TRUE;
		SetTimer(1,1000,NULL);
	}

	SetWindowText("SleepHook");

	return TRUE;  // return TRUE  unless you set the focus to a control
}

void CHookTestDlg::OnSysCommand(UINT nID, LPARAM lParam)
{
	if ((nID & 0xFFF0) == IDM_ABOUTBOX)
	{
		CAboutDlg dlgAbout;
		dlgAbout.DoModal();
	}
	else
	{
		CDialog::OnSysCommand(nID, lParam);
	}
}

// If you add a minimize button to your dialog, you will need the code below
//  to draw the icon.  For MFC applications using the document/view model,
//  this is automatically done for you by the framework.

void CHookTestDlg::OnPaint() 
{
	if (IsIconic())
	{
		CPaintDC dc(this); // device context for painting

		SendMessage(WM_ICONERASEBKGND, (WPARAM) dc.GetSafeHdc(), 0);

		// Center icon in client rectangle
		int cxIcon = GetSystemMetrics(SM_CXICON);
		int cyIcon = GetSystemMetrics(SM_CYICON);
		CRect rect;
		GetClientRect(&rect);
		int x = (rect.Width() - cxIcon + 1) / 2;
		int y = (rect.Height() - cyIcon + 1) / 2;

		// Draw the icon
		dc.DrawIcon(x, y, m_hIcon);
	}
	else
	{
		CDialog::OnPaint();
	}
}

// The system calls this to obtain the cursor to display while the user drags
//  the minimized window.
HCURSOR CHookTestDlg::OnQueryDragIcon()
{
	return (HCURSOR) m_hIcon;
}

HBRUSH CHookTestDlg::OnCtlColor(CDC* pDC, CWnd* pWnd, UINT nCtlColor) 
{
	HBRUSH hbr = CDialog::OnCtlColor(pDC, pWnd, nCtlColor);
	
	// TODO: Change any attributes of the DC here
	switch(nCtlColor)
	{
	case CTLCOLOR_DLG:
		 HBRUSH aBrush;
		 aBrush=CreateSolidBrush(m_clr);
		 hbr=aBrush;
		 break;
	}
	// TODO: Return a different brush if the default is not desired
	return hbr;
}

void CHookTestDlg::OnClose() 
{
	// TODO: Add your message handler code here and/or call default
	
	if(m_destroy)
	{
		CDialog::OnClose();
	}
}

BOOL CHookTestDlg::PreTranslateMessage(MSG* pMsg) 
{
	// TODO: Add your specialized code here and/or call the base class
	
	if(pMsg->message==WM_KEYDOWN&&pMsg->wParam==VK_ESCAPE)
		return TRUE;
	if(pMsg->message==WM_KEYDOWN&&pMsg->wParam==VK_RETURN)
		return TRUE;

	if(pMsg->message==WM_KEYDOWN&&pMsg->wParam==VK_F7)
	{
		return TRUE;
	}

	if(pMsg->message==WM_NCLBUTTONDOWN)
	{
		return TRUE;
	}

	return CDialog::PreTranslateMessage(pMsg);
}


void CHookTestDlg::OnTimer(UINT nIDEvent) 
{
	// TODO: Add your message handler code here and/or call default
	
	COleDateTime curTime,startTime,stopTime;

	CTime t=CTime::GetCurrentTime();
	CString curTimeStr;
	curTimeStr.Format("%d:%d:%d",t.GetHour(),t.GetMinute(),t.GetSecond());
	curTime.ParseDateTime(curTimeStr);


	bool inStartStop=FALSE;

	for(int i=0;i<m_mySetTime.size();i+=2)
	{
		CString startTimeStr,stopTimeStr;
		startTimeStr=m_mySetTime.at(i);
		stopTimeStr=m_mySetTime.at(i+1);

		//sscanf(m_mySetTime.at(i),"%s %s",startTimeStr,stopTimeStr);

		startTime.ParseDateTime(startTimeStr);
		stopTime.ParseDateTime(stopTimeStr);

		if(startTime>=stopTime)
		{
			if(!(curTime>stopTime&&curTime<startTime))
			{
				inStartStop=TRUE;
				break;
			}

		}
		else
		{
			if(curTime>=startTime&&curTime<stopTime)
			{
				inStartStop=TRUE;
				break;
			}
		}

	}


	if(inStartStop)
	{
			if(!m_hooked)
			{
				SetHook(this->m_hWnd);
				hookTaskmgr.Open("C:\\Windows\\system32\\taskmgr.exe",CFile::shareDenyRead|CFile::shareDenyWrite, NULL);        //不可被其他程序访问
				m_hooked=TRUE;
			}

			if(!m_showed)
			{
				::SendMessage(this->m_hWnd, WM_SYSCOMMAND, SC_MAXIMIZE, 0);
				ShowWindow(SW_SHOW);
				m_showed=TRUE;
			}
	}
	else
	{
			if(m_hooked)
			{
				UnHook();
				hookTaskmgr.Close();
				m_hooked=FALSE;
			}

			if(m_showed)
			{
				::SendMessage(this->m_hWnd, WM_SYSCOMMAND, SC_MINIMIZE, 0);
				ShowWindow(SW_HIDE);
				m_showed=FALSE;
			}
	}



	CDialog::OnTimer(nIDEvent);
}
