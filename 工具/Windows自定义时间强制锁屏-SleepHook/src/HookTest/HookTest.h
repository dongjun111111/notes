// HookTest.h : main header file for the HOOKTEST application
//

#if !defined(AFX_HOOKTEST_H__F555F975_EBA5_4316_8676_82EE9F083BF0__INCLUDED_)
#define AFX_HOOKTEST_H__F555F975_EBA5_4316_8676_82EE9F083BF0__INCLUDED_

#if _MSC_VER > 1000
#pragma once
#endif // _MSC_VER > 1000

#ifndef __AFXWIN_H__
	#error include 'stdafx.h' before including this file for PCH
#endif

#include "resource.h"		// main symbols

/////////////////////////////////////////////////////////////////////////////
// CHookTestApp:
// See HookTest.cpp for the implementation of this class
//

class CHookTestApp : public CWinApp
{
public:
	CHookTestApp();

// Overrides
	// ClassWizard generated virtual function overrides
	//{{AFX_VIRTUAL(CHookTestApp)
	public:
	virtual BOOL InitInstance();
	//}}AFX_VIRTUAL

// Implementation

	//{{AFX_MSG(CHookTestApp)
		// NOTE - the ClassWizard will add and remove member functions here.
		//    DO NOT EDIT what you see in these blocks of generated code !
	//}}AFX_MSG
	DECLARE_MESSAGE_MAP()
};


/////////////////////////////////////////////////////////////////////////////

//{{AFX_INSERT_LOCATION}}
// Microsoft Visual C++ will insert additional declarations immediately before the previous line.

#endif // !defined(AFX_HOOKTEST_H__F555F975_EBA5_4316_8676_82EE9F083BF0__INCLUDED_)
